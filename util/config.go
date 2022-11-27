package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func GetConfigDir(configDir string) string {
	if configDir == "" {
		d, err := os.UserConfigDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configDir = d
	}

	// resolve relative paths
	configDir, err := filepath.Abs(path.Join(configDir, ".aocfetch"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return configDir
}

func InitConfigDir(configDir string) error {
	configDir = GetConfigDir(configDir)

	err := MakeIfNotExists(configDir)
	if err != nil {
		return err
	}

	return nil
}
