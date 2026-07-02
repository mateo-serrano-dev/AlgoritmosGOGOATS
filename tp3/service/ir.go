package service

import (
	"fmt"
	"grafos_util"
	"math"
	Constantes "tp3/constantes"
	DB "tp3/database"
	Models "tp3/models"
	"tp3/output"
)

func IrCmd(db *DB.Database, desde, hasta Models.Ciudad, ruta string) {
	grafo := db.ObtenerGrafo()
	if !grafo.PerteneceVertice(desde) || !grafo.PerteneceVertice(hasta) {
		fmt.Println(Constantes.ERR_RECORRIDO_NO_ENCONTRADO)
		return
	}
	padres, dist := grafos_util.Dijkstra(grafo, desde)
	if dist.Obtener(hasta) == math.Inf(1) {
		fmt.Println(Constantes.ERR_RECORRIDO_NO_ENCONTRADO)
		return
	}

	camino := grafos_util.ReconstruirCaminoDijkstra(padres, desde, hasta)
	err := output.ExportarKml(camino, ruta)
	if err != nil {
		fmt.Print(Constantes.ERR_EXPORTAR)
		return
	}
	fmt.Println(output.ConstruirMensajeSalida(camino))
	fmt.Printf(Constantes.TIEMPO_TOTAL, int(dist.Obtener(hasta)))
}
