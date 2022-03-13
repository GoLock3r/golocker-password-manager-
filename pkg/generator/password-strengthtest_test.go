package generator

import (
	"golock3r/server/logger"
	"testing"
)

func test_0lengthpass(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	Gp := GenPassword(0, true)
	ps := Passwordstren(Gp)

	if ps != 0 {
		t.Errorf("Expected Strength 0 of password strength %q", ps)
	}
}
