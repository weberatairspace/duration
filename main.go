package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func usage(msg string, code int) {
	if msg != "" {
		log.Print(msg)
	}

	log.Printf("usage: %s nanoseconds", os.Args[0])
	os.Exit(code)
}

func main() {
	if len(os.Args) != 2 {
		usage("invalid args", 1)
	}

	ns, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		usage(fmt.Sprintf("could not convert nanoseconds to strings: %v", err), 2)
	}

	fmt.Printf("%v nanoseconds is %v", os.Args[1], time.Duration(ns))
}
