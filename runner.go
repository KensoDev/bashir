package bashir

import (
	"fmt"
	"os/exec"
)

type Runner struct {
	Config *Config
}

func NewRunner(config *Config) *Runner {
	return &Runner{
		Config: config,
	}
}

type RunResult struct{}

func (r *Runner) RunCommands() *RunResult {
	generator := NewCommandGenerator(r.Config)

	for _, command := range r.Config.Commands {
		envVars := generator.GetEnvVarsArguments(command)
		args := generator.GetArgs(command)

		commandArgs := []string{"run", "-i"}

		commandArgs = append(commandArgs, envVars...)
		commandArgs = append(commandArgs, command.Container)
		commandArgs = append(commandArgs, command.Command)
		commandArgs = append(commandArgs, args...)

		cmd := exec.Command("docker", commandArgs...)

		fmt.Println(cmd)

		out, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(out))
		fmt.Println(err)
	}

	return &RunResult{}
}
