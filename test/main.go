package main

import "ctlog/ctlog"

func main() {

	ctlog.SetLogLevel("info")
	ctlog.SetLogDir("")

	ctlog.Debugln("debug level")
	ctlog.Infoln("info level")
	ctlog.Warningln("warn level")
	ctlog.Fatalln("fatal level")
}
