package bashir

import "fmt"

type CommandGenerator struct {
	config *Config
	maker  *EnvVarMaker
}

func NewCommandGenerator(config *Config) *CommandGenerator {
	maker := NewEnvVarMaker()

	return &CommandGenerator{
		config: config,
		maker:  maker,
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

func (c *CommandGenerator) GetArgs(command Command) []string {
	defaultArgs := c.config.Defaults.Args
	commandArgs := command.Args
	out := []string{}

	out = append(out, defaultArgs...)
	out = append(out, commandArgs...)

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
		out = append(out, "-v", v)
	}

	return out
}
