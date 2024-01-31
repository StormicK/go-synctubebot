package main

import (
	"log"
	"synctubebot/api"
)

func main() {
	//will be called when the program panics and it is not handled
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	api.Init()
}
