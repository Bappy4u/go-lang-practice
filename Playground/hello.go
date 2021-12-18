package main

import (
	"fmt"
)

func main() {
	var n, item, sum, val int
	fmt.Scanf("%d", &n)

	res := make([]int, 0)
	for i := 0; i < n; i++ {
		sum = 0

		fmt.Scan(&item)
		for j := 0; j < item; j++ {
			fmt.Scan(&val)

			sum += val * val
		}
		res = append(res, sum)

	}

	fmt.Println(res)

}
