package bashir

type Config struct {
	Commands    []Command `yaml:"commands"`
	CommandArgs []string  `yaml:"command_args"`
}

type Command struct {
	WorkingDir string   `yaml:"cwd"`
	Name       string   `yaml:"name"`
	Command    string   `yaml:"command"`
	Args       []string `yaml:"args"`
	VirtualEnv string   `yaml:"virtualenv"`
	ReportTo   []string `yaml:"report_to"`
}
