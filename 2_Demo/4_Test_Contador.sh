echo "######################################################################################"
echo " 4) Lanzar los tests para la validar el contador de peticiones"
echo ""
echo "                                                www.testeandoSoftware.com"
echo "######################################################################################"
echo "Vamos a lanzar un test que comprueba que el Contador de nuestro "
echo "server funciona correctamente incrementandose con cada petición"
echo "Nos indica que test case se está ejecutando y muestra las trazas"
echo "que hayamos puesto con t.Log o t.Logf"
echo "######################################################################################"
echo ""
set -x
go test -run Contador -v
