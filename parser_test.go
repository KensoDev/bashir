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

func (s *ParserSuite) TestParseYamlFileAndGetDescription(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, err := parser.ParseConfigurationFile()

	c.Assert(err, IsNil)
	c.Assert(config.Commands[0].Description, Not(Equals), "")
}

func (s *ParserSuite) TestParseOfYamlFileGetCommand(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, err := parser.ParseConfigurationFile()

	c.Assert(err, IsNil)
	c.Assert(config.Commands[0].ImageName, Equals, "kensodev/bashir")
}

func (s *ParserSuite) TestParseOfYamlFileWithSlackConfig(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, err := parser.ParseConfigurationFile()

	c.Assert(err, IsNil)
	c.Assert(config, NotNil)
	c.Assert(config.Slack.WebhookUrl, Equals, "https://test.com")
}

func (s *ParserSuite) TestParserWithVolumes(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, err := parser.ParseConfigurationFile()

	c.Assert(err, IsNil)
	c.Assert(config, NotNil)
	c.Assert(len(config.Defaults.Volumes), Equals, 1)
}

func (s *ParserSuite) TestParserWithVolumesContent(c *C) {
	configLocation := "fixtures/sample_config.yml"
	parser := NewParser(configLocation)
	config, err := parser.ParseConfigurationFile()

	c.Assert(err, IsNil)
	c.Assert(config, NotNil)
	c.Assert(config.Defaults.Volumes[0], Equals, "~/.aws:/some/.aws")
}
