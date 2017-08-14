package bashir

import (
	"fmt"
	"os"
)

type EnvVarMaker struct{}

func NewEnvVarMaker() *EnvVarMaker {
	return &EnvVarMaker{}
}

func (m *EnvVarMaker) GetEnvVarValue(envVarName string) (string, error) {
	value := os.Getenv(envVarName)
	var err error

	if value == "" {
		err = fmt.Errorf("Could not find value for key: %s", envVarName)
	}

	keyValue := fmt.Sprintf("%s=%s", envVarName, value)

	return keyValue, err
}
