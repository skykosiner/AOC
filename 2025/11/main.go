package main

import (
    "flag"

"github.com/skykosiner/AOC/pkg/utils"
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

lines := utils.ReadFile(inputFile)
}
