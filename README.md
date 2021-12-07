# ruler_api

Api para la gestion y consumo de reglas

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

- [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
- [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno

```shell
RULER_APP_PORT
RULER_DB_USER
RULER_DB_URL
RULER_DB_PASS
RULER_DB_NAME
RULER_DB_SCHEMAS
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con `RULER_...`

### Ejecución del Proyecto

```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/ruler_api

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/ruler_api

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
RULER_APP_PORT=8080 RULER_DB_URL=127.0.0.1:27017 RULER_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile

```shell
# docker build --tag=ruler_api . --no-cache
# docker run -p 80:80 ruler_api
```

### Ejecución docker-compose

```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/ruler_api

#2. Moverse a la carpeta del repositorio
cd ruler_api

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
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
# Not Data
```

## Estado CI

| Develop | Release 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/ruler_api/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/ruler_api) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/ruler_api/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/ruler_api) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/ruler_api/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/ruler_api) |

## Licencia

This file is part of ruler_api.

ruler_api is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ruler_api is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with ruler_api. If not, see https://www.gnu.org/licenses/.
