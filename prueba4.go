package main
import "fmt"

var c, python, java bool
const pi float32 = 3.14

func intercam(x, string, z string) (string, string)/*aca lo que quiere que devolver,si es uno no hace falta parentesis,
cons*/{
	return x, z
}
//inferencia de tipos: hay casos en los que haya variables que devuelve una funcion que el compilador infiere que tipo
// de datos que se indica con ':' antes de asignar 
//en go las funciones puede devolver n cantidad de valores

func main(){
	a,b := intercambiar('wojnvjka'. 'navjnav')
	fmt.Println(a, b)
}

//se puede usar un solo valor(normalmente devuelve el resultado + el errors) de lo que devuelve una funcion, se soluciona con _
func main(){
	_b := intercambiar('wojnvjka'. 'navjnav')
	fmt.Println(b)
}

//para que omita errores de imports sin uso se puede colocar _ para que el copilador omita el error