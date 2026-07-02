package service

import (
	"fmt"
	"grafos_util"
	Constantes "tp3/constantes"
	DB "tp3/database"
	"tp3/output"
)

func ReducirCaminosCmd(db *DB.Database, ruta string) {
	grafo := db.ObtenerGrafo()
	mst, pesoTotal := grafos_util.Prim(grafo)
	fmt.Printf(Constantes.PESO_TOTAL, int(pesoTotal))
	output.ExportarPajek(ruta, mst)
}
