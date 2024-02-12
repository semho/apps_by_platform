package config

import "strings"

type DefaultConfig struct {
	configData map[string]any
}

func (c *DefaultConfig) get(name string) (any, bool) {
	data := c.configData
	var result any
	var found bool
	for _, key := range strings.Split(name, ":") {
		result, found = data[key]
		newSection, ok := result.(map[string]any)
		if !ok || !found {
			return result, found
		}
		data = newSection
	}
	return result, found
}

func (c *DefaultConfig) GetSection(name string) (Configuration, bool) {
	value, found := c.get(name)
	var section Configuration
	if found {
		if sectionData, ok := value.(map[string]any); ok {
			section = &DefaultConfig{configData: sectionData}
		}
	}
	return section, found
}

func (c *DefaultConfig) GetString(name string) (string, bool) {
	value, found := c.get(name)
	var result string
	if found {
		result = value.(string)
	}
	return result, found
}

func (c *DefaultConfig) GetInt(name string) (int, bool) {
	value, found := c.get(name)
	var result int
	if found {
		result = int(value.(float64))
	}
	return result, found
}

func (c *DefaultConfig) GetBool(name string) (bool, bool) {
	value, found := c.get(name)
	var result bool
	if found {
		result = value.(bool)
	}
	return result, found
}

func (c *DefaultConfig) GetFloat(name string) (float64, bool) {
	value, found := c.get(name)
	var result float64
	if found {
		result = value.(float64)
	}
	return result, found
}
