# documentos_crud
API de gestión de documentos

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
DOCUMENTOS_HTTP_PORT=[Puerto asignado para la ejecución del API]
DOCUMENTOS_CRUD__PGUSER=[Usuario de la base de datos]
DOCUMENTOS_CRUD__PGPASS=[Clave del usuario para la conexión a la base de datos]
DOCUMENTOS_CRUD__PGURLS=[Host de conexión]
DOCUMENTOS_CRUD__PGDB=[Nombre de la base de datos]
DOCUMENTOS_CRUD__SCHEMA=[Esquema a utilizar en la base de datos]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con DOCUMENTOS_CRUD__...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/novedades_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/novedades_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
DOCUMENTOS_HTTP_PORT=8094 DOCUMENTOS_CRUD__PGUSER=user DOCUMENTOS_CRUD__PGPASS=password DOCUMENTOS_CRUD__PGURLS=localhost DOCUMENTOS_CRUD__PGDB=bd DOCUMENTOS_CRUD__SCHEMA=schema_new bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=documentos_crud . --no-cache
# docker run -p 80:80 documentos_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/documentos_crud

#2. Moverse a la carpeta del repositorio
cd documentos_crud

#3. Crear un fichero con el nombre **custom.env**
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/documentos_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/documentos_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/documentos_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/documentos_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/documentos_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/documentos_crud) |


## Modelo de Datos
[Modelo de Datos documentos_crud](https://github.com/planesticud/documentos_crud/blob/develop/modelo_documentos_crud.png)

## Licencia

This file is part of documentos_crud.

documentos_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

documentos_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with documentos_crud. If not, see https://www.gnu.org/licenses/.
