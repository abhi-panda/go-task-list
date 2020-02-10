package utilities

import "regexp"

//CheckDateStringFormat checks and tells if the date format provided is correct or not
func CheckDateStringFormat(dt string) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	return re.MatchString(dt)
}
