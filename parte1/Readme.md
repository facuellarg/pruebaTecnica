# Prueba tenica IPCOM
## Parte 1

En este proyecto se pide ofrecer un endpoint que debe dar como resultado el acumulado de ventas de los dias indicados en la peticion.  
Para este endpoint se hizo uso de una de las grandes ventajas de Go, la paralelización, para asi lograr hacer las peticiones de manera asincrona y disminuir bastante el tiempo de respuesta.  
Tambien se uso el framework [echo] para crear las peticiones.

para correr este proyecto debemos situarnos en la carpera _parte1_ y ejecutar el comando   
```sh
$ go run *.go

```
El cual ejecuta el proyecto y descarga las dependencias necesarias.  
Ahora que el proyecto esta corriendo se nos habilida la url con la siguiente estructura:  

http://localhost:8000/ventas/:fecha?dias=int  

a la cual podemos hacerle peticiones de tipo GET.  
Podemos hacer una peticion  corriendo la siguiente instruccion en otro terminal:
```sh
$ curl -X GET http://localhost:8000/ventas/2019-12-12?dias=1
```
Podemos no incluir la variable _dias_ en la query y se tomara solo la fecha especificada en el parametro de _fecha_. Cabe aclarar que la fecha sigue el formato __YYYY_MM_DD__


[echo]:https://echo.labstack.com/guide/request