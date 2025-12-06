package main

import (
	"flag"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

func Zip[T ~[]E, E any](unzipped ...T) [][]E {
	if unzipped == nil {
		return nil
	}

	if len(unzipped) == 0 {
		return [][]E{}
	}

	cols := len(unzipped)
	rows := len(unzipped[0])
	for _, slice := range unzipped {
		if len(slice) < rows {
			rows = len(slice)
		}
	}

	zipped := make([][]E, rows)
	arr := make([]E, rows*len(unzipped))
	for i := 0; i < rows; i++ {
		row := arr[:cols:cols]
		arr = arr[cols:]
		for j := 0; j < cols; j++ {
			row[j] = unzipped[j][i]
		}
		zipped[i] = row
	}

	return zipped
}

func main() {
	test := flag.Bool("test", false, "Run test case or not")
	flag.Parse()

	var inputFile string
	if *test {
		inputFile = "./input.test"
	} else {
		inputFile = "./input"
	}

	lines := utils.ReadFile(inputFile)
}
