package main

import (
	"./triangolo"
	"fmt"
	"math/rand"
	"math"
	"time"
	"os"
	"strconv"
)
const (
	RANGEMIN = 100
	RANGEMAX = 1000
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	triangoli := GeneraTriangoli(n)

	max,min := MinMax(triangoli)

	fmt.Println("Triangolo con area maggiore = ",max.String())
	fmt.Println("Triangolo con area minore = ",min.String())
}

func MinMax(tn []*triangolo.Triangolo)(tMax,tMin triangolo.Triangolo){
	AreaMax, AreaMin := 0.0,math.MaxFloat64

	for _,t := range tn[1:]{
		area := t.Area()
			if AreaMax < area {
				tMax = *t
				AreaMax = area
			}
			if AreaMin > area {
				tMin = *t
				AreaMin = area
			}
	}
	return
}

func GeneraTriangoli(n int) (tn []*triangolo.Triangolo) {
	for i := 0; i < n; {
		rand.Seed(time.Now().UnixNano())
		l1 := RANGEMIN + rand.Float64() * (RANGEMAX - RANGEMIN)
		l2 := RANGEMIN + rand.Float64() * (RANGEMAX - RANGEMIN)
		l3 := RANGEMIN + rand.Float64() * (RANGEMAX - RANGEMIN)

		if t, err := triangolo.NuovoTriangolo(l1, l2, l3); err == nil {
			tn = append(tn, t)
			
			i++
		}
	}
	return
}