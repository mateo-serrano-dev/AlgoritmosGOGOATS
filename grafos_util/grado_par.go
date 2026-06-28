package grafos_util

import (
	TDAGrafo "tdas/grafo"
)

func TieneVerticesGradoPar[T comparable](grafo TDAGrafo.Grafo[T]) bool {
	// Devuelve true si todos los vértices son de grado par
	for iter := grafo.IterVertices(); iter.HayAlgoMas(); iter.Avanzar() {
		contador := 0
		for iterAdyacentes := grafo.IterAdyacentes(iter.VerActual()); iterAdyacentes.HayAlgoMas(); iterAdyacentes.Avanzar() {
			contador++
		}
		if contador%2 != 0 {
			return false
		}
	}
	return true
}
