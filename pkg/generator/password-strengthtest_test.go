package generator

import (
	"golock3r/server/logger"
	"testing"
)

func Test_0lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(0, true)
	ps := Passwordstren(gp)

	if ps != 0 {
		t.Errorf("Expected Strength 0 of password strength %q", ps)
	}
}
func Test_8lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(8, false)
	ps := Passwordstren(gp)

	if ps != 2 {
		t.Errorf("Expected Strength 2 of password strength %q", ps)
	}
}
func Test_16lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(16, false)
	ps := Passwordstren(gp)
	println(ps)
	if ps != 2 {
		t.Errorf("Expected Strength 20 of password strength %q", ps)
	}
}
func Test_8lengthpassWithSpecial(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(8, true)
	ps := Passwordstren(gp)

	if ps != 4 {
		t.Errorf("Expected Strength 2 of password strength %q", ps)
	}
}
func Test_16lengthpassWithSpecial(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	gp := GenPassword(16, true)
	ps := Passwordstren(gp)
	println(ps)
	if ps != 4 {
		t.Errorf("Expected Strength 20 of password strength %q", ps)
	}
}
