/*
 * hl
 *
 * Copy standard input to standard output, highlighting lines that match
 * pattern.
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/karrick/golf"
	"github.com/karrick/golfw"
)

func main() {
	var iowc io.WriteCloser = os.Stdout
	var buf []byte
	var pre, post string
	var prev int
	var useBuffer bool

	optAnsi := golf.String("ansi", "bold", "highlight ansi")
	optBuffer := golf.Bool("buffer", false, "buffers even when writing to TTY")
	optNoBuffer := golf.Bool("no-buffer", false, "prevent buffering when not writing to TTY")
	golf.Parse()

	programName, err := os.Executable()
	if err != nil {
		programName = os.Args[0]
	}
	programName = filepath.Base(programName)

	if golf.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "%s: USAGE: %s [--ansi STRING] [--buffer | --no-buffer] regex\n", programName, programName)
		os.Exit(1)
	}

	patternRE, err := regexp.Compile(golf.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: invalid regex pattern: %s\n", programName, err)
		os.Exit(1)
	}

	if len(*optAnsi) > 0 {
		pre, post, err = ansiCodes(*optAnsi)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: cannot process option argument for -ansi : %s\n", programName, err)
			os.Exit(1)
		}
	}

	if *optBuffer {
		useBuffer = true
	} else if !*optNoBuffer {
		if fi, err := os.Stdout.Stat(); err == nil && fi.Mode()&os.ModeCharDevice == 0 {
			// When not writing to terminal, do not need to write after
			// every newline.
			useBuffer = true
			if iowc, err = golfw.NewWriteCloser(os.Stdout, 512); err != nil {
				// When cannot allocate new write closer, fall back to
				// writing to stdout.
				iowc = os.Stdout
				useBuffer = false
			}
		}
	}
	// fmt.Fprintf(os.Stderr, "useBuffer: %t\n", useBuffer)
	// os.Exit(42)

	buf = append(buf, post...) // very first print should set normal intensity

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
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

		if _, err = iowc.Write(buf); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", programName, err)
			os.Exit(3)
		}

		buf = buf[:0] // reset buffer for next line
		prev = 0
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", programName, err)
		os.Exit(2)
	}

	if useBuffer {
		if err := iowc.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", programName, err)
			os.Exit(3)
		}
	}
}

func ansiCodes(option string) (string, string, error) {
	premap := make(map[int]struct{})
	postmap := make(map[int]struct{})
	var pres, posts []string
	var pre, post int

	for _, arg := range strings.Split(strings.ToLower(option), ",") {
		switch arg {
		case "bold":
			pre = 1
			post = 22
		case "dim", "faint":
			pre = 2
			post = 22
		case "italic":
			pre = 3
			post = 23
		case "underline", "underscore":
			pre = 4
			post = 24
		case "blinking":
			pre = 5
			post = 25
		case "inverse", "reverse":
			pre = 7
			post = 27
		case "hidden", "invisible":
			pre = 8
			post = 28
		case "strikethrough":
			pre = 9
			post = 29
		case "black":
			pre = 30
			post = 0
		case "red":
			pre = 31
			post = 0
		case "green":
			pre = 32
			post = 0
		case "yellow":
			pre = 33
			post = 0
		case "blue":
			pre = 34
			post = 0
		case "magenta":
			pre = 35
			post = 0
		case "cyan":
			pre = 36
			post = 0
		case "white":
			pre = 37
			post = 0
		default:
			return "", "", fmt.Errorf("cannot recognize argument: %q", arg)
		}

		if _, ok := premap[pre]; !ok {
			pres = append(pres, strconv.Itoa(pre))
			premap[pre] = struct{}{}
		}

		if _, ok := postmap[post]; !ok {
			posts = append(posts, strconv.Itoa(post))
			postmap[post] = struct{}{}
		}
	}

	return ansiFromCodes(pres), ansiFromCodes(posts), nil
}

// ansiFromCodes returns an ANSI control sequence corresponding to the list of
// sequences provided.
func ansiFromCodes(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return "\x1b[" + strings.Join(strs, ";") + "m"
}
