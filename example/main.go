package main

import (
	"log"
	"os"

	"github.com/doITmagic/decodestream"
)

func main() {
	stream := decodestream.NewJSONStream()
	go func() {
		for data := range stream.Watch() {
			if data.Error != nil {
				log.Println("error ", data.Error)
			}
			log.Println(data.Data)
		}
	}()

	// Open file to read.
	file, err := os.Open("ports.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//we can use any type that implement io.Reader
	stream.Start(file)

	_, closed := <-stream.Watch()
	if closed {
		os.Exit(0)
	}
}
