package main

import "github.com/rickb777/date/gregorian"

func isLeapYear(year int) bool {
	return gregorian.IsLeap(year)
}

func main() {
	println("Hello, world!")
	println("Is 2020 a leap year?", isLeapYear(2020))
}
