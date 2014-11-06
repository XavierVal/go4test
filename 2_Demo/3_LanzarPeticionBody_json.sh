today=$(date)

echo "Envia peticion - $today"
echo "Enviamos al servidor la cadena de texto fijo que contiene una parte variable (fecha)"
echo "y recibimos un objeto JSON con el contador incrementado, la longitud de la"
echo "cadena que pasamos y las cabeceras que hemos enviado"
echo
(
set -x
curl -i localhost:7777/path/de/miservicio -H "Content-Type: text/plain" -d "El cuerpo contiene datos como la fecha actual: $today"
)
echo
