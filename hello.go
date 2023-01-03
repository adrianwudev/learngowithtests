package main

import "fmt"

const enlishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return enlishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
