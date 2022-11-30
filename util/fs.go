package util

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func MakeIfNotExists(p string) error {
	// if doesn't exist, make it w/ drwxr-xr-x
	if b, err := PathExists(p); !b {
		if err != nil {
			return err
		}

		err = os.MkdirAll(p, 0755)
		if err != nil {
			return err
		}

		return err
	}

	return nil
}

func PathExists(p string) (bool, error) {
	if _, err := os.Stat(p); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

func GetDayFile(configDir string, year, day int) string {
	dir := GetConfigDir(configDir)
	p, err := filepath.Abs(path.Join(dir, strconv.Itoa(year), strconv.Itoa(day), "input.txt"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return p
}

func MakeDayIfNotExists(configDir string, year, day int) error {
	p := path.Dir(GetDayFile(configDir, year, day))
	return MakeIfNotExists(p)
}

func CreateFile(name string, data []byte) error {
	// create file w/ -rw-r--r--
	err := os.WriteFile(name, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func CopyFile(dest, configDir string, year, day int) error {
	p := GetDayFile(configDir, year, day)

	data, err := os.ReadFile(p)
	if err != nil {
		return err
	}

	err = os.WriteFile(dest, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
