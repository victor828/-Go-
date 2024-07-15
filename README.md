# Aprendiendo Go

## Comandos

### go build

Se usa para la copilacon de un programa

### go dock

Imprima la documentacion de un paquete

### go fmt

Formatea los archivos de codigo fuente

### go get

descarga los paquetes y los installa

### go list

Muestra la lista de todos los paquets installados.

### go run

Copila el archivo go y ejecuta el ejecutable.

### go test

correo los test de los archivos que terminan en `_test.go`

### go help

para poder ver la lista de comandos que puede ejecutar go solo necesitas introducir en consola.

## Creacion de modulo

Para la creacion de un modulo en `Go` usaremos el comando

```
// go mod init `nombre del modulo a crear`

go mod init hello
```

## Paquetes

Un paquete te ayuda a llevar funciones a otro lugar
Uno de los casos es el `fmt` que se instala con `Go`.
Una declarado el paquete vamos a introducile funcionamiento en este caso para `main`

```go
func main(){
    fmt.Println("Hola mundo")
}
```

con esto tomamos al paquete `fmt` y usamos uno de sus `metodos`
ahora vamos a correr el paquete.
Para ejecutarlo usamos el comando `go run + nombre del paquete`.

```
go run hello
```

## Tipos de datos de Go

### Integers

Son numeros enteros

- int8: -128 a 127
- int16: -32768 a 32767
- int32: -2147483648 a 2147483647
- int64: -9223372036854775808 a 922337203685477580

### Floats

Numeros decimales.

- float: 1.5
- float32: 1.4e-45 a 3.4e38
- float64: 4.9e-324 a 1.8e308

### Boolean

- bool: true o false

### String

- string: cadena de caracteres
- rune: caracter individual

#### Interpolacion en go

Se puede hacer una interpolacion en go de la siguietne manera

```go
package main
import "fmt"
func main() {
    name := "John"
    age := 30
    fmt.Printf("My name is %s and I am %d years old.\n", name
    , age)
    }
```

### Arrays

- [n]T: array de n elementos de tipo T
  ejemplo:

```go
var arr = [5]int [ 1,2,3,4,5]
```

### Slices

- []T: slice de elementos de tipo T
  ejemplo:

```go
ejemplo
var slice = []int [ 1,2,3,4,5]
```

se dierencia de un array en que no tiene un tamaño fijo y se puede agregar o eliminar elementos dinamicamente.

### Structs

- struct {}: estructura vacía
- struct {campo1 tipo1; campo2 tipo2; ...}: estructura con campos

### Pointers

- \*T: puntero a un valor de tipo T

### Functions

- func nombreFuncion(parametros) tipoRetorno: función con nombre y parámetros

### Interfaces

- interface {}: interfaz vacía
- interface {metodo1; metodo2; ...}: interfaz con métodos

### Maps

- map[K]V: mapa que relaciona claves de tipo K con valores de tipo V

### Channels

- chan T: canal que envía y recibe valores de tipo T
- chan<- T: canal que solo envía valores de tipo T
- <-chan T: canal que solo recibe valores de tipo T

### Enums

- No hay un tipo de dato enum en Go, pero se pueden utilizar constantes para implementar

### Nil

- nil: valor nulo para punteros, interfaces, slices, maps y channels

## Bucles

En `Go` existen dos tipos de bucles `for` y `range`

### For

El bucle `for` es similar al de otros lenguajes de programación.

```go
for var i:= 0 ; i > 5 ; i++ {
    println(i)
}
```

### Range

El bucle `range` es un bucle que itera sobre una colección de elementos.

```go
frutas := []string{"manzana", "pera", "naranja"}
for _, fruta := range frutas {
    println(fruta)
    }

```

En este ejemplo estamos iterando sobre un slice de strings y estamos imprimiendo cada elemento.
El `_` es un placeholder que se utiliza para ignorar el índice del elemento.
Si queremos acceder al índice podemos hacerlo de la siguiente manera.

```go
frutas := []string{"manzana", "pera", "naranja"}
for i, fruta := range frutas {
    println(i, fruta)
    }
```

En este caso estamos imprimiendo el índice y el valor de cada elemento.

# CodeWars

## Convertidor de tallas

### Description

You have clothes international size as text (xs, s, xxl).
Your goal is to return European number size as a number from that size .

Notice that there is arbitrary amount of modifiers (x), excluding m size, and input can be any string.

Linearity
Base size for medium (m) is 38.
(then, small (s) is 36, and large (l) is 40)

The scale is linear; modifier x continues that by adding or subtracting 2 from resulting size.

Invalid cases
Function should handle wrong/invalid sizes.

Valid input has any base size (s/m/l) and any amount of modifiers (x) before base size (like xxl).
Notice that you cannot apply modifier for m size, consider that case as invalid!
Anything else is disallowed and should be considered as invalid (None for Python, 0, false for Go, null for JavaScript).

Invalid examples: xxx, sss, xm, other string

Valid Examples
xs: 34
s: 36
m: 38
l: 40
xl: 42

### Resolucion

```Go
func SizeToNumber(size string) (int, bool) {
  var rsize int
  var vrdd bool
  // Generamos un map con las tallas y su numero
   talls := map[string]int{"s": 36, "m":38, "l":40}
   //Verificamos si termina con m y comienza con x
  if  strings.HasSuffix(size, "m" ) && strings.HasSuffix(size, "x") {
		rsize = 0
		vrdd = false

  } else if strings.HasPrefix(size, "x") {
   //Aqui verificaremos que size comience con x
	if strings.HasSuffix(size,"s") {
	//si el texto termina en s
		var r = strings.Count(size, "x")
	rsize = talls["s"] -  2*r
	} else if strings.HasSuffix(size, "l") {
	 //Si el tecto termina en l
		var r = strings.Count(size, "x")
		rsize = talls["l"] +  2*r
	}
  } else {
	  switch size {
	  //Aqui estan los casos normales
		  case "s": {
			  rsize = talls["s"]
			  vrdd = true
			}
			case "m": {
				rsize = talls["m"]
				vrdd = true
			}
			case "l": {
				rsize = talls["l"]
				vrdd = true
			}
			default: {
				rsize = 0
				vrdd = false
			}
		}
	}
  return rsize, vrdd
}
```

### Otra resulucion Vista

```GO
import (
  "strings"
)

func SizeToNumber(size string) (int, bool) {
    //Guarda el numero de x en el string
  x_count := strings.Count(size, "x")

  //Si el string esta vacio o la cantidad de x es menor al len de size -1
  if size == "" || x_count < len(size)-1 {
    return 0, false
  }

  //
  switch size[len(size)-1] {
    case 's': return 36 - x_count*2, true
    case 'l': return 40 + x_count*2, true
    case 'm': {
      if x_count == 0 {
        return 38, true
      } else {
        return 0, false
      }
    }
    default: return 0, false
  }
}
```

```GO
func SizeToNumber(size string) (int, bool) {
	var count int
	m := map[string]int{
		"x":    2,
		"xxxs": 30,
		"xxs":  32,
		"xs":   34,
		"s":    36,
		"m":    38,
		"l":    40,
		"xl":   42,
		"xxl":  44,
		"xxxl": 46,
	}
	_, found := m[size]
	if found {
		return m[size], true
	}
	slc := strings.Split(size, "")
	uniqWords := []string{}
	uniq := map[string]bool{}
	for _, v := range slc {
		uniq[v] = true
	}
	for k, _ := range uniq {
		uniqWords = append(uniqWords, k)
	}
	if len(slc) > 1 && (slc[0] == "s" || slc[0] == "m" || slc[0] == "l") {

		return 0, false
	}
	if len(uniqWords) > 2 || (len(uniqWords) >= 2 && strings.Contains(size, "m")) {
		return 0, false
	} else {
		if strings.Contains(size, "s") {
			m["x"] = -2

```

## Even or Odd 'Par o Inpar'

### Description

Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.

### Operacion

```GO
func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "Even"
	}
  return "Odd"
}
```

## Clock

### Description

Clock shows h hours, m minutes and s seconds after midnight.
Your task is to write a function which returns the time since midnight in milliseconds.

### Operation

```GO
func Past(h, m, s int) int {
	return (h*3600 + m*60 + s)*1000
}
```
