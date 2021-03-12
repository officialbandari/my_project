package main

import (
	"fmt"
	"time"
)

const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct{

	time time.Time 
	severity string
	message string 

}

var logCh = make(chan logEntry, 50)

func main()  {

	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond )
	
}

func logger()  {

	for entry := range logCh{
		fmt.Printf("%v - [%v]%v \n ", entry.time.Format("2021-03-02T15-:04:05"), entry.severity, entry.message)
	}
	
}