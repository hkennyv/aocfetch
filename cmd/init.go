package cmd

import (
	"aocfetch/util"

	"github.com/spf13/cobra"
)

var ConfigDir string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes aocfetch on your local machine",
	Long:  "Initializes aocfetch. Creates the configuration directory and initial config.",
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	util.InitConfigSpace(ConfigDir)
}

func init() {
	initCmd.Flags().StringVarP(&ConfigDir, "directory", "d", "",
		`Path to desired config directory, defaults to $HOME on Mac, %AppData%
on Windows, and $XDG_CONFIG_HOME on linux`)

	rootCmd.AddCommand(initCmd)
}
