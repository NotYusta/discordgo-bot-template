package config

type FlagArguments struct {
	ConfigPath string
}

type YamlStructure struct {
	BotToken string `yaml:"botToken"`
}
