Manual para Ocultar y Revelar Mensajes con Jsteg en Go
Requisitos Previos:
Tener instalado Git para Windows.
Tener instalado Go en tu sistema Windows.
Instalar Go.
1.	Ingresaremos al siguiente link: https://go.dev/dl/ y descargaremos la versión para Windows 
 

2.	Como cualquier tipo de programas instalaremos normalmente Go y configuramos las variables de entorno.
 

3.	Ahora descargaremos el código fuente de Jsteg.
 

Ocultar Mensaje 

4.	Haremos dos carpetas 
 

5.	Abre visual Studio Code y copia el siguiente código en un nuevo archivo llamado hide.go:
 

6.	Preparar la Imagen Original:

Asegúrate de tener una imagen JPEG llamada original.jpg.
Esta imagen será la que contenga el mensaje oculto.

 
7.	Ejecutar el Programa para Ocultar el Mensaje:

En la misma carpeta que hide.go, abre una terminal.
Ejecuta el siguiente comando para compilar y ejecutar el programa hide.go:

go run hide.go
Esto generará una nueva imagen llamada stego.jpg con el mensaje oculto.

 













Pasos para Revelar el Mensaje Oculto

1.	En la carpeta reveal que creamos anteriormente 
2.	Abre visual Studio Code y copia el siguiente código en un nuevo archivo llamado reveal.go
 

3.	En la carpeta reveal ingresamos la imagen que nos generó el código anterior 
 

4.	Ejecuta el siguiente comando para compilar y ejecutar el programa reveal.go:

go run reveal.go

5.	Esto mostrará el mensaje oculto extraído de la imagen stego.jpg.
 
Siguiendo estos pasos, podrás ocultar y revelar mensajes dentro de imágenes JPEG utilizando Jsteg y Go en tu sistema Windows.

Aquí les dejos los códigos utilizado en esta practica.
