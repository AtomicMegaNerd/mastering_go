package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Configure the log format
	LstdFlags := log.Ldate | log.Ltime | log.Lshortfile

	iLog := log.New(f, "iLog ", LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 3rd edition!")
}
