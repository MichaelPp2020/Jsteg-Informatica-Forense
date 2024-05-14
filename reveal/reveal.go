package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"lukechampine.com/jsteg"
)

func main() {
	// Abre el archivo JPEG que contiene datos ocultos
	inFile, err := os.Open("stego.jpg")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	// Lee el contenido del archivo JPEG
	stegoData, err := ioutil.ReadAll(inFile)
	if err != nil {
		panic(err)
	}

	// Crea un buffer de bytes y escribe los datos del archivo JPEG
	stegoBuffer := bytes.NewBuffer(stegoData)

	// Crea un lector de bytes a partir del buffer
	stegoReader := bytes.NewReader(stegoBuffer.Bytes())

	// Extrae los datos ocultos de la imagen
	message, err := jsteg.Reveal(stegoReader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Mensaje oculto: %s\n", message)
}
