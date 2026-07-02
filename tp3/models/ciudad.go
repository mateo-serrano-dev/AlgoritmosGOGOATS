package models

type Ciudad struct {
	nombre   string
	latitud  float64
	longitud float64
}

func CrearCiudad(nombre string, lat, long float64) Ciudad {
	return Ciudad{nombre, lat, long}
}

func (c Ciudad) Nombre() string {
	return c.nombre
}

func (c Ciudad) Latitud() float64 {
	return c.latitud
}

func (c Ciudad) Longitud() float64 {
	return c.longitud
}
