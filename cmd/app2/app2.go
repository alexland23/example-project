package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	log.Println("Starting Application 2 - CLI Tool")

	// Create flags
	namePtr := flag.String("name", "example", "your name")
	numPtr := flag.Int("number", 0, "an int")
	boolPtr := flag.Bool("bool", false, "a bool")

	flag.Parse()

	fmt.Println("name  :", *namePtr)
	fmt.Println("number:", *numPtr)
	fmt.Println("bool  :", *boolPtr)
}
