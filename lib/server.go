package main

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

func parseDSLConfig(filePath string) (ServerConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return ServerConfig{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	config := ServerConfig{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line == "{" || line == "}" {
			continue
		}

		key, value, err := parseKeyValue(line)
		if err != nil {
			return ServerConfig{}, err
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

func parseKeyValue(line string) (string, string, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid format: %s", line)
	}

	key := strings.TrimSpace(parts[0])
	value := strings.Trim(parts[1], " \";")

	return key, value, nil
}
