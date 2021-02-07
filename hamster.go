package main

import (
	"hamsterclient/client"
	"log"
)

func main() {

	client.SendEvent()
	log.Print("finished")

}
