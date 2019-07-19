// Author: James Mallon <jamesmallondev@gmail.com>
// logit package - lib created to print and write logs
package logit

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"
)

// init function - data and process initialization
func init() {
	Log.Filepath = "log/"
}

// Test function TestGetLogDate to evaluate getLogDate
func TestGetLogDate(t *testing.T) {
	logDate := Log.getLogDate()
	currDate := time.Now().Format("2006/01/02 15:04:05")
	if logDate != currDate {
		t.Errorf("Expected return from getLogDate to be the current date %s, but got %s ",
			time.Now().Format("2006/01/02 15:04:05"), // get the current date
			Log.getLogDate())                         // get the method date
	}
}

// Test function TestCreateDir to evaluate the createDir method
func TestCreateDir(t *testing.T) {
	Log.createDir() // creates the folder
	_, e := os.Stat(Log.Filepath)
	if e != nil { // check for non existent dir
		t.Errorf("Expected the directory to exists.")
	}
	os.Remove(Log.Filepath) // remove the dir
}

// Test function TestCheckPath to evaluate the checkPath method
func TestCheckPath(t *testing.T) {
	e := Log.checkPath()
	if e { // check for non existent dir
		t.Errorf("Expected the directory to not exists.")
	}
}

// Test function TestLoadCategories to evaluate loadCategories method
func TestLoadCategories(t *testing.T) {
	Log.loadCategories()
	if Log.categories["alert"][0] != "Alert:" {
		t.Errorf("Expected Log.categories[\"alert\"][0] == \"Alert\", but got %s", Log.categories["alert"][0])
	}
}

// Test function TestAppendCategories to evaluate AppendCategories
func TestAppendCategories(t *testing.T) {
	newCategory := map[string][]string{
		"checkpoint": {"Checkpoint:", "150.000.000,00"},
	}
	Log.AppendCategories(newCategory)
	if Log.categories["checkpoint"][0] != "Checkpoint:" {
		t.Errorf("Expected Checkpoint:, but got %s ", Log.categories["checkpoint"][0])
	}
}

// Test function TestWriteLog to evaluate WriteLog method
func TestWriteLog(t *testing.T) {
	Log.Filepath = fmt.Sprintf("%s%s.log", "logs/", time.Now().Format("2006_01_02"))
	Log.WriteLog("debug", "Testing...", Log.GetTraceMsg())

	// open and read the first line of the log file
	file, _ := os.Open(Log.Filepath)
	fs := bufio.NewScanner(file)
	fs.Scan()
	fline := fs.Text()

	// check for must have text
	match, _ := regexp.MatchString(".*Debug:.*Testing", fline)
	if !match {
		t.Errorf("Expected to find Debug: in the file")
	}

	os.Remove(Log.Filepath) // remove the file
	os.Remove("logs/")      // remove the dir
}

// Test function BenchmarkWriteLog to evaluate the WriteLog method
func BenchmarkWriteLog(b *testing.B) {
	Log.Filepath = fmt.Sprintf("%s%s.log", "logs/", time.Now().Format("2006_01_02"))
	for i := 0; i < b.N; i++ {
		Log.WriteLog("debug", "Testing...", Log.GetTraceMsg())

	}
	os.Remove(Log.Filepath) // remove the file
	os.Remove("logs/")      // remove the dir
}

// Test function TestGetTraceMsg to evaluate GetTraceMsg method
func TestGetTraceMsg(t *testing.T) {
	pattern := fmt.Sprintf(".*PID: %d", os.Getpid())
	match, _ := regexp.MatchString(pattern, Log.GetTraceMsg())
	if !match {
		t.Errorf("Expected to match the PID")
	}
}
