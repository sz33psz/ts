package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sz33psz/ts"
)

func main() {
	args := os.Args[1:]

	currentTime := time.Now().Unix()
	for _, arg := range args {
		if chg, err := ts.NewChange(arg); err == nil {
			currentTime = chg.Apply(currentTime)
		}
	}
	fmt.Printf("%v\n", currentTime)
}
