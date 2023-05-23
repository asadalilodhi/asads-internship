package main

import "fmt"

const prefix = "Hi, "
const suffix1 = "Asad"
const suffix2 = "Fatima"

func Hello(name string) string {
	return prefix + name
}

func Hi(name string) string {
	return prefix + name
}

func main() {
	fmt.Println(Hello(suffix1))
	fmt.Println(Hi(suffix2))
}
