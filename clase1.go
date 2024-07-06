package main

import "fmt"
func main(){
	var a int = 10
	var b, c int // Puedo crear las variables mas de una en un sitio
	d := "Victor"
	b = 31
	c = 23
	var e int // e = 0 si no se le decladra valor su valor por default sera 0 o depende del tipo creo que es un falsy si no se define
	var f string // f = ""

	fmt.Println("Hello world", a, b, c,d,e,f)
// Punteros
	var x int = 1
	var y int //0
	var ip *int // nil es un puntero a int

	fmt.Println("Pointers",x,y,ip)
	ip = &x // ahora es un puntero a x => 0xc00000a0a8
	y = *ip // y ahora es 1 
	fmt.Println("Pointers 2",ip, y)

	var mySlece = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, s := range mySlece {
		println(`mi numero es`, s, "y su indice es", i)
	}
	// se puede omitir el un valor usando `_`
	/*
	for _, s := range mySlece {
	*/
}