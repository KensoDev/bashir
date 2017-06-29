package bashir

type Config struct {
	Commands    []Command `yaml:"commands"`
	CommandArgs []string  `yaml:"command_args"`
}

type Command struct {
	Name       string   `yaml:"name"`
	Command    string   `yaml:"command"`
	VirtualEnv string   `yaml:"virtualenv"`
	ReportTo   []string `yaml:"report_to"`
}
