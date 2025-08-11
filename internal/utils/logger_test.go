package utils

import (
	"testing"
)

func TestInitLoggerDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("InitLogger panicked: %v", r)
		}
	}()
	InitLogger()
}

func TestInitLoggerMultipleCalls(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("InitLogger panicked on repeated call: %v", r)
		}
	}()
	InitLogger()
	InitLogger()
}