package dummy

type Config struct {
}

func NewDummyConfig() *Config {
	return &Config{}
}

func (c *Config) AuthDatabaseDSN() string {
	return ""
}

func (c *Config) AuthPort() string {
	return ""
}
