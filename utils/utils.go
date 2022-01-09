package utils

import "flag"

func GetConfigFlag() string {
	configPathFlag := flag.String("config", "./config.yml", "config.yml file path")
	flag.Parse()

	return *configPathFlag
}
