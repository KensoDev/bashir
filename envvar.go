package bashir

import (
	"fmt"
	"os"
	"strings"
)

type EnvVarMaker struct{}

func NewEnvVarMaker() *EnvVarMaker {
	return &EnvVarMaker{}
}

func (m *EnvVarMaker) GetEnvVarValue(envVarName string) (string, error) {
	// If the string already contains the value, return it
	if strings.Contains(envVarName, "=") {
		return envVarName, nil
	}

	value := os.Getenv(envVarName)
	var err error

	if value == "" {
		err = fmt.Errorf("Could not find value for key: %s", envVarName)
	}

	keyValue := fmt.Sprintf("%s=%s", envVarName, value)

	return keyValue, err
}
