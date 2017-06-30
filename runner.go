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
	for _, command := range r.Config.Commands {
		fmt.Println(command.WorkingDir)

		cmd = exec.Command(command.Command, command.Args...)
		cmd.Dir = command.WorkingDir

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
