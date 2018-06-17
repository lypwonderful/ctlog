package ctlog

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

const (
	levelDEBUG int = iota
	levelINFO
	levelWARN
	levelERROR
	levelFATAL
)

var (
	logLevel int
)

type logT struct {
	logLevel int
	logDir   string
	userName string
	f        *os.File
}

var ctlog = new(logT)

func init() {
	current, err := user.Current()
	if err == nil {
		ctlog.userName = current.Username
	}

	// Sanitize userName since it may contain filepath separators on Windows.
	ctlog.userName = strings.Replace(ctlog.userName, `\`, "_", -1)
}

func SetLogLevel(level string) {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		{
			logLevel = levelDEBUG
		}
	case "info":
		{
			logLevel = levelINFO
		}
	case "warning":
		{
			logLevel = levelWARN
		}
	case "error":
		{
			logLevel = levelERROR
		}
	case "fatal":
		{
			logLevel = levelFATAL
		}
	default:
		fmt.Println("Only Support LogLevel: debug info warning error fatal,Use default(error)")
		logLevel = levelERROR
	}
}

func SetLogDir(logDir string) {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(ctlog)
	if logDir == "" {
		ctlog.logDir = "./"
		ctlog.createLogFile()
		return
	}
	ctlog.logDir = logDir
	ctlog.createLogFile()
}
