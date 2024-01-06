package main

import (
	"embed"
	"encoding/json"
)

type Config struct {
	ProgramTooltip string   `json:"programTooltip"`
	Actions        []Action `json:"actions"`
}

type Action struct {
	Name    string          `json:"name"`
	Tooltip string          `json:"tooltip"`
	Webhook string          `json:"webhook"`
	Data    json.RawMessage `json:"data"`
	Type    string          `json:"type"`
	Actions []Action        `json:"actions"`
}

//go:embed config.json
var config embed.FS

func ReadConfigFromFile() Config {
	file, err := config.Open("config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	return config
}
