package config

func (c *DefaultConfig) GetStringDefault(name, defVal string) string {
	result, ok := c.GetString(name)
	if !ok {
		result = defVal
	}
	return result
}

func (c *DefaultConfig) GetIntDefault(name string, defVal int) int {
	result, ok := c.GetInt(name)
	if !ok {
		result = defVal
	}
	return result
}

func (c *DefaultConfig) GetBoolDefault(name string, defVal bool) bool {
	result, ok := c.GetBool(name)
	if !ok {
		result = defVal
	}
	return result
}

func (c *DefaultConfig) GetFloatDefault(name string, defVal float64) float64 {
	result, ok := c.GetFloat(name)
	if !ok {
		result = defVal
	}
	return result
}
