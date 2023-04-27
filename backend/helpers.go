package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readEnv(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot read env file: %s", path)
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line_split := strings.Split(line, "=")
		os.Setenv(line_split[0], line_split[1])
	}
}
