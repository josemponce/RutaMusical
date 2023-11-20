
# Gestor de dependencias

Existen herramientas como son `dep` o `glide` que han sido las más utilizadas durante mucho tiempo, sin embargo, resultan ser poco ágiles y el hecho de que no formen parte del proyecto oficial de Go crea muchas complicaciones. Desde 2018, a partir de la versión de Go 1.11 el equipo de Go creó `Go Module` introduciendo el concepto de "modulo" (colecciones de paquetes relacionados) que lo convierte en una herramienta más versátil.

Esta herramienta nos proporciona un enfoque eficaz para gestionar las depencias, permitiéndonos añadir, actualizar o eliminar versiones específicas de dependencias ayudandonos a mantener el control sobre nuestro proyecto.

Para usar esta herrmienta algunos de los comandos a utilziar serán por ejemplo:

- go mod init: incializan un nuevo proyecto generando el archivo go.mod, que contiene la información del módulo, sus dependencias y restriccioens de versión.

- go get: para añadir o actualizar una dependencia.

- go mod tidy: para eliminar las dependencias no utilziadas del archivo go.mod.