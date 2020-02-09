package main

import "fmt"

const HelloHead = "Hello, "

func Hello(name string) string {
	return HelloHead + name
}

func main() {
	fmt.Println(Hello("WuHan"))
}
