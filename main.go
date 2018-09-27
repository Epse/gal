package main

import (
	"errors"
	"fmt"
	"github.com/mgutz/ansi"
	"strconv"
	"time"
	"os"
	"unicode/utf8"
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

func lengthOfMonth(date time.Time) uint {
	month := date.Month()
	if month == time.February {
		leapYear, err := isLeapYear(date.Year())
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

func getFirstDayMonth(date time.Time) time.Weekday {
	month := date.Month()
	year := date.Year()
	location := date.Location()
	return time.Date(year, month, 1, 0, 0, 0, 0, location).Weekday()
}

func formatDayNumber(day uint, date time.Time) (text string) {
	// Convert to base 10
	text = strconv.FormatUint(uint64(day), 10)
	if utf8.RuneCountInString(text) < 2 {
		text = " " + text
	}
	if day == uint(date.Day()) {
		text = ansi.Color(text, "black:white")
	}
	return
}

func printCalendar(date time.Time) {
	firstDay := getFirstDayMonth(date)

	fmt.Printf("  %v %v\n", date.Month().String(), strconv.Itoa(date.Year()))

	fmt.Print("Su Mo Tu We Th Fr Sa\n")
	for i := 0; i < 3*int(firstDay); i++ {
		fmt.Print(" ")
	}
	var i uint = 1
	for ; i <= lengthOfMonth(date); i++ {
		fmt.Print(formatDayNumber(i, date))
		fmt.Print(" ")
		if (i+uint(firstDay))%7 == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
}

func printUsage() {
	fmt.Print("Usage: gal [DATE]\n")
	fmt.Print("DATE is a date string, like yyyy-mm-dd\n")
}

func main() {
	if len(os.Args[1:]) == 0 {
		printCalendar(time.Now())
	} else if len(os.Args[1:]) == 1 {
		arg := os.Args[1]
		iso8081 := "2006-01-02"
		time, err := time.Parse(iso8081, arg)
		if err != nil {
			panic(err)
		}
		printCalendar(time)
	} else {
		printUsage()
	}
}
