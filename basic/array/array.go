package main

import "fmt"

func main() {
	var names [5]string

	x := [5]int{10, 11, 12, 12, 2}

	fmt.Println("i have :" x[1])


	names[0] = "messi"
	names[2] = "ronaldo"

	fmt.Println("My name is: " + names[0])
	fmt.Println("I dont know you" + names[1])

}
