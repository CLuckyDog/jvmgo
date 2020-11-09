package main

import "fmt"

func main() {
	test(1, 2, 3, 4, 5, 9)
}

func test(args ...int) {
	fmt.Println("len(args):", len(args))
	// test(args[2:]...)
}

func test(name string, args ...int) {
	for _, data := range args {
		fmt.Println("data = ", data)
	}
}
