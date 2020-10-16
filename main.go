package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Printf(validateTimeEntry(input))
}

func validateTimeEntry(input string) (durationValue string) {

	value, _ := strconv.Atoi(input)

	seconds := value % 60
	minutes := (value % 3600) / 60
	hours := (value % 86400) / 3600
	days := (value % (86400 * 30)) / 86400

	var d, h, m, s string

	if len(strconv.Itoa(days)) <= 1 {
		d = "0" + strconv.Itoa(days) + "d"
	} else {
		d = strconv.Itoa(days) + "d"
	}

	if len(strconv.Itoa(hours)) <= 1 {
		h = "0" + strconv.Itoa(hours) + "h"
	} else {
		h = strconv.Itoa(hours) + "h"
	}

	if len(strconv.Itoa(minutes)) <= 1 {
		m = "0" + strconv.Itoa(minutes) + "m"
	} else {
		m = strconv.Itoa(minutes) + "m"
	}

	if len(strconv.Itoa(seconds)) <= 1 {
		s = "0" + strconv.Itoa(seconds) + "s"
	} else {
		s = strconv.Itoa(seconds) + "s"
	}

	durationValue = d + h + m + s

	return
}
