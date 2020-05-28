package main

import (
	"fmt"
	"github.com/catoaune/multichannel/channel/sms/linkmobility"
	"os"
)

func main() {
	linkUsername := os.Getenv("link_username")
	linkPassword := os.Getenv("link_password")
	linkPlatformID := os.Getenv("link_platformid")
	linkPartnerPlatformID := os.Getenv("link_partnerplatformid")
	linkSender := os.Getenv("link_sender")
	msg := os.Args[1]
	recipient := os.Args[2]

	linkConfig := linkmobility.NewConfig(linkSender, linkUsername, linkPassword, linkPlatformID, linkPartnerPlatformID)
	fmt.Println("Melding: " + msg)
	fmt.Println("Nr: " + recipient)
	fmt.Println(linkConfig.URL)
	fmt.Println("Username: " + linkUsername)
	fmt.Println("Sender: " + linkSender)
/*
	err := linkConfig.SendNotification(msg, recipient)
	if err != nil {
		log.Println("Error when sending SMS: ${err}")
	}
*/
}
