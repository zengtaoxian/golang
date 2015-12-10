package main
import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second *1)
		c1 <- "Hello"
	}()

	go  func() {
		time.Sleep(time.Second *1)
		c2 <- "World"
	}()

	timeout_cnt := 0
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1 received", msg1)
		case msg2 := <-c2:
			fmt.Println("msg2 received", msg2)
		case <-time.After(time.Second * 1):
			fmt.Println("Time Out")
			timeout_cnt++
		}

		fmt.Println("cnt", timeout_cnt)

		if timeout_cnt > 3 {
			break
		}
	}
}
