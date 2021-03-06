package main

import (
	"os"

	"github.com/catoaune/multichannel/channel/sms/twilio"
)

func main() {
	sid := os.Getenv("sid")
	token := os.Getenv("token")
	number := os.Getenv("number")
	msg := os.Args[1]
	recipient := os.Args[2]
	twilioConfig := twilio.NewConfig(sid, token, number)
	twilioConfig.SendNotification(msg, recipient)
}
