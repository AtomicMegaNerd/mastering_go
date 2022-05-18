package main

import (
	"log"
	"log/syslog"
	"os"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "ch1_logging")

	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(sysLog)
		log.Println("Everything is fine!")
	}
}

func test_log1() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello world")
	}
	log.Panic("Panic: Hello world")
}
