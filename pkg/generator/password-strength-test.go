package generator

import (
	"strings"
)

func Passwordstren(password string) int {
	var strength string
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

	special_chars := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "=", "+", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	capLet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, c := range special_chars {
		if strings.Contains(password, c) {
			hasspec = 1
		}
	}
	for _, c := range numbers {
		if strings.Contains(password, c) {
			hasnum = 1
		}
	}
	for _, c := range capLet {
		c = strings.ToUpper(c)
		if strings.Contains(password, c) {
			hascap = 1
		}
	}

	if length+hascap+hasspec+hasnum+hasnodicword == 4 {
		strength = "your password is decent"
	}
	if length+hascap+hasspec+hasnum+hasnodicword <= 3 {
		strength = "your password sucks figure it out"
	}
	if length+hascap+hasspec+hasnum+hasnodicword == 5 {
		strength = "eyyyyy lets get it your password is strong"
	}
	return strength
}
