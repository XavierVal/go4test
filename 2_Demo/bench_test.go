package main

import (
	"net/http"
	"testing"
)

// Scenario: Este test medirá el tiempo de respuesta del server vs operaciones

// Con go test -test.bench=. se ejecutan todos los benchmark
// Se puede elegir cuales mediante una expresion regular
// Para este "extraño" caso, cuidado con el limite de 'open files'

func BenchmarkEnviarPeticionesGet(b *testing.B) {
	// Es más para benchmark de funciones que para medir las respuestas de una api...
	// pero a modo de ilustración nos puede valer
	// Debemos hacer b.N iteraciones para que pueda medir adecuadamente
	// el valor N nos lo pasa el entorno

	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://127.0.0.1:7777")
		if err != nil {
			b.Fatal(err) // Si hay error, paramos el test
		}
		resp.Body.Close()
	}
}
