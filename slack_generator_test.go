package bashir

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestSlackGenerator(t *testing.T) { TestingT(t) }

type SlackGeneratorSuite struct{}

var _ = Suite(&SlackGeneratorSuite{})

func (s *SlackGeneratorSuite) TestGetSlackMessageForMessageStart(c *C) {
	command := Command{
		Name:        "Migrating the database",
		Description: "This is the description",
	}

	c.Assert(command, NotNil)
}
