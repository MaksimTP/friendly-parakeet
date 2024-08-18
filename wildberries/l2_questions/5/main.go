package main

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{

	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("ok")
}
