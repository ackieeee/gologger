package main

import gologger "github.com/gba-3/gologger"

func main() {
	gologger.SetLogger("DEBUG")
	gologger.Debug("debug message")
	gologger.Info("info message")
	gologger.Warning("warning message")
	gologger.Error("error message")
}
