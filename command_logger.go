package bashir

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

type CommandLogger struct {
	Command     Command
	LogFileName string
}

func NewCommandLogger(command Command) *CommandLogger {
	logFileName := command.Out

	if logFileName == "" {

	}

	return &CommandLogger{
		Command:     command,
		LogFileName: logFileName,
	}
}

func TempFileName(suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join("./", hex.EncodeToString(randBytes)+suffix)
}

func (c *CommandLogger) LogCommandOutput(out []byte) (string, error) {
	fileLocation := filepath.Join("./", c.Command.Out)

	if c.Command.Out == "" {
		fileLocation = TempFileName(".log")
	}

	f, err := os.Create(fileLocation)
	_, err = f.Write(out)

	defer f.Close()

	f.Sync()

	if err != nil {
		return "", err
	}

	fmt.Println("Saved log for %s in %s", c.Command.Name, fileLocation)

	return fileLocation, nil
}
