package database

import (
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	Models "tp3/models"
)

type Database struct {
	ciudades TDADict.Diccionario[string, Models.Ciudad]
	grafo    TDAGrafo.Grafo[Models.Ciudad]
}

func CrearDatabase() *Database {
	ciudades := TDADict.CrearHash[string, Models.Ciudad]()
	db := new(Database)
	db.ciudades = ciudades
	return db
}

func (db *Database) RegistrarCiudad(nombre string, ciudad Models.Ciudad) {
	db.ciudades.Guardar(nombre, ciudad)
}

func (db *Database) ObtenerCiudad(s string) Models.Ciudad {
	return db.ciudades.Obtener(s)
}

func (db *Database) CargarGrafo(g TDAGrafo.Grafo[Models.Ciudad]) {
	db.grafo = g
}
