package main

import (
	"fmt"
	"sync"
)

var waitG sync.WaitGroup

func main() {
	//var c = make(chan string)
	//start := time.Now()
	//
	//go message1(c)
	//go message2(c)
	//go message3(c)
	//
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//
	//close(c)
	//
	//elapsed := time.Since(start)
	//
	//fmt.Println(elapsed)

	//var c chan int
	//
	//if c == nil {
	//	c = make(chan int)
	//}
	////c := make(chan int)
	//
	//fmt.Printf("type of c: %T", c)
	//close(c)

	//c := make(chan int)

	//c <- "No one likes my channel"

	c := make(chan int)
	waitG.Add(1)

	go send(c)
	go receive(c)

	waitG.Wait()
	//time.Sleep(2 * time.Second)
}

func send(c chan int) {
	for i := 1; i < 6; i++ {
		c <- i
	}
	waitG.Done()
}

func receive(c chan int) {
	for {
		fmt.Println(<-c)
	}
	waitG.Done()
}

//func message1(c chan string) {
//	time.Sleep(3 * time.Second)
//	c <- "msg1, delay 3sec"
//}
//
//func message2(c chan string) {
//	time.Sleep(4 * time.Second)
//	c <- "msg2, delay 4sec"
//}
//
//func message3(c chan string) {
//	time.Sleep(2 * time.Second)
//	c <- "msg3, delay 2sec"
//}
