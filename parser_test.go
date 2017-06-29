package bashir

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestParser(t *testing.T) { TestingT(t) }

type ParserSuite struct{}

var _ = Suite(&ParserSuite{})

func (s *ParserSuite) TestParseOfYamlFile(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, err := parser.ParseConfigurationFile()

	c.Assert(err, IsNil)
	c.Assert(config, NotNil)
}

func (s *ParserSuite) TestParseOfYamlFileGetUsers(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, err := parser.ParseConfigurationFile()

	c.Assert(err, IsNil)
	c.Assert(config.Commands[0].ReportTo[0], Equals, "KensoDev")
}

func (s *ParserSuite) TestParseOfYamlFileGetCommandArgs(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, err := parser.ParseConfigurationFile()

	c.Assert(err, IsNil)
	c.Assert(len(config.CommandArgs), Equals, 2)
}