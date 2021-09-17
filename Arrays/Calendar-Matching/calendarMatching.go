// Given two calendar and two corresponding dailyBounds denoting the daily working hours of two employees,
// find out all the possible gaps that's longer than the given meeting duration, so that these two employees
// can both be free for meeting
// solution walkthrough: let's say for a dailyBound ["9:00", "18:00"], insert ["0:00", "9:00"] at the beginning
// of the calendar, and append ["18:00", "23:59"] at the end, so as to simulate another two meetings in one day
// so that the corresponding timespan is not available; then try to transform these updated calendars in numeric
// form and merge afterwards in a merge-sort manner (two pointers, take only the first element, i.e. the start time
// of a meeting into account); flatten the merged calendar to eleminate any overlap, finally iterate through
// the flattened calendar (numeric form, with no overlap, each element in this calendar represents a stand-alone
// meeting), and return all the gaps that fulfills the duration requirement.

package calendarmatching

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Big O: O(c1 + c2) time | O(c1 + c2) space
// where c1, c2 are the length of calendar1, calendar2 respectively
func CalendarMatching(calendar1 [][2]string, dailyBounds1 [2]string, calendar2 [][2]string, dailyBounds2 [2]string, meetingDuration int) [][2]string {
	numericCalendar1 := updateCalendar(calendar1, dailyBounds1)
	numericCalendar2 := updateCalendar(calendar2, dailyBounds2)
	mergedCalendar := mergeCalendars(numericCalendar1, numericCalendar2)
	flattenedCalendar := flattenCalendar(mergedCalendar)

	return getMatchingAvailabilities(flattenedCalendar, meetingDuration)
}

func updateCalendar(calendar [][2]string, dailyBounds [2]string) [][2]int {
	updatedCalendar := append(make([][2]string, 1), calendar...)
	updatedCalendar[0] = [2]string{"0:00", dailyBounds[0]}
	updatedCalendar = append(updatedCalendar, [2]string{dailyBounds[1], "23:59"})

	numericCalendar := make([][2]int, len(updatedCalendar))
	for i := range numericCalendar {
		numericCalendar[i] = [2]int{
			timeToMinutes(updatedCalendar[i][0]),
			timeToMinutes(updatedCalendar[i][1]),
		}
	}
	return numericCalendar
}

func mergeCalendars(calendar1, calendar2 [][2]int) [][2]int {
	merged := [][2]int{}
	var i, j int
	for i < len(calendar1) && j < len(calendar2) {
		meeting1, meeting2 := calendar1[i], calendar2[j]
		if meeting1[0] < meeting2[0] {
			merged = append(merged, meeting1)
			i += 1
		} else {
			merged = append(merged, meeting2)
			j += 1
		}
	}
	if i == len(calendar1) {
		merged = append(merged, calendar2[j:]...)
	} else {
		merged = append(merged, calendar1[i:]...)
	}
	return merged
}

// expert's flatten solution
func flattenCalendar(calendar [][2]int) [][2]int {
	flattened := append([][2]int(nil), calendar[0])
	for i := 1; i < len(calendar); i++ {
		currentMetting := calendar[i]
		// previous meeting is the running time bound in the flattened array,
		// if no overlap between it and current meeting, then it would be set
		// else it will be merged with current meeting
		previousMeeting := flattened[len(flattened)-1]
		currentStart, currentEnd := currentMetting[0], currentMetting[1]
		previousStart, previousEnd := previousMeeting[0], previousMeeting[1]
		if previousEnd >= currentStart {
			// overlap occurs
			newPreviousMeeting := [2]int{previousStart, int(math.Max(float64(previousEnd), float64(currentEnd)))}
			flattened[len(flattened)-1] = newPreviousMeeting
		} else {
			// no overlap, last element in flattened array is safe to be set as a stand-alone
			// meeting, and append current meeting for subsequent check
			flattened = append(flattened, [2]int{currentStart, currentEnd})
		}
	}
	return flattened
}

// // my flatten solution
// func flattenCalendar(calendar [][2]int) [][2]int {
// 	flattenedCalendar := [][2]int{}
// 	i, j := 0, 1
// 	for j < len(calendar) {
// 		currentStart, currentEnd := calendar[i][0], calendar[i][1]
// 		nextStart := calendar[j][0]
// 		if currentEnd < nextStart {
// 			currentBound := [2]int{currentStart, calendar[j-1][1]}
// 			flattenedCalendar = append(flattenedCalendar, currentBound)
// 			i = j
// 			j = i + 1
// 		} else {
// 			j += 1
// 		}
// 	}
// 	// j == len(calendar), whilst i <= len(calendar)-1
// 	// in other word, may be only one last element that could be appended directly
// 	// or more than one elements that should be merged with element i and j-1
// 	currentStart, currentEnd := calendar[i][0], calendar[i][1]
// 	nextEnd := calendar[j-1][1]
// 	currentBound := [2]int{currentStart, int(math.Max(float64(currentEnd), float64(nextEnd)))}
// 	flattenedCalendar = append(flattenedCalendar, currentBound)
// 	return flattenedCalendar
// }

func getMatchingAvailabilities(calendar [][2]int, meetingDuration int) [][2]string {
	matchingAvailabilities := make([][2]string, 0)
	for i := 1; i < len(calendar); i++ {
		previousMeetingEnd := calendar[i-1][1]
		currentMeetingStart := calendar[i][0]
		if currentMeetingStart-previousMeetingEnd >= meetingDuration {
			matchingAvailabilities = append(matchingAvailabilities, [2]string{
				minutesToTime(calendar[i-1][1]),
				minutesToTime(calendar[i][0]),
			})
		}
	}
	return matchingAvailabilities
}

// transform literal time into numeric value for ease of comparison, e.g. "1:30" => 90
func timeToMinutes(literalTime string) int {
	timeComponents := strings.Split(literalTime, ":")
	hours, _ := strconv.ParseInt(timeComponents[0], 10, 0)
	minutes, _ := strconv.ParseInt(timeComponents[1], 10, 0)

	return int(60*hours + minutes)
}

func minutesToTime(minute int) string {
	hours := fmt.Sprint(math.Floor(float64(minute) / 60))
	mins := fmt.Sprint(minute % 60)
	if minute%60 < 10 {
		mins = fmt.Sprintf("0%v", mins) // padding with 0 if mins is single digit
	}
	return fmt.Sprintf("%v:%v", hours, mins)
}
