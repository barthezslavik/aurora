package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ServerConfig struct {
	Host        string
	Port        int
	StaticFiles string
}

func ParseDSLConfig(filePath string) (ServerConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return ServerConfig{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	config := ServerConfig{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line == "{" || line == "}" || strings.HasPrefix(line, "Server") {
			continue
		}

		key, value, err := ParseKeyValue(line)
		if err != nil {
			return ServerConfig{}, fmt.Errorf("invalid configuration line: %w", err)
		}

		switch key {
		case "Host":
			config.Host = value
		case "Port":
			port, err := strconv.Atoi(value)
			if err != nil {
				return ServerConfig{}, fmt.Errorf("invalid Port value: %w", err)
			}
			config.Port = port
		case "StaticFiles":
			config.StaticFiles = value
		}
	}

	return config, scanner.Err()
}
