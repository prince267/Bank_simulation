// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import "fmt"
import "time"

func banksimmulation(id int, customers <-chan int, results chan<- int) {

	for j := range customers {
		fmt.Println("cashier :", id, " coustomers ", j, " started")
		time.Sleep(3 * time.Second)
		fmt.Println("cashier :", id, " coustomers ", j, " stopped")
		results <- 1
	}
}

func main() {
	customers := make(chan int, 100)
	results := make(chan int, 100)
	for c := 1; c <= 3; c++ {
		go banksimmulation(c, customers, results)
	}
	for j := 1; j <= 5; j++ {
		customers <- j
	}
	close(customers)
	for a := 1; a <= 5; a++ {
		<-results
	}
}
