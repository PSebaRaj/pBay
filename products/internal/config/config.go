package config

import (
	"flag"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	HTTPServer struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`

	DBConn struct {
		Dialect   string `yaml:"dialect"`
		Address   string `yaml:"address"`
		Database  string `yaml:"database"`
		Port      string `yaml:"port"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Protocol  string `yaml:"protocol"`
		ParseTime bool   `yaml:"parse_time"`
	} `yaml:"db"`

	CacheConn struct {
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		DBNumber int    `yaml:"db"`
	} `yaml:"redis"`

	CurrencyGRPCConn struct {
		Address string `yaml:"address"`
		Port    string `yaml:"port"`
	} `yaml:"currency_grpc"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file, that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./../../configs/config.yaml", "path to config file")

	flag.Parse()

	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}
