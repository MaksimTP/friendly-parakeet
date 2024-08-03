package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	numWorkers := 0
	fmt.Println("Choose number of workers: ")
	fmt.Scan(&numWorkers)

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)

	ch := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {

		go func() {

			for v := range ch {
				select {
				case <-sigs:
					fmt.Println("SIGINT!")
					close(ch)
					os.Exit(1)
				default:
					break
				}
				fmt.Printf("Num:\t%2d, Worker:\t%2d\n", v, i+1)
			}
		}()
	}

	for {
		ch <- rand.Int() % 100
	}

}
