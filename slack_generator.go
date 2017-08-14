package bashir

import "fmt"

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
