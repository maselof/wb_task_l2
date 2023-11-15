package main

import (
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	errLog := log.New(os.Stderr, "Error: ", 0)

	currTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		errLog.Fatal(err)
	}

	log.Printf("currTime = %s", currTime)

}
