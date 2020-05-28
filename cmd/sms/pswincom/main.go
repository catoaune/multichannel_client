package main

import (
	"fmt"
	"os"
	"github.com/catoaune/multichannel/channel/sms/pswincom"
)

func main() {
	username := os.Getenv("pswincom_username")
	password := os.Getenv("pswincom_password")
	sender := os.Getenv("pswincom_sender")
	msg := os.Args[1]
	recipient := os.Args[2]
	pswincomConfig := pswincom.NewConfig(username, password, sender)
	err := pswincomConfig.SendNotification(msg, recipient)
	if err != nil {
		fmt.Println("Error: {err}")
	}
}
