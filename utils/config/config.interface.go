package config

type FlagArguments struct {
	ConfigPath string
}

type ConfStructure struct {
	BotToken string `yaml:"botToken"`
	Options  struct {
		Activity struct {
			Enable      bool   `yaml:"enable"`
			Type        string `yaml:"type"`
			Description string `yaml:"description"`
			Url         string `yaml:"url"`
		} `yaml:"activity"`
		Prefix string `yaml:"prefix"`
	} `yaml:"options"`
}
