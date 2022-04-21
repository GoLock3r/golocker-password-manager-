package generator

import (
	"strings"
)
//Tests the strength of the password based off of length, contains capital letters, contains numbers, and special characters. Will have a dictonary implimented at somepoint to detect dictionary words
func Passwordstren(password string) int {
	var strength int
	var length int
	var hascap int
	var hasspec int
	var hasnum int
	var hasnodicword int

	if len(password) > 8 {
		length = 1
	} else {
		length = 0
	}

	specialChars := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "=", "+"}
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	capLet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, c := range specialChars {
		if strings.Contains(password, c) {
			hasspec = 1
			break
		}
	}
	for _, c := range numbers {
		if strings.Contains(password, c) {
			hasnum = 1
			break
		}
	}
	for _, c := range capLet {
		c = strings.ToUpper(c)
		if strings.Contains(password, c) {
			hascap = 1
			break
		}
	}

	if length+hascap+hasspec+hasnum+hasnodicword == 4 {
		strength = length + hascap + hasspec + hasnum + hasnodicword
		//maybe log the strength of password decent
	}
	if length+hascap+hasspec+hasnum+hasnodicword <= 3 {
		strength = length + hascap + hasspec + hasnum + hasnodicword
		// weaaak
	}
	if length+hascap+hasspec+hasnum+hasnodicword == 5 {
		strength = length + hascap + hasspec + hasnum + hasnodicword
		//strongest we can tell most likely
	}
	return strength
}
