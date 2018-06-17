package ctlog

import (
	"fmt"
	"github.com/golang/glog"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var createLogDirOnce = sync.Once{}

// level of debug
func Debugln(v ...interface{}) {
	if logLevel >= levelDEBUG {
		log.SetPrefix("DEBUG\t")
		log.Output(2, fmt.Sprintln(v))
	}
}

// level of info
func Infoln(v ...interface{}) {
	if logLevel >= levelINFO {
		log.SetPrefix("INFO\t")
		log.Output(2, fmt.Sprintln(v))
	}
}

// level of warning
func Warningln(v ...interface{}) {
	log.SetPrefix("WARN\t")
	if logLevel >= levelWARN {
		log.Output(2, fmt.Sprintln(v))
	}
}

// level of fatal
func Fatalln(v ...interface{}) {
	log.SetPrefix("FATAL\t")
	if logLevel >= levelFATAL {
		log.Output(2, fmt.Sprintln(v))
	}
}

func createLogDir() {
	// ignore error
	os.MkdirAll(ctlog.logDir, os.ModePerm)
}

func logFileName() (fileName, link string) {
	t := time.Now()
	fileName = fmt.Sprintf("%s.log.%04d%02d%02d-%02d%02d%02d",
		ctlog.userName,
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
	return fileName, fileName + "-link"
}

func (l *logT) createLogFile() (err error) {
	createLogDirOnce.Do(createLogDir)
	fileName, link := logFileName()
	logFilePath := filepath.Join(ctlog.logDir, fileName)

	l.f, err = os.Create(logFilePath)
	if err == nil {
		symlink := filepath.Join(ctlog.logDir, link)
		os.Remove(symlink)            // ignore err
		os.Symlink(fileName, symlink) // ignore err
		return nil
	}
	return fmt.Errorf("Create Log File Fail: %v", err)
}

func (l *logT) outPut(calldepth int, s string) error {
	return log.Output(calldepth, s)
}

func (l *logT) Write(buf []byte) (n int, err error) {
	if l.f == nil {
		return len(buf), nil
	}
	return l.f.Write(buf)
}

func Test() {
	glog.Error()
}
