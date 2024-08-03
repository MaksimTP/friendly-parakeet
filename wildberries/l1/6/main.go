package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// func main() {
// 	go func() {
// 		fmt.Println("Hello world")
// 	}()
// }

// func main() {
// 	quit := make(chan int)

// 	go func() {
// 		for {
// 			select {
// 			case <-quit:
// 				return
// 			default:
// 				fmt.Println("Hello")
// 			}
// 		}
// 	}()

// 	quit <- 1
// }

func main() {
	forever := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				forever <- 1
				return
			default:
				fmt.Println("hello")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(1500 * time.Millisecond)
		cancel()
	}()

	<-forever
	fmt.Println("finish")
}
