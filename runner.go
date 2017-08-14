package bashir

import (
	"fmt"
	"log"
	"os/exec"
)

type Runner struct {
	Config         *Config
	SlackClient    *SlackClient
	SlackGenerator *SlackGenerator
}

func NewRunner(config *Config) *Runner {
	slackClient := NewSlackClient(config)
	slackGenerator := NewSlackGenerator()

	return &Runner{
		Config:         config,
		SlackClient:    slackClient,
		SlackGenerator: slackGenerator,
	}
}

type RunResult struct{}

func (r *Runner) RunCommands() *RunResult {
	generator := NewCommandGenerator(r.Config)

	for _, command := range r.Config.Commands {
		message := r.SlackGenerator.GetSlackMessageForStart(command)
		err := r.SlackClient.SendMessage(message)

		if err != nil {
			log.Fatal(err)
		}

		envVars := generator.GetEnvVarsArguments(command)
		args := generator.GetArgs(command)

		commandArgs := []string{"run", "-i"}

		commandArgs = append(commandArgs, envVars...)
		commandArgs = append(commandArgs, command.ImageName)
		commandArgs = append(commandArgs, command.Command)
		commandArgs = append(commandArgs, args...)

		cmd := exec.Command("docker", commandArgs...)

		fmt.Println(cmd)

		out, err := cmd.CombinedOutput()

		message = r.SlackGenerator.GetSlackMessageForEnd(command, err, out)
		err = r.SlackClient.SendMessage(message)

		fmt.Println(string(out))
		fmt.Println(err)
	}

	return &RunResult{}
}
