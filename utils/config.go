package utils

import (
	"flag"
	"path"
)

var configurationFlags = getFlag()

type FlagArguments struct {
	ConfigPath *string
}

type ConfigStructure struct {
	Hits      int64
	Time      int64
	CamelCase string `yaml:"camelCase"`
}

func getFlag() FlagArguments {
	configPathFlag := flag.String("config", "./config.yml", "config.yml file path")
	flag.Parse()
	configPath := path.Join(*configPathFlag)

	return FlagArguments{
		ConfigPath: &configPath,
	}
}

func GetConfigurationFlags() *FlagArguments {
	return &configurationFlags
}
