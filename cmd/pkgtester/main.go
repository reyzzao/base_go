package main

import "fmt"

func hello() string{
	return "hello pkgtester"
}

func main() {
	fmt.Println(hello())
}
