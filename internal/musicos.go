package internal

import "time"

type Musico struct {
	Nombre 		string
	Instrumento 		string
	DiasMaximosTocar 	int
	DiasDisponibles 	[]time.Time
}
