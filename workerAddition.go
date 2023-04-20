package main
import (
"fmt"
"time"
)

func main() {
start := time.Now()
jobs := make(chan int, 100000)
results := make(chan int, 50000)

go worker(jobs, results)
go worker(jobs, results)

	for i := 0; i < 100000; i=i+2 { // put jobs in the channel for the works to do
	jobs <- i
	jobs<-(i+1)
	}
	close(jobs) // close the channel because you are done putting jobs in
	for j := 0; j < 50000; j++ { // print out the answers for the jobs that the workers came up with
	fmt.Println(j, " ", <-results)
	}
	elapsed := time.Since(start)
	fmt.Printf("It took %s", elapsed)
		}

func worker(jobs <-chan int, results chan<- int) {
for n := range jobs {
results <- sum(n,<-jobs)
}
}

func sum(n int, m int)int { // compute the fib for the number that is sent it
	return n+m
	
	

	}
