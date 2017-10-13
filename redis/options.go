package redis

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
