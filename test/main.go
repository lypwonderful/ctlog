package main

import (
	"ctlog/ctlog"
	"time"
)

func main() {

	ctlog.SetLogLevel("info")
	//ctlog.SetLogDir("", "ct")

	for {
		ctlog.Debugln("debug level")
		ctlog.Infoln("info level")
		ctlog.Warningln("warn level")
		ctlog.Fatalln("fatal level")
		time.Sleep(1 * time.Millisecond)
	}
}
