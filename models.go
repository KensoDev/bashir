package bashir

type Config struct {
	Defaults    DefaultConfig `yaml:"defaults"`
	Commands    []Command     `yaml:"commands"`
	CommandArgs []string      `yaml:"command_args"`
}

type DefaultConfig struct {
	EnvVars []string `yaml:"envvars"`
	Args    []string `yaml:"args"`
}

type Command struct {
	EnvVars    []string `yaml:"envvars"`
	WorkingDir string   `yaml:"cwd"`
	Container  string   `yaml:"container"`
	Name       string   `yaml:"name"`
	Command    string   `yaml:"command"`
	Args       []string `yaml:"args"`
	ReportTo   []string `yaml:"report_to"`
}
