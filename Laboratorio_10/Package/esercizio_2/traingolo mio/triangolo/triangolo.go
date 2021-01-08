package triangolo

import "errors"
import "math"
import "fmt"

type Triangolo struct {
    lato1, lato2, lato3 float64
}

func NuovoTriangolo(l1, l2, l3 float64) (t *Triangolo, err error){
	if l1 + l2 > l3 && l1 + l3 > l2 && l2 + l3 > l1{
		return &Triangolo{l1,l2,l3}, nil
	}else{
		return nil, errors.New("impossibile creare un triangolo con le misure specificate")
	}
}

func (t Triangolo) Perimetro() float64{
	return t.lato1 + t.lato2 + t.lato3
}

func (t Triangolo) Area() float64{
	p := t.Perimetro() / 2
	return math.Sqrt(p * (p-t.lato1) * (p-t.lato2) * (p-t.lato3))
}

func (t Triangolo) String() string{
	return fmt.Sprintf("Triangolo con lati %.2f, %.2f, %.2f", t.lato1, t.lato2, t.lato3)
}