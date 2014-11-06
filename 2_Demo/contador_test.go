package main

import (
	"testing"
)

// Scenario: Este test comprobará que el incremento del contador de requests recibidas

// Aunque esta no es exactamente la finalidad del framework de pruebas de Go, podemos usarlo
// como ilustración de la potencia de "testing"

// Las funciones que empiezan con Test se ejecutan como un caso de prueba
func TestContador(t *testing.T) {
	const url = "http://localhost:7777"

	mapa, err := enviaPeticion(url, "Petición de prueba A")
	if err != nil {
		t.Error(err) //si hay error, el test falla
	}
	// sabemos que es un flota64, un "number" en JSON (podriamos comporbarlo como parte del test)
	var contador1 = mapa["contador"].(float64)

	// lanzamos otra petión
	mapa, err = enviaPeticion(url, "Petición de prueba B")
	if err != nil {
		t.Error(err)
	}

	// Ahora comprobamos que el contador de peticiones se ha incrementado
	var contador2 = mapa["contador"].(float64)

	//Deberia ser estrictamente creciente
	if contador1 >= contador2 { /* (negar para ver "fail") */
		t.Errorf("contador no creciente: %.0f -> %.0f\n", contador1, contador2)
		// En caso de error mostramos porque falla el test
	}

	// Si invocamos go test -v podemos ver esta traza
	// con go test sólo mostrará los errores en caso de que los haya
	t.Logf("contador %.0f -> %.0f\n", contador1, contador2)
}
