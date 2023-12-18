package main

import (
	"errors"
	"strings"
)

type DSLLine struct {
	Directive string
	Arguments []string
}

func parseDSLLine(line string) (DSLLine, error) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return DSLLine{}, errors.New("empty line")
	}

	directive := parts[0]
	arguments := parts[1:]

	return DSLLine{Directive: directive, Arguments: arguments}, nil
}

func parseKeyValue(line string) (string, string, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return "", "", errors.New("invalid format")
	}

	key := strings.TrimSpace(parts[0])
	value := strings.Trim(parts[1], " \";")

	return key, value, nil
}
