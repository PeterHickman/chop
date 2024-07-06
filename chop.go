package main

import (
	"bufio"
	"flag"
	"fmt"
	toolbox "github.com/PeterHickman/toolbox"
	"os"
	"strings"
)

var header string
var wanted string
var unwanted string

func dropdead(message string) {
	fmt.Println(message)
	os.Exit(3)
}

func reportMatch(lines []string) {
	text := strings.Join(lines, "\n")

	x := false
	if wanted != "" && strings.Contains(text, wanted) {
		x = true
	}
	if unwanted != "" && strings.Contains(text, unwanted) {
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

		if strings.Contains(line, header) {
			reportMatch(lines)
			lines = []string{}
		}

		lines = append(lines, line)
	}
	reportMatch(lines)

	readFile.Close()
}

func init() {
	var h = flag.String("header", "", "The string that indicates the start of the block")
	var w = flag.String("wanted", "", "The string that indicates that we want the block")
	var u = flag.String("unwanted", "", "The string that indicates that we do not want the block")

	flag.Parse()

	header = *h
	if header == "" {
		dropdead("--header is required\n")
	}

	wanted = *w
	unwanted = *u

	if wanted == "" && unwanted == "" {
		wanted = header
	}
}

func main() {
	for _, filename := range flag.Args() {
		if toolbox.FileExists(filename) {
			process(filename)
		} else {
			dropdead(fmt.Sprintf("[%s] is not a real file\n", filename))
		}
	}
}
