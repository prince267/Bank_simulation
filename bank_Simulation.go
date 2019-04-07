package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func banksimmulation(id int, customers <-chan int, results chan<- int, coustomerTime int) {

	for j := range customers {
		fmt.Println((time.Now()).Format("2006-01-02 15:04:05")+" --> Cashier ", id, ": Customer ", j, " Started")
		time.Sleep(time.Duration(coustomerTime) * time.Second)
		fmt.Println((time.Now()).Format("2006-01-02 15:04:05")+" --> Cashier ", id, ": Customer ", j, " Completed")
		results <- 1
	}
}

func main() {
	numCashier := os.Args[1]
	numCoustomers := os.Args[2]
	timePerCustomer := os.Args[3]
	NoOfcashier, _ := strconv.Atoi(strings.Split(numCashier, "=")[1])
	NoOfcoustomers, _ := strconv.Atoi(strings.Split(numCoustomers, "=")[1])
	coustomerTime, _ := strconv.Atoi(strings.Split(timePerCustomer, "=")[1])

	customers := make(chan int, NoOfcoustomers)
	results := make(chan int, NoOfcoustomers)
	fmt.Println((time.Now()).Format("2006-01-02 15:04:05") + " --> Bank Simulation Started")
	for c := 1; c <= NoOfcashier; c++ {
		go banksimmulation(c, customers, results, coustomerTime)
	}
	for cus := 1; cus <= NoOfcoustomers; cus++ {
		customers <- cus
	}
	close(customers)
	for res := 1; res <= NoOfcoustomers; res++ {
		<-results
	}
	fmt.Println((time.Now()).Format("2006-01-02 15:04:05") + " --> Bank Simulation Ended")
}
