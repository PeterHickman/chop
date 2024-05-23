package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var options = make(map[string]string)
var has_wanted bool
var has_unwanted bool
var files = []string{}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func hasKey(h map[string]string, k string) bool {
	_, err := h[k]
	return err
}

func dropdead(message string) {
	fmt.Println(message)
	os.Exit(3)
}

func reportMatch(lines []string) {
	text := strings.Join(lines, "\n")

	x := false
	if has_wanted && strings.Contains(text, options["wanted"]) {
		x = true
	}
	if has_unwanted && strings.Contains(text, options["unwanted"]) {
		x = false
	}

	// return if text == ''

	if x {
		fmt.Println(text)
		fmt.Println()
	}
}

func process(filename string) {
	readFile, err := os.Open(filename)

	if err != nil {
		dropdead(fmt.Sprintf("Unable to read %s", filename))
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lines := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.Contains(line, options["header"]) {
			reportMatch(lines)
			lines = []string{}
		}

		lines = append(lines, line)
	}
	reportMatch(lines)

	readFile.Close()
}

func opts() {
	for i := 0; i < len(os.Args[1:]); i++ {
		k := os.Args[(1 + i)]

		if strings.HasPrefix(k, "--") {
			if hasKey(options, k) {
				dropdead(fmt.Sprintf("--%s has already been used\n", k))
			} else {
				i++
				options[k[2:]] = os.Args[(1 + i)]
			}
		} else {
			if fileExists(k) {
				files = append(files, k)
			} else {
				dropdead(fmt.Sprintf("[%s] is not a real file\n", k))
			}
		}
	}

	has_wanted = hasKey(options, "wanted")
	has_unwanted = hasKey(options, "unwanted")

	if hasKey(options, "header") {
		if hasKey(options, "wanted") || hasKey(options, "unwanted") {
			// This is good
		} else {
			options["wanted"] = options["header"]
		}
	} else {
		dropdead("--header is required\n")
	}
}

func main() {
	opts()

	for _, filename := range files {
		process(filename)
	}
}
