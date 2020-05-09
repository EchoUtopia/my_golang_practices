package main

import "time"

func main() {
	go runServer()
	go func() {
		time.Sleep(1 * time.Second)
		runClient()
	}()
	time.Sleep(10 * time.Hour)
}
