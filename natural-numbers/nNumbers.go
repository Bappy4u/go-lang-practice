package main

import "fmt"

func main() {
	//2+4+6+........+200

	// lets find i, an = a+(n-1)d
	// 2i = 200
	// n = 100
	fmt.Print("Enter nth number: ")
	var n int
	fmt.Scanf("%d", &n)
	fmt.Print("Here is the sum of 2+4+6+......+", n*2, " is: ")
	sum := n * (n + 1)
	fmt.Println(sum)

}
