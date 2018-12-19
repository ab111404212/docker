package logger

import (
	"log"
	"os"
)

var logs *log.Logger

func init() {
	InitLogSys()
}
func InitLogSys() {
	logFile := "main.log"
	file, _ := os.Create(logFile)
	logs = log.New(file, "", log.Lshortfile|log.LUTC|log.Ldate|log.LstdFlags|log.Ltime)
}

func Println(v ...interface{}) {
	logs.Println(v...)
}
