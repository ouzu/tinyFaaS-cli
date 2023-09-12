package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v2"
)

type TFConfig struct {
	ConfigPort int
	COAPPort   int
	Host       string
}

func LoadConfig(c *cli.Context) *TFConfig {
	configFile := c.String("config")
	if configFile != "" {
		// Read the config file into config variable
		file, err := os.ReadFile(configFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var config TFConfig

		if err := toml.Unmarshal(file, &config); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		return &config
	}

	return nil
}

func GetManagerUrl(c *cli.Context) string {
	url := "http://localhost:8080"

	config := LoadConfig(c)

	if config != nil {
		if config.Host != "" {
			url = fmt.Sprintf("http://%s:%d", config.Host, config.ConfigPort)
		}
	}

	return url
}

func GetCoapUrl(c *cli.Context) string {
	url := "coap://localhost:5683"

	config := LoadConfig(c)

	if config != nil {
		if config.Host != "" {
			url = fmt.Sprintf("%s:%d", config.Host, config.COAPPort)
		}
	}

	return url
}
