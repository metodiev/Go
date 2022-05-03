package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	syslog, err := syslog.New(syslog.Log_Info|syslog.LOG_LOCAL7, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(syslog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")

	syslog, err = syslog.New(syslog.LOG_MAIL, "Some Program")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(syslog)
	}

	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will see you this?")
}