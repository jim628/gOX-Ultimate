package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type (
	Offset struct {
		X int `toml:"x"`
		Y int `toml:"y"`
	}

	Board struct {
		PathS  string `toml:"pathS"`
		PathG  string `toml:"pathG"`
		Offset Offset `toml:"offset"`
	}

	Config struct {
		Board Board `toml:"board"`
	}
)

func loadConfig(path string) (*Config, error) {
	f := "../example-config/config.toml"
	if _, err := os.Stat(f); err != nil {
		f = "example-config/config.toml"
	}

	var cfg Config
	_, err := toml.DecodeFile(f, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg, nil
}

func loadBoard(f *bufio.Scanner, start string, end string) (string, error) {
	var sb strings.Builder
	for f.Scan() {
		ln := f.Text()
		if strings.Contains(ln, end) {
			return sb.String(), nil
		}
		ln = strings.Replace(ln, ";", "", -1)
		sb.WriteString(ln)
		sb.WriteByte('\n')
	}

	return sb.String(), fmt.Errorf("couldn't parse %s", start)
}
