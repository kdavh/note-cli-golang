package main

import (
	"fmt"
	"os"
)

func errExit(e error) {
	if e != nil {
		fmt.Printf("ERROR, EXITING: %v", e)
		os.Exit(1)
	}
}
