package apiserver

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	Port     string `json:"port"`
	Database string `json:"database"`
}

func LoadConfig() *Config {
	// Check if environment variable CONFIG_PATH exist
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Println("CONFIG_PATH variable is not set. Will use default path: config/server.json")
		configPath = "config/server.json"
	}

	// Check if file exists
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("Config %s does not exist", configPath)
	}

	// Try open file
	f, err := os.Open(configPath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)

	var data string

	// Read file and save each line into data variable
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		data += line
	}

	//Initialize a Config struct
	cfg := &Config{}

	// Unmarshal data into Config struct
	if err := json.Unmarshal([]byte(data), cfg); err != nil {
		log.Fatalf("ERROR: %s\n", err)
		return nil
	}

	return cfg
}
