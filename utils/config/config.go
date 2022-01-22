package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
)

var configurationFlags = getFlag()
var configurationYaml = getYaml()

func getFlag() *FlagArguments {
	configPathFlag := flag.String("config", "./config.yml", "config.yml file path")
	flag.Parse()
	configPath := path.Join(*configPathFlag)
	flagArguments := FlagArguments{
		ConfigPath: configPath,
	}

	return &flagArguments
}

func getYaml() *YamlStructure {
	cfgFile := &YamlStructure{}
	yamlFile, err := ioutil.ReadFile(GetConfigurationFlags().ConfigPath)
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
	return configurationFlags
}

func GetConfigurationYaml() *YamlStructure {
	return configurationYaml
}
