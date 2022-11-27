# Advent of Code Fetch

## Overview

A CLI tool used to fetch [Advent of Code](https://adventofcode.com/about)
input by year/day/part.

## Usage

```
$ ./aocfetch --help 
Fetches Advent of Code inputs by year and day from your CLI! Downloads today's input by default (if valid)

Usage:
  aocfetch [flags]

Flags:
  -d, --day int    The desired day to fetch input for. Defaults to today, valid range is [1-25]. (default 26)
  -h, --help       help for aocfetch
  -y, --year int   The AOC calendar year, defaults to current year. (default 2022)
```

## License

[MIT LICENSE](./LICENSE)
