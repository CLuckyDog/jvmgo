package main

import "fmt"

type Book struct {
	name string
	count int
}

func main() {
	//test1(1, 2, 3, 4, 5, 9)
	//sb := make([]Book,5)
	//for i := 1; i < 5; i++ {
	//	sb[i] =readBook(sb)
	//}
}

//func readBook(b Book) Book{
//	fmt.Printf("正在读书: %v\n",b.name)
//}

func test1(args ...int) {
	fmt.Println("len(args):", len(args))
	// test(args[2:]...)
}

func test(name string, args ...int) {
	for _, data := range args {
		fmt.Println("data = ", data)
	}
}
