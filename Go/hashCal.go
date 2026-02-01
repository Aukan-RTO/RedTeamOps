package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
)

// Utilidad de línea de comandos
// para calcular los hashes MD5, SHA1 y SHA256 de un archivo.
// Uso: hashcal -F <archivo>

type Options struct {
	// InputFile es la ruta del archivo que se va a procesar.
	InputFile string
}

// callOptions parsea las banderas de la línea de comandos y devuelve
// una estructura Options con los valores obtenidos.
func callOptions() *Options {
	// Mensaje de ayuda que se muestra al ejecutar sin parámetros adecuados.
	fmt.Print("Este programa permite calcular los siguientes hash para un archivo: MD5, SHA1, SHA256\n\n")

	// Definimos la bandera -F para indicar el archivo de entrada.
	inputFile := flag.String("F", "", "Ingresar archivo ")
	flag.Parse()

	return &Options{
		InputFile: *inputFile,
	}
}

// main es el punto de entrada del programa.
func main() {
	menu := callOptions()

	// Verificar que el usuario haya ingresado un archivo.
	if menu.InputFile == "" {
		log.Fatal("[!] Debe ingresar un archivo.\nEjemplo: hashcal -F <File>")
	}

	// Leer todo el contenido del archivo en memoria.
	entradaData, err := os.ReadFile(menu.InputFile)
	if err != nil {
		log.Fatal("[!] Error al leer el archivo.")
	}

	// Calcular los hashes.
	md5sum := md5.Sum(entradaData)
	sha1sum := sha1.Sum(entradaData)
	sha256sum := sha256.Sum256(entradaData)

	// Mostrar los resultados en hexadecimal.
	fmt.Printf("MD5: %x\n", md5sum)
	fmt.Printf("SHA1: %x\n", sha1sum)
	fmt.Printf("SHA256: %x\n", sha256sum)
}
