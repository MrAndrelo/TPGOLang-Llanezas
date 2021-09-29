# TPGOLang-Llanezas

Alumno : José Andrés Llanezas
Trabajo Práctico del Seminario de Go Lang  2021



CONSIGNA
Crear una funcion que dada una cadena con un formato determinado genere una instancia de una estructura con los valores de los campos correspondientes al formato de la cadena.
La cadena de caracteres tiene el siguiente formato:

los primeros dos caracteres son el tipo
los segundos dos caracteres son el largo del valor
los siguientes N caracteres son el valor, donde N es el valor del bullet anterior.


******************************

Para resolver la consigna comencé con el modelo que genera una estructura a partir de un string "input" harcodeado desde el Main. Lo toma y subdivide en Type, Value y Length y chequea que cada parte sea válida y se corresponda con los requerimientos.

usé la libreria errors para manejar los errores y enviar mensajes al usuario acerca del error específico.
Libreria strconv para convertir string a diferentes tipos de dato. mas que nada Int para el es split del Input en segmentos según su posición.

El main solo pasa el input y devuelve el string parseado o el error sea cual sea el caso.
ambos se conectan pdesde main en el import.





