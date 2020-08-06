package main

import (
	"bufio"
	"log"
	"os"

	"github.com/valaymerick/mower/parser"
)

// ParseFile parses a mow configuration file and returns the corresponding Config
func ParseFile(name string) *parser.Config {
	f, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error opening config file: %v", err)
	}
	defer f.Close()

	p := parser.NewParser(*bufio.NewReader(f))
	return p.Parse()
}
