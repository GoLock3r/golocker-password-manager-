package generator

import (
	"golock3r/server/logger"
	"strings"
	"testing"
)

func Test_zero_password(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	GP := GenPassword(0, false)
	if GP != "" {
		t.Errorf("Expected empty string got something that wasnt an empty string %q ", GP)
	}
}

func Test_outofsize_password(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	GP := GenPassword(18, false)
	if GP != "" {
		t.Errorf("Expected an empty string got somthing that wasnt an empty string %q", GP)
	}
}

func Test_password_gen_length(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	GP := GenPassword(16, false)
	if len(GP) != 19 {
		t.Errorf("Length of password was expected to be 16 instead it was %q", GP)
	}
}

func Test_password_gen_spec(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	GP := GenPassword(16, true)

	special_chars := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "=", "+", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	contains := false
	for _, c := range special_chars {
		if strings.Contains(GP, c) {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("Password does not contain a special character / number %q", GP)
	}
}
