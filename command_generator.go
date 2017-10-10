package bashir

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	prompt "github.com/segmentio/go-prompt"
)

type CommandGenerator struct {
	config  *Config
	maker   *EnvVarMaker
	matcher *regexp.Regexp
}

func NewCommandGenerator(config *Config) *CommandGenerator {
	maker := NewEnvVarMaker()
	matcher := regexp.MustCompile("(.*):ask\\?")

	return &CommandGenerator{
		config:  config,
		maker:   maker,
		matcher: matcher,
	}
}

func (c *CommandGenerator) GetEnvVarArgs(envVars []string) []string {
	out := []string{}

	for _, v := range envVars {
		val, err := c.maker.GetEnvVarValue(v)
		if err != nil {
			// warn? exit?
		}
		out = append(out, "-e", val)
	}

	return out
}

func (c *CommandGenerator) GetEnvVarsArguments(command Command) []string {
	defaultEnvVars := c.config.Defaults.EnvVars
	commandEnvVars := command.EnvVars

	out := []string{}

	out = append(out, c.GetEnvVarArgs(defaultEnvVars)...)
	out = append(out, c.GetEnvVarArgs(commandEnvVars)...)

	return out
}

func (c *CommandGenerator) IsArgumentAskable(arg string) (substring string, ask bool) {
	substring = ""
	ask = c.matcher.MatchString(arg)

	if ask {
		substring = c.matcher.FindStringSubmatch(arg)[1]
	}

	return
}

func (c *CommandGenerator) GetArgumentValueOrAsk(arguments []string) []string {
	out := []string{}

	for _, val := range arguments {
		name, askable := c.IsArgumentAskable(val)

		if askable {
			val = prompt.String(fmt.Sprintf("Value for %s?", name))
		}

		out = append(out, val)
	}

	return out
}

func (c *CommandGenerator) GetArgs(command Command) []string {
	defaultArgs := c.config.Defaults.Args
	commandArgs := command.Args
	out := []string{}

	out = append(out, c.GetArgumentValueOrAsk(defaultArgs)...)
	out = append(out, c.GetArgumentValueOrAsk(commandArgs)...)

	return out
}

func (c *CommandGenerator) GetVolumeArguments(command Command) []string {
	defaultVolumeArgs := c.config.Defaults.Volumes
	out := []string{}

	out = append(out, c.GetVolumeArgs(defaultVolumeArgs)...)

	fmt.Println(out)

	return out
}

func (c *CommandGenerator) GetVolumeArgs(volumes []string) []string {
	out := []string{}

	for _, v := range volumes {

		if v[0] == '~' {
			dir := os.Getenv("HOME")
			v = filepath.Join(dir, v[1:])
		}

		out = append(out, "-v", v)
	}

	return out
}
