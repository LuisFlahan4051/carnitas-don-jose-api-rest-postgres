# Instalación

Sigue los pasos a continuación para instalar y configurar el proyecto.

## Requisitos previos

Asegúrate de tener instalado lo siguiente antes de comenzar:

- [Node.js](https://nodejs.org) (versión X.X.X o superior)
- [Git](https://git-scm.com)

## Paso 1: Clonar el repositorio

```

git clone https://github.com/tu-usuario/tu-repositorio.git

```

## Paso 2: Instalar dependencias

```

cd tu-repositorio

```

## Paso 3: Configurar variables de entorno

Modifica el archivo `go.env` en la raíz del proyecto y configura las variables de entorno necesarias para tu equipo local.

## Paso 4: instala postgres usando docker con las siguientes variables de entorno:

```

POSTRES_USER=root
POSTGRES_PASSWORD=4051
POSTGRES_DB=carnitas_db

```

## Paso 5: IMPORTANTE, edita las variables en `init.sh` de acuerdo a lo que definiste en go.env, recuerda que tambien tienen que ser las mismas entradas del paso anterior:

```

docker cp -L .\database\schema\sql postgres_carnitas:.\scripts
docker exec -it postgres_carnitas /bin/bash

> chmod +x /scripts/init.sh
> chmod +x /scripts/restore.sh
> cd /scripts/
> ./restore.sh
> ./init.sh
> exit

```

## To restore the database, use the script restore.sh, and repeat step 5 once the following line has been executed

```
docker exec postgres_carnitas rm -rf ./scripts
```

## Paso 6: ejecuta la api solo escribiendo air en consola dentro del directorio raíz y luego abre el enlace definidor por la API

```

air

http://localhost:8080/users?admin_username=root&admin_password=root

```

¡Listo! Ahora deberías tener la aplicación funcionando en tu entorno local.

## Contribuir

Si deseas contribuir a este proyecto, sigue los pasos a continuación:

1. Crea una nueva rama: `git checkout -b mi-rama`
2. Realiza tus cambios y haz commit: `git commit -am 'Agrega nuevas funcionalidades'`
3. Sube tus cambios a la rama: `git push origin mi-rama`
4. Abre una Pull Request en GitHub

## Problemas y preguntas

Si tienes algún problema o pregunta, no dudes en abrir un issue en el repositorio.

Recuerda reemplazar `tu-usuario` y `tu-repositorio` con tu información correspondiente.

Espero que esto te sea útil. ¡Buena suerte con tu proyecto!
