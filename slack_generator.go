package bashir

import (
	"fmt"
	"strings"
)

type SlackGenerator struct{}

func NewSlackGenerator() *SlackGenerator {
	return &SlackGenerator{}
}

func (g *SlackGenerator) GetSlackMessageForStart(command Command) string {
	var message string

	if command.Description == "" {
		message = fmt.Sprintf("Starting: %s", command.Name)
		return message
	}

	message = "%s\n```%s```"

	message = fmt.Sprintf(message, command.Name, command.Description)

	return message
}

func (g *SlackGenerator) GetSlackMessageForEnd(command Command, runErr error) string {
	result := "Command finished successfully"

	if runErr != nil {
		mentions := strings.Join(command.ReportTo, ", ")
		result = fmt.Sprintf("Command failed to run: cc/ %s. Result was %s", mentions, runErr)
	}

	message := fmt.Sprintf("Finished running: %s. %s", command.Name, result)
	return message
}
