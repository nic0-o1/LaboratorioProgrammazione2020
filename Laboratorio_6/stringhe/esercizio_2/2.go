package main

import "fmt"

func main() {
	var n int
	var toRepeat string

	fmt.Scan(&n, &toRepeat)

	for i := 0; i < n; i++ {
		fmt.Print(toRepeat)
		if i < n-1 {
			fmt.Print("-")
		}
	}

}
