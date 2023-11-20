package internal

import "time"

type Musico struct {
	Nombre 		string
	DiasMaximosTocar 	int
	DiasDisponibles 	[]time.Time
}
