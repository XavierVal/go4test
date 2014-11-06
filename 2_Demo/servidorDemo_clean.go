// server.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Servidor que recibe peticiones http y responde con datos varios como
// - cabeceras que se han usado en la petición
// - contador de peticiones recibidas en el server
// - tamaño del cuerpo del mensaje recibido

type Contador struct {
	value int
	sync.Mutex
}

// Funcion para incremetar el contador
func (c *Contador) IncAndGet() int {
	c.Lock()         // Acceso exclusivo al contador
	defer c.Unlock() // Liberamos el mutex
	c.value++
	return c.value
}

func main() {
	log.Println("Server para testeandoSoftware.com en Go:")
	var contador = &Contador{value: 0}
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(response, "Leyendo body", 500)
			return
		}
		mapa := make(map[string]interface{})
		mapa["cabecera"] = request.Header
		mapa["contador"] = contador.IncAndGet()
		mapa["longitud"] = len(body)

		if string(body) != "" {
			log.Printf("[%d] Recibido: %s", contador.value, string(body))
		} else {
			log.Printf("[%d] Sin cuerpo", contador.value)
		}

		response.Header().Set("content-type", "application/json")

		json, err := json.MarshalIndent(mapa, " ", strings.Repeat(" ", 3))
		if err != nil {
			http.Error(response, "Generando JSON", 500)
			return
		}
		response.Write(json)
	})
	log.Fatal(http.ListenAndServe(":7777", nil))
}
