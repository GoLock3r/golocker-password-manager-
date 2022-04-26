package generator

import (
	"golock3r/server/logger"
	"testing"
)
//tests 0 length password
func Test_0lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(0, true)
	ps := Passwordstren(gp)

	if ps != 0 {
		t.Errorf("Expected Strength 0 of password strength %q", ps)
	}
}
//tests 8 length password with no special characters 
func Test_8lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(8, false)
	ps := Passwordstren(gp)

	if ps != 2 {
		t.Errorf("Expected Strength 2 of password strength %q", ps)
	}
}
//tests 16 length password with no special characters
func Test_16lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(16, false)
	ps := Passwordstren(gp)
	if ps != 2 {
		t.Errorf("Expected Strength 2 of password strength %q", ps)
	}
}
//tests 8 length password with special characters
func Test_8lengthpassWithSpecial(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(8, true)
	ps := Passwordstren(gp)

	if ps >=2 {
		t.Errorf("Expected Strength of at least 3 got %q", ps)
	}
}
//tests 16 length password with special characters
func Test_16lengthpassWithSpecial(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(16, true)
	ps := Passwordstren(gp)
	if ps < 3 {
		t.Errorf("Expected Strength of at least 3 got %q", ps)
	}
}
