package cmd

import (
	"aocfetch/util"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// args
var (
	Year int
	Day  int
)

var syncCmd = &cobra.Command{
	Use:   "sync [year] [day]",
	Short: "Syncs Advent of Code input (defaults to the current year)",
	Long:  "Sync Advent of Code input for a given year, defaults to the current year",
	Args:  handleArgs,
	Run:   runSync,
}

func runSync(cmd *cobra.Command, args []string) {
	now := time.Now()

	// ensure configdir has been initialized
	err := util.InitConfigDir(ConfigDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// default case - sync all current year
	if Year == 0 && Day == 0 {
		fmt.Printf("Syncing current year - %d\n", now.Year())
		err := syncYear(now.Year())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// if only year is given - sync the year
	} else if Year != 0 && Day == 0 {
		fmt.Printf("Syncing %d\n", Year)
		err := syncYear(Year)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// if only day is given - sync this day of the current year
	} else if Day != 0 && Year == 0 {
		fmt.Printf("Syncing %d day %d\n", now.Year(), Day)
		err := syncDay(now.Year(), Day)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// otherwise both year and day are given - sync specific year/date combo
	} else {
		fmt.Printf("Syncing %d day %d\n", Year, Day)
		err := syncDay(Year, Day)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func handleArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MaximumNArgs(2)(cmd, args); err != nil {
		return err
	}

	if len(args) == 2 {
		year, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Year needs to be an integer:", err.Error())
			return err
		}
		day, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Day needs to be an integer:", err.Error())
			return err
		}

		err = util.ValidateDate(year, day)
		if err != nil {
			return err
		}

		Year = year
		Day = day
	}

	if len(args) == 1 {
		year, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Year needs to be an integer:", err.Error())
			return err
		}

		err = util.ValidateYear(year)
		if err != nil {
			return err
		}

		Year = year
	}

	return nil
}

func syncYear(year int) error {
	if !util.YearIsStarted(year) {
		msg := fmt.Sprintf("✗ you're early! aoc %d has not started yet!", year)
		return errors.New(msg)
	}

	for day := 1; day <= 25; day++ {
		err := syncDay(year, day)
		if err != nil {
			return err
		}
	}

	return nil
}

func syncDay(year, day int) error {
	// very important to cache to disk to not upset the topaz!
	p := util.GetDayFile(ConfigDir, year, day)
	if e, _ := util.PathExists(p); e {
		fmt.Printf("- AOC %d day %d already synced\n", year, day)
		return nil
	}

	err := util.MakeDayIfNotExists(ConfigDir, year, day)
	if err != nil {
		return err
	}

	b, err := util.FetchDay(year, day)
	if err != nil {
		return err
	}

	err = util.CreateFile(p, b)
	if err != nil {
		return err
	}

	fmt.Printf("✓ Synced AOC %d day %d\n", year, day)

	return nil
}
