package utils

import (
	"errors"
	"go-dc-bot/utils/config"
)

var configClass = ConfigClass{
	GetConfigurationYaml:  config.GetConfigurationYaml,
	GetConfigurationFlags: config.GetConfigurationFlags,
}

func GetConfig() *ConfigClass {
	return &configClass
}

func RemoveArrayFromIndex(s []string, index int) ([]string, error) {
	if index >= len(s) {
		return nil, errors.New("out of Range Error")
	}

	return append(s[:index], s[index+1:]...), nil
}
