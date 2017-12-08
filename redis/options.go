package redis

import (
	"encoding/json"
	"errors"
	"strconv"
)

// ConfigRedis Configuration structure of Redis instance
type ConfigRedis struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

// Config Configuration structure pulled from JSON file
type Config struct {
	Master string `json:"master"`

	Redis ConfigRedis `json:"redis"`

	Sentinel struct {
		Hostname string                 `json:"hostname"`
		Port     int                    `json:"port"`
		Config   map[string]interface{} `json:"config"`
	} `json:"sentinel"`
}

// MarshalJSON Converts our password to a blank on output
func (r *ConfigRedis) MarshalJSON() ([]byte, error) {
	type Alias ConfigRedis

	return json.Marshal(&struct {
		Password string `json:"password,omitempty"`
		*Alias
	}{
		Password: "",
		Alias:    (*Alias)(r),
	})
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
