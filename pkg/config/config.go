package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Config - This struct will hold configuration components.
type Config struct {
	Resolver     string   `json:"resolver"`
	ResolverPort int      `json:"resolverPort"`
	Domains      []string `json:"domains"`
	JumpHost     string   `json:"jumpHost"`
	SplitString  string   `json:"splitString"`
}

//GetConfig gets the configuration values for the api using the file in the supplied configPath.
func GetConfig(configPath string, debug bool) (Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("could not find the config file at path %s", configPath)
	}
	if debug {
		log.Println("Loading Configuration File: " + configPath)
	}
	return loadConfigFromFile(configPath)
}

//if the config loaded from the file errors, no defaults will be loaded and the app will exit.
func loadConfigFromFile(configPath string) (conf Config, err error) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Printf("Error opening config file: %v", err)
	} else {
		defer file.Close()

		err = json.NewDecoder(file).Decode(&conf)
		if err != nil {
			log.Printf("Error decoding config file: %v", err)
		}
	}

	return conf, err
}
