package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Scanf("%d", &n)

	res := make([]int, n) 
	fmt.Println(res)
	for i:=0; i<n; i++{
		fmt.Scanf("%d", &res[i])
		
	}
	
	fmt.Println(res)

}