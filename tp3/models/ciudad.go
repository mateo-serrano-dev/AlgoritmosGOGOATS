package models

type Ciudad struct {
	nombre   string
	latitud  float64
	longitud float64
}

func CrearCiudad(nombre string, lat, long float64) Ciudad {
	return Ciudad{nombre, lat, long}
}
