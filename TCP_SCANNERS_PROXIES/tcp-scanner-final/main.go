package main 

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int){
	for p:=range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <-0 // indicate failed
			continue
		}
		conn.Close()
		results <- p // port number that successed
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i:=0; i<cap(ports); i++ {
		// create workers
		go worker(ports, results)
	}


	// feed jobs 
	go func() {
		for i:=1; i<=1024; i++ {
			ports<-i
		}
	}()

	// no need to wg.Wait()
	// if # of results == # of jobs, then finished, just close
	for i:=0; i<1024; i++ {
		port := <-results
		if port != 0 {
			openports=append(openports, port)
		}
	}

	close(ports)
	close(results)
	// sort results
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
