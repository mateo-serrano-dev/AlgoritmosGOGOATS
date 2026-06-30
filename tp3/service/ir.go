package service

import (
	"fmt"
	"grafos_util"
	"math"
	"strings"
	Constantes "tp3/constantes"
	DB "tp3/database"
	Models "tp3/models"
	"tp3/output"
)

func construirMensajeSalida(camino []Models.Ciudad) string {
	ciudades := make([]string, 0)
	for _, ciudad := range camino {
		ciudades = append(ciudades, ciudad.Nombre())
	}
	return strings.Join(ciudades, " -> ")
}

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

	camino := grafos_util.ReconstruirCamino(padres, desde, hasta)
	err := output.ExportarKml(camino, ruta)
	if err != nil {
		fmt.Print(Constantes.ERR_EXPORTAR)
		return
	}
	fmt.Println(construirMensajeSalida(camino))
	fmt.Printf(Constantes.TIEMPO_TOTAL, int(dist.Obtener(hasta)))
}
