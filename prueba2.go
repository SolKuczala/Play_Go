package main

import (
	"fmt"
	"math/rand"
)

// todos los archivos de una carpeta deben permanecer al mismo paquete

func mainDos() {
	fmt.Println("mi numero blah", rand.Intn(10)) //println(print line)
}

//los archivos comparten visibilidad, como una variable, que se puede usar desde otro archivo
//variables const e imports no usados no compila
//y para exportar algo para otro paquete se usa con mayuscula
