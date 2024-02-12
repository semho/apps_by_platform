package config

import (
	"encoding/json"
	"os"
	"strings"
)

func Load(fileName string) (Configuration, error) {
	var data []byte
	var config Configuration
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	m := map[string]any{}
	err = decoder.Decode(&m)
	if err != nil {
		return nil, err
	}
	config = &DefaultConfig{configData: m}
	return config, err
}
