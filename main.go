package holidays

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var holidayDate = time.Date(2021, 10, 01, 02, 34, 58, 651387237, time.UTC)

	fmt.Println(CountHoliday(holidayDate))
	loc, _ := time.LoadLocation("UTC")
	today := time.Now().In(loc)
	var resumeDatetime = time.Date(today.Year(), today.Month(), today.Day(), 01, 30, 0, 0, time.UTC)
	var cutOffDatetime = time.Date(today.Year(), today.Month(), today.Day(), 8, 30, 0, 0, time.UTC)
	fmt.Println("currentTime: ", today.Local())
	fmt.Println("resumeDatetime: ", resumeDatetime.Local())
	fmt.Println("cutOffDatetime:", cutOffDatetime.Local())

	fmt.Println("Today is :", today.Weekday())
	in := time.Now().In(loc).Add(750 * time.Minute)
	if in.After(cutOffDatetime) {
		fmt.Println("CurrentDate is after cutoff time", in.Local())
		resumeDatetime = resumeDatetime.Add(24 * time.Hour)
		cutOffDatetime = cutOffDatetime.Add(24 * time.Hour)
	}
	fmt.Println("resumeDatetime: ", resumeDatetime.Local())
	fmt.Println("cutOffDatetime:", cutOffDatetime.Local())
}

//func inTimeSpan(resumeDatetime, cutOffDatetime, check time.Time) bool {
//	return check.After(resumeDatetime) && check.Before(cutOffDatetime)
//}

func inOperationTime(currentTime time.Time) (bool, *int, string) {
	loc, _ := time.LoadLocation("UTC")
	var message string = ""
	currentTime = currentTime.In(loc)
	var resumeDatetime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 01, 30, 0, 0, time.UTC)
	var cutOffDatetime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 8, 30, 0, 0, time.UTC)
	// same both after cutoff and resume
	if currentTime.After(cutOffDatetime) && currentTime.After(resumeDatetime) {
		resumeDatetime = resumeDatetime.Add(24 * time.Hour)
		diff := int(math.Abs(resumeDatetime.Sub(currentTime).Seconds()))
		message = "after same cutoff time. " + currentTime.Local().String()
		return false, &diff, message
	} else if currentTime.After(resumeDatetime) && currentTime.Before(cutOffDatetime) {
		message = "currentTime is " + currentTime.Local().String() + " between" + resumeDatetime.Local().String() +
			" to " + cutOffDatetime.Local().String()
		diff := 0
		return true, &diff, message
	} else {
		message = "where am I ?" + currentTime.Local().String()
		diff := int(math.Abs(currentTime.Sub(resumeDatetime).Seconds()))
		return false, &diff, message
	}
}

func CountHoliday(holidayDate time.Time) int {
	var durationInDays int
	for {
		isHoliday, _ := IsHoliday(holidayDate) // true
		if isHoliday {
			holidayDate = holidayDate.Add(24 * time.Hour)
			durationInDays++
		} else {
			break
		}
	}
	return durationInDays
}
