package bashir

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Parser struct {
	ConfigLocation string
}

func NewParser(configLocation string) *Parser {
	return &Parser{
		ConfigLocation: configLocation,
	}
}

func (p *Parser) ParseConfigurationFile() (*Config, error) {
	config := Config{}

	file, err := ioutil.ReadFile(p.ConfigLocation)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(file), &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
