package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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

type AppConfig struct {
	Tags     []string       `yaml:"tags"`
	Database DatabaseConfig `yaml:"database"`
	Twitter  TwitterConfig  `yaml:"twitter"`
	YouTube  YouTubeConfig  `yaml:"youtube"`
	loaded   bool
}

var cfg AppConfig

func GetConfig() AppConfig {

	if !cfg.loaded {
		cfg = AppConfig{}

		// read yaml
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}

		yaml.Unmarshal(yamlFile, &cfg)
		cfg.loaded = true
		log.Println("Config loaded")
		log.Printf("Tags: %v", cfg.Tags)
	}

	return cfg
}
