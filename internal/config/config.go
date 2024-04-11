package config

import (
	"encoding/json"
	"os"
)

type RunConfig struct {
	BatchSize string `json:"batch-size"`
	Frequency string `json:"frequency"`
}

func ReadConfig() (RunConfig, error) {
	// read config
	file, err := os.Open("config.json")
	if err != nil {
		return RunConfig{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// parse config
	var config RunConfig
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return RunConfig{}, err
	}

	return config, nil
}