package utilities

import (
	"time"
)

//SecondsLeftInDay returns as the name suggests seconds left in a day , needed for correction while checking duedates
func SecondsLeftInDay() int64 {
	t := time.Now()
	var result int64
	result += int64((t.Hour() * 3600))
	result += int64((t.Minute() * 60))
	result += int64((t.Second()))
	return (86399 - result)
}

//SecondsOccuredInDay returns as the name suggests seconds left in a day , needed for correction while checking duedates
func SecondsOccuredInDay() int64 {
	t := time.Now()
	var result int64
	result += int64((t.Hour() * 3600))
	result += int64((t.Minute() * 60))
	result += int64((t.Second()))
	return (result + 1)
}
