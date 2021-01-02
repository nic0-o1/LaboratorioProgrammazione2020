package main

import(
	"os"
	"bufio"
	"fmt"
)


func main(){
	fmt.Println(count("operazioni_1.txt"))
}

func count(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
	   return 0, err
	}
	defer file.Close()
 
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
	   fmt.Println(scanner.Text())
	}
	return count, nil
 }