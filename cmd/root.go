package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"aocfetch/util"

	"github.com/spf13/cobra"
)

// flags
var Year int
var Day int

var rootCmd = &cobra.Command{
	Use:   "aocfetch",
	Short: "aocfetch - Advent of Code input fetcher",
	Long:  "Fetches Advent of Code inputs by year and day from your CLI! Downloads today's input by default (if valid)",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	err := validateFlags(Year, Day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := util.FetchInput(Year, Day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Stdout.Write(body)
}

func validateFlags(year, day int) error {
	now := time.Now()
	s := ""

	if year < 2015 || year > now.Year() {
		s += "Advent of Code started in 2015, so year must be between 2015 and the current year inclusive.\n"
	}

	if day < 1 || day > 25 {
		s += "Advent of Code only runs in December up until the 25th, so days must be 1-25 inclusive.\n"
	}

	if s != "" {
		return errors.New("Errors validating flags:\n" + s)
	}

	return nil
}

func Main() {
	now := time.Now()

	rootCmd.Flags().IntVarP(&Year, "year", "y", now.Year(), "The AOC calendar year, defaults to current year.")
	rootCmd.Flags().IntVarP(&Day, "day", "d", now.Day(), "The desired day to fetch input for. Defaults to today, valid range is [1-25].")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
