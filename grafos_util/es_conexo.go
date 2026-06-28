package grafos_util

import (
	TDAGrafo "tdas/grafo"
)

func EsConexo[T comparable](grafo TDAGrafo.Grafo[T]) bool {
	iter := grafo.IterVertices()
	if !iter.HayAlgoMas() {
		return false // Inválido
	}
	origen := iter.VerActual()
	padres, _ := Bfs(grafo, origen)
	return padres.Cantidad() == grafo.CantidadVertices()
}
