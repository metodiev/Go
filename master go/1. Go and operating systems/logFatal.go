package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	syslog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(syslog)
	}

	log.Fatal(syslog)
	fmt.Println("Will you see this?")

}