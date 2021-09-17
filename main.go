// golang entry for each AlgoExpert challenge
// #[Arrays]/Calendar Matching

package main

import (
	"fmt"

	array "github.com/letempest/algoexpert-in-action/Arrays/Calendar-Matching"
)

func main() {
	calendar1 := [][2]string{
		{"9:00", "10:30"},
		{"12:30", "13:00"},
		{"16:00", "18:00"},
	}
	dailyBounds1 := [2]string{"9:00", "20:00"}
	calendar2 := [][2]string{
		{"10:00", "11:30"},
		{"12:30", "14:30"},
		{"14:30", "15:00"},
		{"16:00", "17:00"},
	}
	dailyBounds2 := [2]string{"10:00", "18:30"}

	meetingDuration := 30

	availableGaps := array.CalendarMatching(calendar1, dailyBounds1, calendar2, dailyBounds2, meetingDuration)
	fmt.Printf("%#v\n", availableGaps)
}
