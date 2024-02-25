package main

import (
	"fmt"
	"time"
)

const TimeInterval = 10

func main() {
	for {
		fmt.Println("Test Shifu")
		time.Sleep(time.Second * time.Duration(TimeInterval))
	}
}
