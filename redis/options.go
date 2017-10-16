package redis

import (
	"errors"
	"strconv"
)

// Config Configuration structure pulled from JSON file
type Config struct {
	Master string `json:"master"`

	Redis struct {
		Hostname string `json:"hostname"`
		Port     int    `json:"port"`
		Password string `json:"password"`
	} `json:"redis"`

	Sentinel struct {
		Hostname string                 `json:"hostname"`
		Port     int                    `json:"port"`
		Config   map[string]interface{} `json:"config"`
	} `json:"sentinel"`
}

// ConvertToString Takes an interface value and ensures it's a string
func (c *Config) ConvertToString(data interface{}) (string, error) {
	if str, ok := data.(string); ok {
		return str, nil
	}

	if float, ok := data.(float64); ok {
		return strconv.FormatFloat(float, 'f', -1, 64), nil
	}

	return "", errors.New("String conversion of config value failed")
}
