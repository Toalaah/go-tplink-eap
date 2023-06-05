package main

import (
	"log"
	"os"

	"github.com/toalaah/go-tplink-eap/pkg/tplink"
)

func main() {
	baseaddr := os.Getenv("TPLINK_ADDR")
	username := os.Getenv("TPLINK_USERNAME")
	password := os.Getenv("TPLINK_PASSWORD")
	radioID := 1

	c := tplink.NewClient(baseaddr, username, password)

	ssids, err := c.GetSSIDDs(radioID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Currently have %d SSIDs configured on radioID %d: %v\n", len(ssids), radioID, ssids)

	ssid := tplink.SSID{
		Ssidname:          "guest_network",
		Ssidbcast:         0,
		SecurityMode:      3,
		PskVersion:        3,
		PskCipher:         3,
		PskKey:            "strongpassword",
		PskKeyUpdate:      0,
		WpaVersion:        3,
		WepSelect:         0,
		Guest:             1,
		Limit:             false,
		LimitDownloadUnit: 1,
		LimitUploadUnit:   1,
	}

	ssids, err = c.CreateSSIDD(ssid, radioID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully created new SSID on radioID %d: %v\n", radioID, ssids)

	found := false
	idx := 0
	for i, s := range ssids {
		if s.Ssidname == ssid.Ssidname {
			found = true
			idx = i
		}
	}

	if !found {
		log.Fatal("Failed to find SSID in list of SSIDs")
	}

	if _, err = c.DeleteSSIDD(ssids[idx], radioID); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully deleted SSID on radioID %d: %v\n", radioID, ssids[idx])
}
