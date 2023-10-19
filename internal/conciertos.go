package internal

import "time"

type Concierto struct {
	Fecha 			time.Time
	HoraInicio 		time.Time
	NombrePueblo 		string
	DineroPagado 		float64
	Duracion 		int // minutos
	Distancia		float64 // km
}
