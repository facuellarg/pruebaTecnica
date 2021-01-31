# Prueba tenica IPCOM
## Parte 2

En este proyecto se pide leer un archivo csv con el siguiente formato:
```text
organizacion, usuario, rol 
org1,jperez,admin 
org1,jperez,superadmin
org1,asosa,writer
org2,jperez,admin 
org2,rrodriguez,writer
org2,rrodriguez,editor
```
 y dar como resultado una estructura de tipo json en la cual esta informacion este mas organizada.
 ```
 [
   { organization: “org1”, 
     users: [{
       username: “jperez”,
       roles: [“admin”,”superadmin”] }
       {
       username: “asosa”,
       roles: [“writer”] }
    ]
   }, 
  {
      organization:”org2”, 
    … 
}]
 ```
Lo primero que debemos hacer es situarnos en la carpera _parte2_,  luego tendremos varias opciones para correr el programa, las cuales podemos ver usando el comando: 
```sh
$ go run *.go -h
```
el cual nos da como resultado:
```text
  -in string
        path of input file (default "example.in")
  -out string
        set value if you want the output in a file
```
El parametro __in__ nos permite especificar la ruta del archivo csv de entrada, el cual por defecto es _example.in_. El parametro __out__ nos permite especificar el arhivo de salida, por defecto se usa la consola.
por ejemplo si queremos que el programa lea el archivo _myInput.csv_ y nos de el resultado en el archivo _myOutput.json_ correremos el siguiente comando

```
$ go run *.go -in=myInput.csv -out=myOutput.json
```
Esto nos crear el archivo _myOutput.json_ con la salida correspondiente al archivo de entrada.

El proyecto contiene el archivo de ejemplo _example.in_, para correr el programa y ver el resultado en consola usamos el comando 
```
$ go run *.go
```