package utils

import (
	"strings"
	"time"
)

func GenerateRegistrationNumber() string {
	currentTime := time.Now()
	twoLetterMonth := currentTime.Format("Jan")[:2]
	formattedTime := currentTime.Format("06010215040506")
	registrationNumber := strings.ToUpper(twoLetterMonth) + "00" + formattedTime
	return registrationNumber
}
