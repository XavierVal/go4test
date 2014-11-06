echo "######################################################################################"
echo " 5) Lanzar los tests para la validar la funcion que mide la longitud de las peticiones"
echo ""
echo "                                                 www.testeandoSoftware.com"
echo "######################################################################################"
echo "Lanzamos los test para que se ejecute TestLongitud en modo verbose"
echo "Nos indica que test case se está ejecutando y muestra las trazas"
echo "que hayamos puesto con t.Log o t.Logf. Como se apoya en 'testing/quick''"
echo "generará varias trazas en el servidor con diferentes cadenas para el body"
echo "de la peticion"
echo "######################################################################################"
echo ""
set -x
go test -run Longitud -v

