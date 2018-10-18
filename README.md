# documentos_crud
API de documentos, Integración con CI
documentos_crud master/develop
 ## Requirements
Go version >= 1.8.
 ## Preparation:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/documentos_crud
 ## Run
 Definir los valores de las siguientes variables de entorno:
  - `DOCUMENTOS_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `DOCUMENTOS_CRUD__PGUSER`: Usuario de la base de datos
 - `DOCUMENTOS_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `DOCUMENTOS_CRUD__PGURLS`: Host de conexión
 - `DOCUMENTOS_CRUD__PGDB`: Nombre de la base de datos
 - `DOCUMENTOS_CRUD__SCHEMA`: Esquema a utilizar en la base de datos
 ## Example:
DOCUMENTOS_HTTP_PORT=8095 DOCUMENTOS_CRUD__PGUSER=postgres DOCUMENTOS_CRUD__PGPASS=password DOCUMENTOS_CRUD__PGURLS=localhost DOCUMENTOS_CRUD__PGDB=local DOCUMENTOS_CRUD__SCHEMA=core_new bee run
 ## Model DB
![image](https://github.com/udistrital/documentos_crud/blob/develop/modelo_documentos_crud.png).
