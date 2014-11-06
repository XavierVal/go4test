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

// La concurrencia no es problema independientemente del numero de peticiones
type Contador struct {
	value      int
	sync.Mutex // Utilidad de "bajo" nivel.
	// Canales sería otra opción a utilizar
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
		// Leemos el cuerpo. Cuando salgamos, que se cierre el Reader
		defer request.Body.Close()
		// lo leemos entero, sin más ceremonia
		body, err := ioutil.ReadAll(request.Body)
		// Sin excepciones que rompan el flujo "sorpresivamente", tratamos el error cuando aparece
		if err != nil {
			http.Error(response, "Leyendo body", 500)
			return
		}
		// se prepara la respuesta con un mapa string->cualquier tipo
		mapa := make(map[string]interface{})
		mapa["cabecera"] = request.Header       // un objeto complejo
		mapa["contador"] = contador.IncAndGet() // incrementamos y recuperamos
		// con peticiones concurrentes sin confictos
		mapa["longitud"] = len(body) // calculamos la longitud del cuerpo

		// imprimimos en el log que se ha recibido una petición y el cuerpo del mismo
		if string(body) != "" {
			log.Printf("[%d] Recibido: %s", contador.value, string(body))
		} else {
			log.Printf("[%d] Sin cuerpo", contador.value)
		}

		// establecemos el content-type de la respuesta
		response.Header().Set("content-type", "application/json")
		response.Header().Set("help", "La respuesta de este servidor incluye las cabeceras que se han usado en la petición, el contador de peticiones recibidas en el server y el tamaño del cuerpo del mensaje recibido")

		// Para el ejemplo preferimos un JSON indentado con tres espacios para que
		// sea más legible, MarshalIndent y strings.Repeat hacen el trabajo
		json, err := json.MarshalIndent(mapa, " ", strings.Repeat(" ", 3))
		if err != nil {
			http.Error(response, "Generando JSON", 500)
			return
		}
		response.Write(json)
	})
	log.Fatal(http.ListenAndServe(":7777", nil))
}
