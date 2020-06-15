package main

import (
	"flag"
	"os"
)

func main() {
	args := os.Args[1:]
	millis := flag.Bool("m", false, "milliseconds mode")
	flag.Parse()
}
