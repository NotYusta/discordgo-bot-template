package utils

import "go-dc-bot/utils/config"

var configClass = ConfigClass{
	GetConfigurationYaml:  config.GetConfigurationYaml,
	GetConfigurationFlags: config.GetConfigurationFlags,
}

func GetConfig() *ConfigClass {
	return &configClass
}
