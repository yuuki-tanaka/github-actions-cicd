package main

import "fmt"

func Greeting() string {
	return "Hello"
}

func Greeting2(g string) string {
	return g + "World"
}

func main() {
	fmt.Println(Greeting())
}
