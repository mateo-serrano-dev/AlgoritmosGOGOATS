package grafos_util

import (
	TDACola "tdas/cola"
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func OrdenTopologico[T comparable](grafo TDAGrafo.Grafo[T]) []T {
	// Precondicion: el grafo debe ser dirigido
	grados := TDADict.CrearHash[T, int]()
	for iter := grafo.IterVertices(); iter.HayAlgoMas(); iter.Avanzar() {
		v := iter.VerActual()
		if !grados.Pertenece(v) {
			grados.Guardar(v, 0)
		}

		for iterAdyacente := grafo.IterAdyacentes(v); iterAdyacente.HayAlgoMas(); iterAdyacente.Avanzar() {
			w := iterAdyacente.VerActual()
			if !grados.Pertenece(w) {
				grados.Guardar(w, 0)
			}

			grados.Guardar(w, grados.Obtener(w)+1)
		}
	}

	q := TDACola.CrearColaEnlazada[T]()
	for iter := grafo.IterVertices(); iter.HayAlgoMas(); iter.Avanzar() {
		if grados.Obtener(iter.VerActual()) == 0 {
			q.Encolar(iter.VerActual())
		}
	}

	res := make([]T, 0)
	for !q.EstaVacia() {
		v := q.Desencolar()
		res = append(res, v)
		for iter := grafo.IterAdyacentes(v); iter.HayAlgoMas(); iter.Avanzar() {
			w := iter.VerActual()
			grados.Guardar(w, grados.Obtener(w)-1)
			if grados.Obtener(w) == 0 {
				q.Encolar(w)
			}
		}
	}

	if len(res) == grafo.CantidadVertices() {
		return res
	}
	return nil
}
