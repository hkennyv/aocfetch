package util

import (
	"errors"
	"time"
)

func ValidateDate(year, day int) error {
	s := ""

	err := ValidateYear(year)
	if err != nil {
		s += "\n" + err.Error()
	}

	err = ValidateDay(day)
	if err != nil {
		s += "\n" + err.Error()
	}

	if s != "" {
		return errors.New("error(s) validating date:" + s)
	}

	return nil
}

func ValidateYear(year int) error {
	now := time.Now()

	if year < 2015 || year > now.Year() {
		return errors.New("advent of code started in 2015, so year must be between 2015 and the current year inclusive")
	}

	return nil
}

func ValidateDay(day int) error {
	if day < 1 || day > 25 {
		return errors.New("advent of code only runs in December up until the 25th, so days must be 1-25 inclusive")
	}

	return nil
}
