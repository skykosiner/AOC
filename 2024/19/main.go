package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var rules, lines []string

	bytes, _ := os.ReadFile("./input")
	str := strings.Split(string(bytes), "\n\n")
	lines = strings.Split(str[1], "\n")

	for _, itm := range strings.Split(str[0], ", ") {
		rules = append(rules, itm)
	}

	regexStr := fmt.Sprintf(`^(?:%s)+$`, strings.Join(rules, "|"))

	r := regexp.MustCompile(regexStr)

	sum := 0
	for _, word := range lines {
		if r.MatchString(word) {
			sum++
		}
	}

	fmt.Println(sum)
}
