package main

import "fmt"
func main(){
	var a int = 10

	fmt.Println("Hello world", a)


	var mySlece = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, s := range mySlece {
		println(`mi numero es`, s, "y su indice es", i)
	}
	// se puede omitir el un valor usando `_`
	/*
	for _, s := range mySlece {
	*/
}