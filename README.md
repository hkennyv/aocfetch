# Advent of Code Fetch

## Overview

A CLI tool used to fetch [Advent of Code](https://adventofcode.com/about)
input by year/day/part.

## Features

- [x] Fetch input from [Advent of Code](https://adventofcode.com/about)
- [x] Configurable config path
- [x] Cache input in `.aocfetch`
- [x] Sync command to sync all inputs given a year and/or day
- [ ] Clean command to clear year and/or day input(s)
- [ ] Setup Github Actions for CI + releases
- [ ] Copy input from cache to destination
- [ ] Fetch session from browser - maybe (w/ [kooky](https://github.com/zellyn/kooky)?)
- [ ] Setup godoc

## Usage

### Sync

Use the `aocfetch sync` command to download/sync inputs. It caches the inputs
to disk to avoid extra requests to the AOC server.

Inputs are synced to the following locations depending on your system:

- **Windows**: `%AppData%\.aocfetch`
- **MacOS**: `$HOME/Library/Application\ Support\.aocfetch`
- **Linux**: `$XDG_CONFIG_HOME/.aocfetch` or `$HOME/.config/.aocfetch`


- By default, it syncs the current year's input, you can use this command each
  day to pull the latest day's input and it will skip all previously synced 
  ones:

```
$ ./aocfetch sync
```

- To sync a full year's inputs (e.g sync all 2015 inputs), use:

```
$ ./aocfetch sync 2015
```

- To sync a single day's input (e.g. sync 2015 day 2 input), use:

```
$ ./aocfetch sync 2015 2
```

For more info, use the `--help` flag.

```
Sync Advent of Code input for a given year, defaults to the current year

Usage:
  aocfetch sync [year] [day] [flags]

Flags:
  -h, --help   help for sync
```

### More help

Other commands are still WIP, but you can use the `--help` flag anytime to
get more help.

```
$ ./aocfetch --help
Fetches Advent of Code inputs by year and day from your CLI! Downloads today's input by default (if valid)

Usage:
  aocfetch [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  sync        Syncs Advent of Code input (defaults to the current year)

Flags:
  -d, --directory string   Path to desired config directory, defaults to $HOME on Mac, %AppData%
                           on Windows, and $XDG_CONFIG_HOME on linux
  -h, --help               help for aocfetch
  -v, --verbose            Enables verbose logging

Use "aocfetch [command] --help" for more information about a command.)
```

## License

[MIT LICENSE](./LICENSE)
