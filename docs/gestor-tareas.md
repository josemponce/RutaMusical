
# Gestor de tareas

Necesitamos escoger una herramienta que nos permita automatizar ciertas tareas que sean muy repetitivas a lo largo del desarrollo de nuestro proyecto. Para empezar definiré los **criterios** por los cuales me decantaré por una elección u otra.

- **Lenguaje utilizado**: que se adapte bien y pueda resultar fácil de utilizar sobre el lenguaje escogido.
- **Más usados**: que sea usado por mucha gente, más seguidores en github nos dirá que se trata de uno de los recomendados.

Tras estar investigando un poco he dado con los que creo que son más usados, entre ellos:

- `Mage`: Bien integrado en el ecosistema de Go, aunque su sintaxis pueda resultar más complicada para la gente que no esté familiarizada con Go como es mi caso, a parte de que cuenta con bastante apoyo en github.

- `Task`: Tiene una sintaxis más simple y fácil de entender, y es conocido y utilizado en la comunidad de Go, se tiene que crear en un archivo YAML y además cuenta con un gran apoyo en github con casi nueve mil estrellas en github.

- `Make`: El más reconocido en el desarrollo de software, ampliamente compatible y con una estructura clara. El problema sería la complejidad de uso de este así como no ser específico para Go. 

En principio creo que usaré **Task** como gestor de tareas por que creo que será más fácil adaptarme y uso dentro de la comunidad aunque no descarto pasar a **Mage** cuando me encuentre más cómodo con Go. 