package main

import "fmt"

func main() {
	c := "ciao, come va?"
	
	fmt.Println(c[0],c[len(c)-2],c[len(c)-1],len(c))
}
