package bashir

type Config struct {
	Debug       bool          `yaml:"debug"`
	Slack       SlackConfig   `yaml:"slack"`
	Defaults    DefaultConfig `yaml:"defaults"`
	Commands    []Command     `yaml:"commands"`
	CommandArgs []string      `yaml:"command_args"`
}

type SlackConfig struct {
	WebhookUrl string `yaml:"webhook_url"`
	Channel    string `yaml:"channel"`
	Icon       string `yaml:"icon"`
	BotName    string `yaml:"bot_name"`
}

type DefaultConfig struct {
	EnvVars []string `yaml:"envvars"`
	Args    []string `yaml:"args"`
	Volumes []string `yaml:"volumes"`
}

type Command struct {
	EnvVars     []string `yaml:"envvars"`
	WorkingDir  string   `yaml:"cwd"`
	ImageName   string   `yaml:"image_name"`
	Out         string   `yaml:"out"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Command     string   `yaml:"command"`
	Args        []string `yaml:"args"`
	ReportTo    []string `yaml:"report_to"`
}
