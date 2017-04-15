// You are given the following information, but you may prefer to do some research for yourself.

// 1 Jan 1900 was a Monday.
// Thirty days has September, April, June and November.
// All the rest have thirty-one, Saving February alone, Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
package main

import "fmt"

// This slice represents the first week day of the month (1 to 7) in each month (slice position)
// Starting with year 1901
const year1901 = []int{3, 6, 6, 2, 4, 7, 2, 5, 1, 3, 6, 1}

// Get total sundays that fall in the first day of the month
func countFirstSundays(yearFirstDays []int) int {
	total := 0
	for _, i := range yearFirstDays {
		if i == 1 {
			total++
		}
	}
	return total
}

// Check if year is leap
func isLeapYear(year int) bool {
	return (year % 4) == 0
}

// Set array with values corresponding to days from respective year
func setNextYearFirstDays(yearDays *[]int, year int) {
	for month := range *yearDays {
		if month == 1 && isLeapYear(year) {
			(*yearDays)[month] += 2
		} else {
			(*yearDays)[month]++
		}

		if (*yearDays)[month] > 7 {
			(*yearDays)[month] -= 7
		}
	}
}

func main() {
	totalSundays := 0

	yearFirsMonthtDays := year1901

	for year := 1901; year <= 2000; year++ {
		totalSundays += countFirstSundays(yearFirsMonthtDays)
		setNextYearFirstDays(&yearFirsMonthtDays, year)
	}

	fmt.Printf("total sundays: %d\n\n", totalSundays)
}
