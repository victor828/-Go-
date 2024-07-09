package main

import (
	"fmt"
	"strings"
)
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
	var z int = x
	x = 2
	fmt.Println("Pointers 2",ip, y,z)

	var mySlece = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, s := range mySlece {
		println(`mi numero es`, s, "y su indice es", i)
	}

	nombre := "Victor"
	fmt.Println(strings.Split(nombre, ""))
	println(strings.HasSuffix(nombre, "r"))
} 
