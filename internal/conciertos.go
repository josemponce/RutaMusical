package internal

import "time"

type Concierto struct {
	Fecha 			time.Time
	HoraInicio 		time.Time
	NombrePueblo 		string
	DineroPagado 		float64
	DuracionMinutos	uint 
	Costes			float64 
}
