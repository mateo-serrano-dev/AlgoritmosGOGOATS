package service

import (
	"bufio"
	"fmt"
	"grafos_util"
	"os"
	"strings"
	TDAGrafo "tdas/grafo"
	Constantes "tp3/constantes"
	DB "tp3/database"
)

// Queda en este archivo porque este parsing es especifico de este comando
func parsearRecomendaciones(db *DB.Database, ruta string) TDAGrafo.GrafoNoPesado[string] {
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf(Constantes.ERR_ABRIR_ARCHIVO)
		return nil
	}
	defer archivo.Close()

	grafo := TDAGrafo.CrearGrafoNoPesado[string](true)
	// No nos interesa si es pesado o no
	// Usamos strings por comodidad

	for iter := db.ObtenerGrafo().IterVertices(); iter.HayAlgoMas(); iter.Avanzar() {
		v := iter.VerActual()
		grafo.AgregarVertice(v.Nombre())
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		arista := strings.Split(linea, ",")
		v := strings.TrimSpace(arista[0])
		w := strings.TrimSpace(arista[1])
		grafo.AgregarArista(v, w)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf(Constantes.ARCHIVO_CSV_ERR)
		return nil
	}
	return grafo
}

func ItinerarioCmd(db *DB.Database, ruta string) {
	grafo := parsearRecomendaciones(db, ruta)
	if grafo == nil {
		return
	}
	camino := grafos_util.OrdenTopologico(grafo)
	if camino == nil {
		fmt.Println(Constantes.ERR_RECORRIDO_NO_ENCONTRADO)
		return
	}
	salida := strings.Join(camino, " -> ")
	fmt.Println(salida)
}
