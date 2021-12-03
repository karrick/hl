# hl

Read and highlight regular expression matches.

## Description

Copy standard input to standard output, highlighting lines that match
pattern.

## Usage

The below example tails a file, copying lines to the console, but
highlighting any substrings that match the term `failure`.

```Bash
tail -f /var/log/messages | hl failure
```

### Changing the -ansi command line flag

By default, `hl` uses ANSI codes to embolden text to highlight a
match, and non-bold text for all other text. Other ANSI codes are
supported as shown in the below table.

NOTE: Not all ANSI codes work on every terminal.

| -ansi option argument | effect  |
|-----------------------|---------|
| bold                  | ESC[1m  |
| dim, faint            | ESC[2m  |
| italic                | ESC[3m  |
| underline, underscore | ESC[4m  |
| blinking              | ESC[5m  |
| reverse, inverse      | ESC[7m  |
| hidden, invisible     | ESC[8m  |
| strikethrough         | ESC[9m  |
| black                 | ESC[30m |
| red                   | ESC[31m |
| green                 | ESC[32m |
| yellow                | ESC[33m |
| blue                  | ESC[34m |
| magenta               | ESC[35m |
| cyan                  | ESC[36m |
| white                 | ESC[37m |

#### Combining different ANSI effects

```Bash
tail -f /var/log/messages | hl -ansi reverse,red,bold failure
```

## Installation

If you don't have the Go programming language installed, then you'll
need to install it from your package manager, or download and install
it from [https://golang.org/dl](https://golang.org/dl).

Once you have Go installed:

    $ go get github.com/karrick/hl
