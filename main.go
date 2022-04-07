package main

import "fmt"

func main() {
	fmt.Println("Hello")
	foo()
	fmt.Println("something else")
	
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	bar()
}

func foo() {
	fmt.Println("Im in foo")
	count := 4
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}

func bar() {
	fmt.Println("and then we exited")
}