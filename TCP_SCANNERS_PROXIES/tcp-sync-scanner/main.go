package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup){
	for p := range ports {	// consume what is in the ports
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	// create a channel with buffer size of 100
	ports := make(chan int, 100)
	var wg sync.WaitGroup 
	for i:=0; i<cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i:=1; i<=1024; i++ {
		wg.Add(1)	// increment counter
		ports<-i	// send port number into ports
	}
	// wait for the counter to become zero
	wg.Wait()
	// close ports 
	close(ports)
}
