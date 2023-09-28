package main

import (
	"fmt"
	"main/notify"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	err := notify.FlashConsoleWindow()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
}
