package database

import (
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	Models "tp3/models"
)

type Database struct {
	ciudades TDADict.Diccionario[string, TDAGrafo.Nodo[Models.Ciudad]]
}

func (db *Database) RegistrarCiudad(nombre string, nodo TDAGrafo.Nodo[Models.Ciudad]) {
	db.ciudades.Guardar(nombre, nodo)
}

func (db *Database) ObtenerCiudad(s string) TDAGrafo.Nodo[Models.Ciudad] {
	return db.ciudades.Obtener(s)
}
