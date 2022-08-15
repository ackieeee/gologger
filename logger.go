package gologger

import (
	"encoding/json"
	"log"
	"os"
)

var logger Config

type Config struct {
	MinLevel string
	Levels   []string
}

type Logger struct {
	Msg   string `json:"msg"`
	Level string `json:"level"`
	Pid   int    `json:"pid"`
}

func SetLogger(level string) {
	levels := []string{"ERROR", "WARNING", "INFO", "DEBUG"}

	ls := []string{}
	for _, l := range levels {
		ls = append(ls, l)
		if l == level {
			break
		}
	}
	logger = Config{MinLevel: level, Levels: ls}
}

func print(level string, msg string) {
	if !containsLevel(level) {
		return
	}
	pid := os.Getpid()
	l := Logger{
		Msg:   msg,
		Level: level,
		Pid:   pid,
	}
	buf, err := json.Marshal(l)
	if err != nil {
		return
	}
	log.Println(string(buf))
}

func Debug(msg string) {
	print("DEBUG", msg)
}

func Info(msg string) {
	print("INFO", msg)
}

func Warning(msg string) {
	print("WARNING", msg)
}

func Error(msg string) {
	print("ERROR", msg)
}

func containsLevel(level string) bool {
	for _, l := range logger.Levels {
		if level == l {
			return true
		}
	}
	return false
}
