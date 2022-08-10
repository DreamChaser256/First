package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

const name string = "test"

func testfun() error {
	return errors.New("this")
}
func main() {
	str := ReadFile()
	fmt.Println(str)
	fmt.Println("----end----")
	testfun()
}

func ReadFile() string {
	f, err := ioutil.ReadFile("/home/test")
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}
