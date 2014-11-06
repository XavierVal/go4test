// 1_basicServer.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello World, testeandoSoftware.com!") //saludamos por la salida estándar

	// Añadimos un manejador para el path "/", una simple función hace el trabajo
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		//registramos la peticion, con los nombres de los campos para que sea más cómodo (%+v)
		log.Printf("%+v\n\n", request)

		// Se escribe en el Writer de la respuesta como en cualquier otro Writer
		// (fichero, socket, memoria ... cualquiera que implemente el interfaz
		// con un método Write(data []byte)
		fmt.Fprintln(response, "OK", "Received at ", time.Now())
	})

	//Si no conseguimos arrancar, dejamos una traza al menos
	log.Fatal(http.ListenAndServe(":7777", nil))
}

// Posible error si ya hay un servidor levantado en ese puerto
// $ go run 1_basicServer
// Hello World!
// 2014/10/22 18:43:54 listen tcp :7777: bind: address already in use
// exit status 1
