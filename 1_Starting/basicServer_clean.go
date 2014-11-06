// server_clean.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello World, testeandoSoftware.com!")

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		log.Printf("%+v\n\n", request)
		fmt.Fprintln(response, "OK", "Received at ", time.Now())
	})

	log.Fatal(http.ListenAndServe(":7777", nil))
}
