package main

import "fmt"
import "time"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "hello"
		fmt.Println("write \"hello\" done!")

		channel <- "World"
		fmt.Println("write go sleep...")
		time.Sleep(3*time.Second)
		channel <- "channel"

		fmt.Println("write \"channel\" done!")
	}()

	time.Sleep(2*time.Second)
	fmt.Println("Reader Wake up...")

	msg := <-channel
	fmt.Println("Reader: ", msg)

	msg = <-channel
	fmt.Println("Reader: ", msg)

	msg = <- channel
	fmt.Println("Reader: ", msg)
}
