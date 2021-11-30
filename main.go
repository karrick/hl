/*
 * hl
 *
 * Copy standard input to standard output, highlighting lines that match
 * pattern.
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	optAnsi := flag.String("ansi", "bold", "highlight ansi")
	flag.Parse()

	programName, err := os.Executable()
	if err != nil {
		programName = os.Args[0]
	}
	programName = filepath.Base(programName)

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "%s: USAGE: %s [-ansi STRING] regex\n", programName, programName)
		os.Exit(2)
	}

	patternRE, err := regexp.Compile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: invalid regex pattern: %s\n", programName, err)
		os.Exit(2)
	}

	var buf []byte
	var pre, post string

	switch strings.ToLower(*optAnsi) {
	case "bold":
		pre = "\033[1m"
		post = "\033[22m"
	case "dim", "faint":
		pre = "\033[2m"
		post = "\033[22m"
	case "italic":
		pre = "\033[3m"
		post = "\033[23m"
	case "underline", "underscore":
		pre = "\033[4m"
		post = "\033[24m"
	case "blinking":
		pre = "\033[5m"
		post = "\033[25m"
	case "inverse", "reverse":
		pre = "\033[7m"
		post = "\033[27m"
	case "hidden", "invisible":
		pre = "\033[8m"
		post = "\033[28m"
	case "strikethrough":
		pre = "\033[9m"
		post = "\033[29m"
	case "black":
		pre = "\033[30m"
		post = "\033[0m"
	case "red":
		pre = "\033[31m"
		post = "\033[0m"
	case "green":
		pre = "\033[32m"
		post = "\033[0m"
	case "yellow":
		pre = "\033[33m"
		post = "\033[0m"
	case "blue":
		pre = "\033[34m"
		post = "\033[0m"
	case "magenta":
		pre = "\033[35m"
		post = "\033[0m"
	case "cyan":
		pre = "\033[36m"
		post = "\033[0m"
	case "white":
		pre = "\033[37m"
		post = "\033[0m"
	default:
		fmt.Fprintf(os.Stderr, "%s: cannot recognize -ansi option argument: %q\n", programName, *optAnsi)
		os.Exit(2)
	}

	buf = append(buf, post...) // very first print should set normal intensity

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var prev int

		matches := patternRE.FindAllStringSubmatchIndex(line, -1)

		for _, tuple := range matches {
			buf = append(buf, line[prev:tuple[0]]...)     // print text before match
			buf = append(buf, pre...)                     // bold intensity
			buf = append(buf, line[tuple[0]:tuple[1]]...) // print match
			buf = append(buf, post...)                    // normal intensity
			prev = tuple[1]
		}

		buf = append(buf, line[prev:]...) // print remaining text after final match
		buf = append(buf, '\n')

		if _, err = os.Stdout.Write(buf); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", programName, err)
			os.Exit(1)
		}

		buf = buf[:0] // reset buffer for next line
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", programName, err)
		os.Exit(1)
	}
}
