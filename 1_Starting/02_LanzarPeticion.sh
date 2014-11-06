echo "Hacemos una peticion al servidor sencillo con campos de la cabecera"
echo "arbitarios y un path /a/b/c/d. Enviamos un cuerpo con tres bytes. Estos"
echo "valores serán visibles en la traza del servidor"
set -x
curl -v localhost:7777/a/b/c -H "Content-Type: text/plain" -H "cabecera1: valor1" -H "cabecera1: valor2" -d "ABC"

