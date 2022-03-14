package generator

import (
	"golock3r/server/logger"
	"testing"
)

func Test_0lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	Gp := GenPassword(0, true)
	ps := Passwordstren(Gp)

	if ps != 0 {
		t.Errorf("Expected Strength 0 of password strength %q", ps)
	}
}
func Test_8lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	Gp := GenPassword(8, false)
	ps := Passwordstren(Gp)

	if ps > 9 {
		t.Errorf("Expected Strength 0 of password strength %q", ps)
	}
}
func Test_16lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	Gp := GenPassword(16, false)
	ps := Passwordstren(Gp)

	if ps > 20 {
		t.Errorf("Expected Strength 0 of password strength %q", ps)
	}
}
