package main

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
	"strconv"
	"github.com/mgutz/ansi"
)

func isLeapYear(year int) (bool, error) {
	if year < 0 {
		return false, errors.New("Can't calculate negative years")
	}
	if year%4 != 0 {
		return false, nil
	}
	if year%100 != 0 {
		return true, nil
	}
	if year%400 != 0 {
		return false, nil
	}
	return true, nil
}

func lengthOfMonth(month time.Month) uint {
	if month == time.February {
		leapYear, err := isLeapYear(time.Now().Year())
		if err != nil {
			panic(err)
		}
		if leapYear {
			return 29
		} else {
			return 28
		}
	}
	if month <= 7 {
		// month % 2 is 1 on odd months, and those happen to have 31 days
		return 30 + uint(month%2)
	}
	if month%2 == 0 {
		return 31
	}
	return 30
}

func getFirstDayThisMonth() time.Weekday {
	month := time.Now().Month()
	year := time.Now().Year()
	location := time.Now().Location()
	return time.Date(year, month, 1, 0, 0, 0, 0, location).Weekday()
}

func formatDayNumber(day uint) (text string) {
	// Convert to base 10
	text = strconv.FormatUint(uint64(day), 10)
	if day == uint(time.Now().Day()) {
		text = ansi.Color(text, "black:white")
	}
	if utf8.RuneCountInString(text) < 2 {
		text = " " + text
	}
	return
}

func printCalendar() {
	firstDay := getFirstDayThisMonth()

	fmt.Printf("  %v %v\n", time.Now().Month().String(), strconv.Itoa(time.Now().Year()))

	fmt.Print("Su Mo Tu We Th Fr Sa\n")
	for i := 0; i < 3*int(firstDay); i++ {
		fmt.Print(" ")
	}
	var i uint = 1
	for ; i <= lengthOfMonth(time.Now().Month()); i++ {
		fmt.Print(formatDayNumber(i))
		fmt.Print(" ")
		if (i + uint(firstDay)) % 7 == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
}

func main() {
	printCalendar()
}
