package service

import (
	"fmt"
	"grafos_util"
	Constantes "tp3/constantes"
	DB "tp3/database"
	Modelos "tp3/models"
	"tp3/output"
)

func ViajeCmd(db *DB.Database, origen Modelos.Ciudad, ruta string) {
	grafo := db.ObtenerGrafo()
	if !grafos_util.EsConexo(grafo) || !grafos_util.TieneVerticesGradoPar(grafo) {
		fmt.Println(Constantes.ERR_RECORRIDO_NO_ENCONTRADO)
		return
	}
	pila := grafos_util.Hierholzer(grafo, origen)
	camino, tiempo := grafos_util.ReconstruirCaminoHierholzer(grafo, pila)
	err := output.ExportarKml(camino, ruta)
	if err != nil {
		fmt.Print(Constantes.ERR_EXPORTAR)
		return
	}
	fmt.Println(output.ConstruirMensajeSalida(camino))
	fmt.Printf(Constantes.TIEMPO_TOTAL, int(tiempo))
}
