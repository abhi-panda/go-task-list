package utilities

import (
	"regexp"
	"strconv"
	"strings"
)

//CheckDateStringFormat checks and tells if the date format provided is correct or not
func CheckDateStringFormat(dt string) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	matched := re.MatchString(dt)
	if matched == true {
		date := strings.Split(dt, "-")
		year, _ := strconv.Atoi(date[0])
		month, _ := strconv.Atoi(date[1])
		day, _ := strconv.Atoi(date[2])
		if year > 1970 && month < 13 && month > 0 && day > 0 && day < 32 && matched {
			return true
		}
		return false
	}
	return false

}
