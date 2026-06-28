package grafos_util

import (
	TDACola "tdas/cola"
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func Bfs[T comparable](grafo TDAGrafo.Grafo[T], origen T) (TDADict.Diccionario[T, T], TDADict.Diccionario[T, int]) {
	visitados := TDADict.CrearHash[T, bool]()
	padres := TDADict.CrearHash[T, T]()
	orden := TDADict.CrearHash[T, int]()
	q := TDACola.CrearColaEnlazada[T]()
	visitados.Guardar(origen, true)
	padres.Guardar(origen, origen)
	orden.Guardar(origen, 0)
	q.Encolar(origen)
	for !q.EstaVacia() {
		v := q.Desencolar()
		for iterAdyacentes := grafo.IterAdyacentes(v); iterAdyacentes.HayAlgoMas(); iterAdyacentes.Avanzar() {
			w := iterAdyacentes.VerActual()
			if !visitados.Pertenece(w) {
				visitados.Guardar(w, true)
				padres.Guardar(w, v)
				orden.Guardar(w, orden.Obtener(v)+1)
				q.Encolar(w)
			}
		}
	}
	return padres, orden
}
