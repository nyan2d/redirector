package config

import (
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Host struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

type Config struct {
	Address string `yaml:"address"`
	Hosts   []Host `yaml:"hosts"`
}

func Read(r io.Reader) (*Config, error) {
	var conf Config
	if err := yaml.NewDecoder(r).Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func ReadFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Read(file)
}

func (cf *Config) HostsMap() map[string]string {
	hostsmap := map[string]string{}
	for _, host := range cf.Hosts {
		hostsmap[strings.TrimSpace(host.Source)] = strings.TrimSpace(host.Destination)
	}
	return hostsmap
}
