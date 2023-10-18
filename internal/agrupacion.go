package internal

type AgrupacionMusical struct {
	NombreAgrupacion 	string
	Musicos 			[]Musico
	
	// Coordenadas pueblo de la agrupacion
	Latitud 			float64
	Longitud 			float64
}