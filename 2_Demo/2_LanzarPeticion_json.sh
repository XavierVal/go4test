echo "Enviamos al servidor la cadena que pasamos como primer argumento (\"$1\")"
echo "y recibimos un objeto JSON con el contador incrementado, la longitud de la"
echo "cadena que pasamos y las cabeceras que hemos enviado"
echo
(
set -x
curl -i localhost:7777/a/b/c  -H "Content-Type: text/plain" -d "$1"
)
echo
