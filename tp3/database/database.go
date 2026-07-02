package database

import (
	"fmt"
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	Models "tp3/models"
)

type Database struct {
	ciudades TDADict.Diccionario[string, Models.Ciudad]
	grafo    TDAGrafo.GrafoPesado[Models.Ciudad]
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

func (db *Database) CargarGrafo(g TDAGrafo.GrafoPesado[Models.Ciudad]) {
	db.grafo = g
}

func (db *Database) ObtenerGrafo() TDAGrafo.GrafoPesado[Models.Ciudad] {
	return db.grafo
}

func (db *Database) ImprimirCiudades() {
	fmt.Println("Ciudades: ")
	for it := db.ciudades.Iterador(); it.HayAlgoMas(); it.Avanzar() {
		_, city := it.VerActual()
		fmt.Println(city.Nombre())
	}
}
