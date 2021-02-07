package main

import (
	"github.com/Budlee/hamsterclient/client"
	"log"
)

func main() {

	client.SendEvent()
	log.Print("finished")

}
