package utils

import "flag"

func ParseConfigPath() string{
	configPath := ""
	flag.StringVar(&configPath, "c", "configuration.json", "set configuration")
	flag.Parse()
	return configPath
}
