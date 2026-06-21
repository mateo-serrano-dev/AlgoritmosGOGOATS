package database

import (
	TDADict "tdas/diccionario"
	Models "tp3/models"
)

type Database struct {
	ciudades TDADict.Diccionario[string, Models.Ciudad]
}

func (db *Database) RegistrarCiudad(nombre string, ciudad Models.Ciudad) {
	db.ciudades.Guardar(nombre, ciudad)
}

func (db *Database) ObtenerCiudad(s string) Models.Ciudad {
	return db.ciudades.Obtener(s)
}
