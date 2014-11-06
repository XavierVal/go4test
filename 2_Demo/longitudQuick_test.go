package main

import (
	"testing"
	"testing/quick"
)

// Scenario: Este test comprobará que la longitud del body de nuestra peticiones
// y la respuesta de longitud calculada en el server coinciden.

// Nos vamos a ayudar de la funcion "quick" que genera un conjunto aleatorio
// de entradas para verficar si es cierto un supuesto,
// generando un error si el supuesto no se cumple.

func TestLongitud(t *testing.T) {
	const url = "http://localhost:7777"
	//Establecemos el supuesto que debe cumplirse
	supuesto := func(textoDePrueba string) bool {
		mapa, err := enviaPeticion(url, textoDePrueba)
		if err != nil {
			return false // esta cadena no ha ido bien
		}

		// Extraemos la longitud de la respuesta obtenida del server
		var longitudRespuesta = mapa["longitud"].(float64)

		// Deberia devolver la longitud de la cadena que hemos enviado
		var resultadoEsperado = len(textoDePrueba)

		// Validamos si lo obtenido coincide con lo esperado
		return int(longitudRespuesta) == resultadoEsperado
	}
	if err := quick.Check(supuesto, nil); err != nil {
		t.Error(err) //en caso de error muestralo
	}
}
