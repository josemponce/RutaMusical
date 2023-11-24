
# Gestor de dependencias

Existen herramientas como son `dep` o `glide` que han sido las más utilizadas durante mucho tiempo, sin embargo, resultan ser poco ágiles y el hecho de que no formen parte del proyecto oficial de Go crea muchas complicaciones. Desde 2018, a partir de la versión de Go 1.11 el equipo de Go creó `Go Module` introduciendo el concepto de "módulo" (colecciones de paquetes relacionados) que lo convierte en una herramienta más versátil.

Esta herramienta nos proporciona un enfoque eficaz para gestionar las dependencias, permitiéndonos añadir, actualizar o eliminar versiones específicas de dependencias ayudándonos a mantener el control sobre nuestro proyecto.

Para usar esta herramienta algunos de los comandos a utilizar serán por ejemplo:

- go mod init: inicializan un nuevo proyecto generando el archivo go.mod, que contiene la información del módulo, sus dependencias y restricciones de versión.

- go get: para añadir o actualizar una dependencia.

- go mod tidy: para eliminar las dependencias no utilizadas del archivo go.mod.