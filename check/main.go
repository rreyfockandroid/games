package main

import "fmt"

func main() {
	fmt.Println("fire")

	i := 5

	for j := 0; j < 10; j++ {
		i--
		fmt.Println(i)
	}
}
