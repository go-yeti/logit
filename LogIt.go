// Author: James Mallon <jamesmallondev@gmail.com>
// logit package - lib created to print and write logs
package logit

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Struct type syslog -
type syslog struct {
	file       *os.File
	Filepath   string
	log        *log.Logger
	categories map[string][]string
}

// to be used as an external pointer to the syslog struct type
var Syslog *syslog

// init function - initialize values and performs a pre instantiation to make this lib
// methods work as static methods and avoid external instantiation of the struct
func init() {
	lg := syslog{} // pre instantiation
	lg.Filepath = fmt.Sprintf("%s%s.log", "logs/", time.Now().Format("2006_01_02"))
	lg.loadCategories() // loads all categories
	var e error
	if !lg.checkPath() {
		e = lg.createDir()
	}
	if e == nil {
		lg.file, _ = os.OpenFile(lg.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 1444)
		lg.log = log.New(lg.file, "", log.Ldate|log.Ltime)
	}
	Syslog = &lg // exported variable receives the instance
}

// getLogDate method - returns a string with the log format date
func (this *syslog) getLogDate() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// createDir function - function attempts to create the log file dir in case it doesn't exists
func (this *syslog) createDir() (err error) {
	err = os.MkdirAll(filepath.Dir(this.Filepath), 0755)
	if err != nil {
		msg := fmt.Sprintf("Logit error: path %s doesn't exists or is not writable and cannot be created",
			this.Filepath)
		fmt.Printf("%s %s on %s\n", this.getLogDate(),
			msg, this.GetTraceMsg())
	}
	return
}

// checkPath method - verifies if the directory exists and is writable
func (this *syslog) checkPath() bool {
	if _, err := os.Stat(filepath.Dir(this.Filepath)); os.IsNotExist(err) {
		return false
	}
	return true
}

// loadCategories method - loads all categories
func (this *syslog) loadCategories() {
	this.categories = map[string][]string{
		"emergency": {"Emergency:", "an emergency"},
		"alert":     {"Alert:", "an alert"},
		"critical":  {"Critical:", "a critical"},
		"error":     {"Error:", "an error"},
		"warning":   {"Warning:", "a warning"},
		"notice":    {"Notice:", "a notice"},
		"info":      {"Info:", "an info"},
		"debug":     {"Debug:", "a debug"},
	}
}

// AppendCategories method - it allow the user to append new categories
func (this *syslog) AppendCategories(newCategories map[string][]string) {
	for k, v := range newCategories {
		this.categories[k] = v
	}
}

// WriteLog method - writes the message to the log file
func (this *syslog) WriteLog(category string, msg string, trace string) {
	val, res := this.categories[category]
	if !res {
		fmt.Printf("%s %s The category %s does not exists on %s\n",
			time.Now().Format("2006/01/02 15:04:05"),
			this.categories["warning"][0], category, this.GetTraceMsg())
		this.log.Printf("%s (non existent category) %s on %s", category, msg, trace)
	} else {
		this.log.Printf("%s %s on %s", val[0], msg, trace)
	}
	defer this.file.Close()
}

// GetTraceMsg method - get the full error stack trace
func (this *syslog) GetTraceMsg() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d PID: %d", frame.File, frame.Line, os.Getpid())
}
