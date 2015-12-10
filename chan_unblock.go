package main
import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "World"
	}()

	for {
		select {
		case msg1 := <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("nothing received!")
			time.Sleep(time.Second)
		}
	}
}
