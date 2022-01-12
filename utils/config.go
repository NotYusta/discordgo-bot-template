package utils

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
)

var configurationFlags = getFlag()
var configurationYaml = getYaml()

type FlagArguments struct {
	ConfigPath *string
}

type ConfigYamlStructure struct {
	BotToken string `yaml:"botToken"`
}

func getFlag() FlagArguments {
	configPathFlag := flag.String("config", "./config.yml", "config.yml file path")
	flag.Parse()
	configPath := path.Join(*configPathFlag)

	return FlagArguments{
		ConfigPath: &configPath,
	}
}

func getYaml() *ConfigYamlStructure {
	cfgFile := &ConfigYamlStructure{}
	yamlFile, err := ioutil.ReadFile(*GetConfigurationFlags().ConfigPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, cfgFile)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return cfgFile
}

func GetConfigurationFlags() *FlagArguments {
	return &configurationFlags
}

func GetConfigurationYaml() *ConfigYamlStructure {
	return configurationYaml
}
