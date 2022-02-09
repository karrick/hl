---
date: February 3, 2022
section: 1
title: HL
---

# NAME

**hl** - read and highlight regular expression matches

# SYNOPSIS

**hl pattern**  

**hl** \[**--ansi** **attribute**\] \[**--buffer**\] **pattern**  

**hl** \[**--ansi** **attribute**\] \[**--no-buffer**\] **pattern**

# DESCRIPTION

The **hl** utility copies standard input to standard output,
highlighting lines that match **pattern**, which must be a valid regular
expression. The optional command line arguments are as follows:

**--ansi** **attribute**  
Highlights pattern matches using the ANSI code corresponding to
**attribute**.

By default, **hl** uses ANSI codes to embolden text to highlight match
text, and non-bold text for all other text. Other ANSI codes are
supported as shown in the below table. Several rows contain multiple
attributes, indicating they are intended synonyms.

> |                         |           |
> |:------------------------|:----------|
> |                         |           |
> | attribute               | ANSI code |
> |                         |           |
> | ----------------------- | --------- |
> |                         |           |
> | bold                    | ESC\[1m   |
> |                         |           |
> | dim\|faint              | ESC\[2m   |
> |                         |           |
> | italic                  | ESC\[3m   |
> |                         |           |
> | underline\|underscore   | ESC\[4m   |
> |                         |           |
> | blinking                | ESC\[5m   |
> |                         |           |
> | reverse\|inverse        | ESC\[7m   |
> |                         |           |
> | hidden\|invisible       | ESC\[8m   |
> |                         |           |
> | strikethrough           | ESC\[9m   |
> |                         |           |
> | black                   | ESC\[30m  |
> |                         |           |
> | red                     | ESC\[31m  |
> |                         |           |
> | green                   | ESC\[32m  |
> |                         |           |
> | yellow                  | ESC\[33m  |
> |                         |           |
> | blue                    | ESC\[34m  |
> |                         |           |
> | magenta                 | ESC\[35m  |
> |                         |           |
> | cyan                    | ESC\[36m  |
> |                         |           |
> | white                   | ESC\[37m  |

**NOTE**: Not all ANSI codes work on every terminal.

**--buffer**  
Buffers its output even when it detects that it is writing to a
character device.

By default when **hl** detects that its standard output is directed to a
character device, like a **TTY**, it will write each line of output
immediately after processing it. However, if **hl** detects that its
standard output is not directed to a character device, for instance when
it is being piped into another process, it will use a small memory
buffer to reduce the number of system write calls it makes. However,
when given the **--buffer** command line flag, **hl** will buffer its
output even when it detects that it is writing to a character device.

**--no-buffer**  
Does not buffer its output even when its output is not a character
device. Use this when piping the output of **hl** to another process,
and each complete line should be flushed immedidately after it is
processed.

# EXIT VALUES

**0**  
Success

**1**  
Syntax or usage error

**2**  
Error reading input

**3**  
Error writing output

# EXAMPLES

**Example 1: Highlighting a text pattern from a stream**  
This example tails a file and highlights any instances of the word
**failure** using a bold typeface:

<!-- -->


    # tail -f /var/log/messages | hl failure

**Example 2: Highlighting multiple patterns from a stream**  
**hl** invocations may be chained together to highlight multiple
patterns using different ANSI codes. When doing so, it is common to turn
off buffering in every instance of **hl** that pipes its output to
another program. This example will tail a file and highlight **BAD**
with red, and **OK** with green.

<!-- -->


    # tail -f /var/log/messages | hl --ansi red --no-buffer BAD | hl --ansi green OK

# AUTHORS

Karrick McDermott **https://linkedin.com/in/karrick**
