package internal

type Ruta struct {
	Conciertos 		[]ConciertoConMusico
	GananciaTotal 		float64
}

type ConciertoConMusico struct {
	Concierto		Concierto
	MusicosAsignados	[]Musico
}
