
package CoreMailer

import (
	"regexp"
)

const emailRegexPattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`


func IsValidEmail(email string) bool {

	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)

}


