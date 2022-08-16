package gologger

import (
	"encoding/json"
	"os"
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
		{
			"INFO",
			Config{
				MinLevel: "INFO",
				Levels:   []string{"ERROR", "WARNING", "INFO"},
			},
		},
		{
			"WARNING",
			Config{
				MinLevel: "WARNING",
				Levels:   []string{"ERROR", "WARNING"},
			},
		},
		{
			"ERROR",
			Config{
				MinLevel: "ERROR",
				Levels:   []string{"ERROR"},
			},
		},
		{
			"TEST",
			Config{
				MinLevel: "TEST",
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

func TestPrint(t *testing.T) {
	t.Helper()
	backup := os.Stdout
	defer func() {
		os.Stderr = backup
	}()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("fail pipe: %v", err)
	}
	os.Stdout = w

	SetLogger("DEBUG")
	print("INFO", "info message")
	w.Close()

	log := Logger{}
	if err := json.NewDecoder(r).Decode(&log); err != nil {
		t.Fatalf("fail decode json: %v", err)
	}

	if log.Msg != "info message" {
		t.Fatalf("unexpected result. expect=%s actual=%s", "info message", log.Msg)
	}
}
