package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	V string `json:"v"`
}

/** 加载配置文件 */
func loadConfig(configFile string) {
	// read config file
	file, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	// parse config file
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	// use config values
	if config.V != "NyarukoDNSv1" {
		panic("config file var error")
	}
}
