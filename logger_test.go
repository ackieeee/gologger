package gologger

import (
	"reflect"
	"testing"
)

func TestSetLogger(t *testing.T) {
	testCases := []struct {
		level  string
		expect Config
	}{
		{
			"DEBUG",
			Config{
				MinLevel: "DEBUG",
				Levels:   []string{"ERROR", "WARNING", "INFO", "DEBUG"},
			},
		},
	}
	for _, testCase := range testCases {
		SetLogger(testCase.level)
		if !reflect.DeepEqual(testCase.expect, logger) {
			t.Fatalf("unexpected result. expect=%v actual=%v\n", testCase.expect, logger)
		}
	}
}
