package logger

import (
	"log"
	"os"
	"time"
)

//var Pname string

//LogErr will log the process logs into the project-1/logs directory corresponding to the correct process name
func LogErr(pName string) {
	yr, month, day := time.Now().Date()
	myTime := string(yr) + "_" + string(month) + "_" + string(day)
	path := "logs/" + pName + "/" + myTime
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetOutput(file)
}

// func init() {
// 	LogErr(Pname)

// }
