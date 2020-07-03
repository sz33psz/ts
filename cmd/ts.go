package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sz33psz/ts"
)

func main() {
	args := os.Args[1:]

	currentTime := time.Now().Unix()
	printIso := false
	for _, arg := range args {
		if arg == "-t" {
			printIso = true
			continue
		}
		if override, err := strconv.ParseInt(arg, 10, 64); err == nil {
			currentTime = override
		}
		if chg, err := ts.NewChange(arg); err == nil {
			currentTime = chg.Apply(currentTime)
		}
	}

	if printIso {
		t := time.Unix(currentTime, 0).UTC()
		fmt.Println(t.Format(time.RFC3339))
	} else {
		fmt.Printf("%v\n", currentTime)
	}
}
