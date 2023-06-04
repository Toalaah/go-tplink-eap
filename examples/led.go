package main

import (
	"fmt"
	"log"
	"os"

	"github.com/toalaah/go-tplink-eap/pkg/tplink"
)

func main() {
	baseAddr := os.Getenv("TPLINK_ADDR")
	username := os.Getenv("TPLINK_USERNAME")
	password := os.Getenv("TPLINK_PASSWORD")

	c := tplink.NewClient(baseAddr, username, password)

	info, err := c.GetLedStatus()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Got LED status %s\n", info.Enable)
	}

	var newStatus tplink.LedStatus
	if info.Enable == string(tplink.LedStatusOn) {
		newStatus = tplink.LedStatusOff
	} else {
		newStatus = tplink.LedStatusOn
	}

	info, err = c.SetLedStatus(newStatus)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Set LED status to %s\n", info.Enable)
	}

}
