package cmd 

import (
    "io/ioutil"
    "log"
    "gopkg.in/yaml.v2"
)

// Config to use
type Config struct {
    Bait      string `yaml:"bait"`
    Debug     bool   `yaml:"debug"`
    Key	      string `yaml:"key"`
    LogPath   string `yaml:"log"`
    Threshold int    `yaml:"threshold"`
}

// GetConfig of user
func (c *Config) GetConfig(config string) *Config {
    if config == "" {
        config = "config.yml"
    }
    yamlFile, err := ioutil.ReadFile(config)
    if err != nil {
        log.Printf("Error in %v ", err)
    }   
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }   
    return c
}
