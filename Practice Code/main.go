//https://golang.org/
package main

import "fmt"

func main() {	
	fmt.Println("This is my program")
	testPointers()
}

func testPointers(){
	// My way
	fmt.Println("Typical Sqaure", sqaure(5))

	// Go way
	myNum:= 6
	goSquare(&myNum)
	fmt.Println("Go Sqaure", myNum)
}