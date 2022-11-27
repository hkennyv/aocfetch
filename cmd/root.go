package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// flags
var Verbose bool
var ConfigDir string

var rootCmd = &cobra.Command{
	Use:   "aocfetch",
	Short: "aocfetch - Advent of Code input fetcher",
	Long:  "Fetches Advent of Code inputs by year and day from your CLI! Downloads today's input by default (if valid)",
}

func Main() {
	rootCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Enables verbose logging")
	rootCmd.Flags().StringVarP(&ConfigDir, "directory", "d", "",
		`Path to desired config directory, defaults to $HOME on Mac, %AppData%
on Windows, and $XDG_CONFIG_HOME on linux`)
	rootCmd.AddCommand(syncCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
