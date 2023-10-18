package internal

type Musico struct {
	Nombre 				string
	Instrumento 		string
	DiasMaximosTocar 	int
	DiasDisponibles 	[]time.Time
}