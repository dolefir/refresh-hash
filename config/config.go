package config

import (
	"errors"
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Main defines the properties of the application configuration.
type Main struct {
	APIServer APIServer `yaml:"api-server"`
	Ticker    Ticker    `yaml:"ticker"`
	Logger    Logger    `yaml:"logger"`
}

// APIServer defines API server configuration.
type APIServer struct {
	HTTP HTTP `yaml:"http"`
}

// HTTP defines HTTP section of the API server configuration.
type HTTP struct {
	ListenAddr      string        `yaml:"listen-address"`
	GracefulTimeout time.Duration `yaml:"graceful-timeout"`
	Network         string        `yaml:"network"`
	AddrGrpc        string        `yaml:"address-grpc"`
}

// Ticker defines Ticker section of the API server configuration.
type Ticker struct {
	Timer   time.Duration `yaml:"timer"`
	Timeout time.Duration `yaml:"time-out"`
}

// Logger defines logger section of the API server configuration.
type Logger struct {
	Mode      string `yaml:"mode"`
	LogFormat string `yaml:"log-format"`
	LogLevel  string `yaml:"log-level"`
}

// NewConfig returns config environment reads file from config.yaml.
func NewConfig(configPath string) *Main {
	cfg := &Main{}
	log.Printf("configPath := %s", configPath)
	if err := readConfigFile(configPath, cfg); err != nil {
		log.Printf("read config file error %s", err)
	}

	return cfg
}

func readConfigFile(name string, cfg interface{}) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return errors.New("read config file error")
	}

	conf, err := os.Open(name)
	if err != nil {
		log.Printf("open file for reading %s: %v", name, err)
		return err
	}

	data, err := io.ReadAll(conf)
	if err != nil {
		log.Printf("unable to read the %s: %v", name, err)
		return err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		log.Printf("unmarshal %s: %v", name, err)
		return err
	}

	return nil
}
