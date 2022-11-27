package util

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func InitConfigSpace(p string) error {
	// make config path if not provided
	if p == "" {
		d, err := os.UserConfigDir()
		if err != nil {
			return err
		}
		p = d
	}

	// resolve relative paths if given
	p, err := filepath.Abs(path.Join(p, ".aocfetch"))
	if err != nil {
		return err
	}

	err = makeIfNotExists(p)
	if err != nil {
		return err
	}

	return nil
}

func makeIfNotExists(p string) error {
	// if doesn't exist, make it w/ drwxr-xr-x
	if b, err := pathExists(p); !b {
		if err != nil {
			return err
		}

		fmt.Printf("%s does not exist - creating for the first time\n", p)
		err = os.Mkdir(p, 0755)
		if err != nil {
			return err
		}

		return err
	}

	return nil
}

func pathExists(p string) (bool, error) {
	if _, err := os.Stat(p); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}
