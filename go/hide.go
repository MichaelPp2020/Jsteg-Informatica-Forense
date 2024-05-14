package main

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"lukechampine.com/jsteg"
)

func main() {
	// Abre el archivo JPEG original
	inFile, err := os.Open("original.jpg")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	// Lee el contenido del archivo JPEG
	srcData, err := ioutil.ReadAll(inFile)
	if err != nil {
		panic(err)
	}

	// Decodifica la imagen JPEG
	srcImage, _, err := image.Decode(bytes.NewReader(srcData))
	if err != nil {
		panic(err)
	}

	// Crea un buffer de bytes para escribir los datos ocultos
	var buf bytes.Buffer

	// Mensaje a ocultar
	message := "Mensaje oculto Jsteg para informatica Forense Alumno: Michael Palacios! "

	// Oculta el mensaje en la imagen
	err = jsteg.Hide(&buf, srcImage, []byte(message), nil)
	if err != nil {
		panic(err)
	}

	// Escribe la imagen con el mensaje oculto en un archivo
	outFile, err := os.Create("stego.jpg")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = outFile.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}

	fmt.Println("Mensaje oculto en la imagen.")
}
