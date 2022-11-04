package main

import "fmt"

func main() {
	fmt.Print("enter your name: ")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("your name is", name)
}
