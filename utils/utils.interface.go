package utils

import "go-dc-bot/utils/config"

type ConfigClass struct {
	GetConfigurationFlags func() *config.FlagArguments
	GetConfigurationYaml  func() *config.ConfStructure
}
