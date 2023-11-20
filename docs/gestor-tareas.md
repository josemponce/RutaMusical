
# Gestor de tareas

Necesitamos escoger una herramienta que nos permita automatizar ciertas tareas que sean muy repetitivas a lo largo del desarrollo de nuestro proyecto. Para empezar definiré los **criterios** por los cuales me decantaré por una elección u otra.

- **Lenguaje utilizado**: que se adapte bien y pueda resultar fácil de utilzar sobre el lenguaje escogido.
- **Popularidad**: que sea usado por mucha gente nos dirá que se trata de uno de los recomendables.
- **Facilidad de uso**: que no sea muy complicado entender el funcionamiento del mismo.

Tras estar investigando un poco he dado con los que creo que son más usados, entre ellos:

- `Mage`: posiblemente el más utilizado pricipalmente por estar bien integrado en el ecosistmema de Go, aunque su sistaxis pueda resultar más complicada para la gente que no esté familiarizada con Go como es mi caso.

- `Task`: Tiene una sintaxis mas simple y fácil de entender, y es conocdio y utilizado en la comunidad de Go, se tiene que crear en un archivo YAML.

- `Make`: El mas reconocido en el desarrollo de software, ampliamente compatible y con una estructura clara. El problema sería la complejidad de uso de este así como no ser específico para Go. 

En principio creo que usaré **Task** como gestor de tareas por su simplicidad y uso dentro de la comunidad aunque no descarto pasar a **Mage** cuando me encuentre más cómodo con Go. 