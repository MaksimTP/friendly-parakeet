package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	tmr := time.NewTimer(duration)
	<-tmr.C
}

func main() {
	fmt.Println("hello")
	sleep(1 * time.Second)
	fmt.Println("bye")
}
