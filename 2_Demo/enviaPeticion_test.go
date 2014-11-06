package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Funcion auxiliar: Vamos a definir una funcion para que envíe datos a nuestro server y recoja la respuesta

// Las funciones pueden devover múltiples valores, ideal para indicar errores "fuera de banda"
func enviaPeticion(url string, body string) (mapa map[string]interface{}, err error) {
	// Obtenemos un Reader desde un string sencillamente con strings.NewReader
	// El interfaz Reader es implementado por ficheros, sockets, buffers, ...
	// Todos ellos podrían pasarse como body para el POST

	// Hacemos el envío de nuestro POST con los datos en el body
	resp, err := http.Post(url, "text/plain", strings.NewReader(body))

	// Como siempre, los errores no nos pillarán con la guardia baja,
	// los comprobamos según puedan pasar
	if err != nil {
		return nil, err // en caso de error devolvemos el error (no necesita ningún dato más)
	}
	defer resp.Body.Close()

	// Recogemos la respuesta
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err // igual que el anterior
	}

	// Y ahora obtenemos el "map" de la respuesta desde el JSON
	err = json.Unmarshal(respBody, &mapa)
	if err != nil {
		return nil, err
	}
	// si no a habido ningun error devolvemos el contenido de la respuesta obtenida
	return mapa, nil
}
