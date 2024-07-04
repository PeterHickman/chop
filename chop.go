package main

// TODO: need to use the flag package

import (
	"bufio"
	"fmt"
	toolbox "github.com/PeterHickman/toolbox"
	"os"
	"strings"
)

var options = make(map[string]string)
var hasWanted bool
var hasUnwanted bool
var files = []string{}

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
	if hasWanted && strings.Contains(text, options["wanted"]) {
		x = true
	}
	if hasUnwanted && strings.Contains(text, options["unwanted"]) {
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
		dropdead("Unable to read " + filename)
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
			if toolbox.FileExists(k) {
				files = append(files, k)
			} else {
				dropdead(fmt.Sprintf("[%s] is not a real file\n", k))
			}
		}
	}

	hasWanted = hasKey(options, "wanted")
	hasUnwanted = hasKey(options, "unwanted")

	if hasKey(options, "header") {
		if hasWanted || hasUnwanted {
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
