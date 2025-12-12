package main

import (
	"flag"
	"log/slog"
	"os"
	"strings"
)

func main() {
	test := flag.Bool("test", false, "Run test case or not")
	flag.Parse()

	var inputFile string
	if *test {
		inputFile = "./input.test"
	} else {
		inputFile = "./input"
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		slog.Error("Couldn't read file provided.", "path", path, "error", err)
		os.Exit(1)
	}

	content := strings.Split(string(bytes), "\n\n")
	var presents []string
}
