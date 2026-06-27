package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World NOT LOL"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
