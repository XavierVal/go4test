echo "######################################################################################"
echo " 6) Lanzar los tests para la validar el rendimiento de nuestras funciones de test"
echo "    Primero lanzará todos los nuestros tests y luego medirá los tiempos y el rendimiento"
echo ""
echo "                                                www.testeandoSoftware.com"
echo "######################################################################################"
echo "Lanza los benchmark en modo verbose para nuestros test (contador y longitud). "
echo "Ademas ejecuta los test primero"
echo "Generará un montón de trazas en el servidor"
echo "######################################################################################"
echo ""
set -x
go test -test.bench=. -v

