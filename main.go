package main

import (
	"errors"
	"fmt"
)

const name string = "Tom"

func main() {
	showError()
	showError2()
}

func callTom() {
	fmt.Println("my name is Tom")
}

func showError() error {
	fmt.Println("my name is Tom")
	return errors.New("it is a fault")
}

func showError2() error {
	fmt.Println("my name is Jerry")
	return errors.New("it is a fault2")
}
