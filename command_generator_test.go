package bashir

import (
	"fmt"
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func TestCommandGenerator(t *testing.T) { TestingT(t) }

type CommandGeneratorSuite struct{}

var _ = Suite(&CommandGeneratorSuite{})

func (s *CommandGeneratorSuite) TestEnvVars(c *C) {
	_ = os.Setenv("TEST_A", "FAKE_VALUE")
	_ = os.Setenv("TEST_B", "FAKE_VALUE")
	_ = os.Setenv("TEST_C", "FAKE_VALUE")

	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, _ := parser.ParseConfigurationFile()

	command := config.Commands[0]
	commandGenerator := NewCommandGenerator(config)
	envArgs := commandGenerator.GetEnvVarsArguments(command)

	c.Assert(len(envArgs), Equals, 6)
}

func (s *CommandGeneratorSuite) TestEnvVarValue(c *C) {
	_ = os.Setenv("TEST_A", "FAKE_VALUE1")
	_ = os.Setenv("TEST_B", "FAKE_VALUE2")
	_ = os.Setenv("TEST_C", "FAKE_VALUE3")

	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, _ := parser.ParseConfigurationFile()

	command := config.Commands[0]
	commandGenerator := NewCommandGenerator(config)
	envArgs := commandGenerator.GetEnvVarsArguments(command)

	_ = os.Unsetenv("TEST_A")
	_ = os.Unsetenv("TEST_B")
	_ = os.Unsetenv("TEST_C")

	c.Assert(envArgs[1], Equals, "TEST_A=FAKE_VALUE1")
}

func (s *CommandGeneratorSuite) TestEnvVarValueBlank(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, _ := parser.ParseConfigurationFile()

	command := config.Commands[0]
	commandGenerator := NewCommandGenerator(config)
	envArgs := commandGenerator.GetEnvVarsArguments(command)

	c.Assert(envArgs[1], Equals, "TEST_A=")
}

func (s *CommandGeneratorSuite) TestArgs(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, _ := parser.ParseConfigurationFile()

	command := config.Commands[0]
	commandGenerator := NewCommandGenerator(config)
	args := commandGenerator.GetArgs(command)

	c.Assert(len(args), Equals, 7)
}

func (s *CommandGeneratorSuite) TestArgsValues(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, _ := parser.ParseConfigurationFile()

	command := config.Commands[0]
	commandGenerator := NewCommandGenerator(config)
	args := commandGenerator.GetArgs(command)

	c.Assert(args[0], Equals, "-x")
}

func (s *CommandGeneratorSuite) TestVolumeArgs(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, _ := parser.ParseConfigurationFile()

	command := config.Commands[0]
	commandGenerator := NewCommandGenerator(config)
	args := commandGenerator.GetVolumeArguments(command)

	c.Assert(args[0], Equals, "-v")
	c.Assert(args[1], Equals, fmt.Sprintf("%s/.aws:/some/.aws", os.Getenv("HOME")))
}

func (s *CommandGeneratorSuite) TestIsArgumentAskable(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, _ := parser.ParseConfigurationFile()

	commandGenerator := NewCommandGenerator(config)
	sub, askable := commandGenerator.IsArgumentAskable("environment:ask?")

	c.Assert(sub, Equals, "environment")
	c.Assert(askable, Equals, true)
}

func (s *CommandGeneratorSuite) TestPathExpansion(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, _ := parser.ParseConfigurationFile()

	commandGenerator := NewCommandGenerator(config)
	volumes := commandGenerator.GetVolumeArgs([]string{"~/.aws:/root/.aws"})

	c.Assert(len(volumes), Equals, 2)
}
