package main

import {
	"fmt"
	"strings"
}

func main() {
	for i := 0; i < 5; i++ {
		go printA()
	}
	fmt.Println("123")
}

func printA () {
	fmt.Println("abc")
}
