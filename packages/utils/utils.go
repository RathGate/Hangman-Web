package utils

import "regexp"

func IsAlpha(str string) bool {
	return regexp.MustCompile(`^[A-Za-z]+$`).MatchString(str)
}
