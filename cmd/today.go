package cmd

import (
	"aocfetch/util"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var (
	CopyDir  string
	CopyName string
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "fetch today's puzzle",
	Long:  "fetch today's puzzle, defaults to current working directory",
	Run:   runToday,
}

func init() {
	todayCmd.Flags().StringVar(&CopyDir, "dest", ".", "directory to fetch today's input to. defaults to the current day (UTC-5 AOC time)")
	todayCmd.Flags().StringVar(&CopyName, "name", "input.txt", "name to save input as. defaults to 'input.txt'")
}

func runToday(cmd *cobra.Command, args []string) {
	loc := time.FixedZone("UTC-5", -5*60*60)
	now := time.Now()
	t := now.In(loc)

	err := util.ValidateDate(t.Year(), t.Day())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = util.InitConfigDir(ConfigDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = syncDay(t.Year(), t.Day())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dest := filepath.Join(CopyDir, CopyName)

	if b, _ := util.PathExists(dest); b {
		fmt.Printf("- %s already exists\n", dest)
		return
	}

	err = util.CopyFile(dest, ConfigDir, t.Year(), t.Day())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("✓ %s copied from cache\n", dest)
}
