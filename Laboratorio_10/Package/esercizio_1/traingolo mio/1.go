package main

import (
	"fmt"
	"os"
	"strconv"

	"./triangolo"
)

func main() {
	l1, _ := strconv.Atoi(os.Args[1])
	l2, _ := strconv.Atoi(os.Args[2])
	l3, _ := strconv.Atoi(os.Args[3])

	if tria, err := triangolo.NuovoTriangolo(float64(l1), (float64(l2)), float64(l3)); err == nil {
		fmt.Println("Perimetro triangolo = ", tria.Perimetro())
		fmt.Println("Area triangolo =", tria.Area())
	} else {
		fmt.Println("Errore:", err)
	}

}
