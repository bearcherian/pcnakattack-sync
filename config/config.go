package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YouTubeConfig struct {
	ApiKey string `yaml:"apiKey"`
}

type TwitterConfig struct {
	ConsumerKey    string `yaml:"consumerKey"`
	ConsumerSecret string `yaml:"consumerSecret"`
	Token          string `yaml:"token"`
	Secret         string `yaml:"secret"`
}

type DatabaseConfig struct {
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Tags    []string      `yaml:"tags"`
	Database DatabaseConfig `yaml:"database"`
	Twitter TwitterConfig `yaml:"twitter"`
	YouTube YouTubeConfig `yaml:"youtube"`

}

func GetConfig() Config {
	cfg := Config{}

	// read yaml
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(yamlFile, &cfg)

	return cfg
}
