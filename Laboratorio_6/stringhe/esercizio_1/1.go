package main

import "fmt"

func main(){
	var size int

	fmt.Scan(&size)

	if size <=0{
		fmt.Println("Dimensione non sufficiente")
	}else{
		stringa:=""
		for i:=0;i<size;i++{
			stringa+="*"
			fmt.Println(stringa)
		}
	}

}