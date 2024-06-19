# Aprendiendo Go

## Creacion de modulo

Para la creacion de un modulo en `Go` usaremos el comando

```
// go mod init `nombre del modulo a crear`

go mod init hello
```

## Paquetes

Un paquete te ayuda a llevar funciones a otro lugar
Uno de los casos es el `fmt` que se instala con `Go`.
Una declarado el paquete vamos a introducile funcionamiento

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

## lista de comando

para poder ver la lista de comandos que puede ejecutar go solo necesitas introducir en consola.

```
go help
```

## Tipos de datos de Go

### Integers

- int8: -128 a 127
- int16: -32768 a 32767
- int32: -2147483648 a 2147483647
- int64: -9223372036854775808 a 922337203685477580
- uint8: 0 a 255
- uint16: 0 a 65535
- uint32: 0 a 4294967295
- uint64: 0 a 18446744073709551615
- rune: alias para int32
- byte: alias para uint8
- int: depende de la arquitectura del sistema
- uint: depende de la arquitectura del sistema

### Floats

- float32: 1.4e-45 a 3.4e38
- float64: 4.9e-324 a 1.8e308

### Complex

- complex64: float32 + float32i
- complex128: float64 + float64i

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
