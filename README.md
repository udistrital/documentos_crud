# documentos_crud
API de gestión de documentos

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `documentos_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/planesticud/documentos_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `DOCUMENTOS_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `DOCUMENTOS_CRUD__PGUSER`: Usuario de la base de datos
 - `DOCUMENTOS_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `DOCUMENTOS_CRUD__PGURLS`: Host de conexión
 - `DOCUMENTOS_CRUD__PGDB`: Nombre de la base de datos
 - `DOCUMENTOS_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
DOCUMENTOS_HTTP_PORT=8094 DOCUMENTOS_CRUD__PGUSER=user DOCUMENTOS_CRUD__PGPASS=password DOCUMENTOS_CRUD__PGURLS=localhost DOCUMENTOS_CRUD__PGDB=bd DOCUMENTOS_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![image](https://github.com/planesticud/documentos_crud/blob/develop/modelo_documentos_crud.png).
