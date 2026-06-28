package grafos_util

import (
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDALista "tdas/lista"
)

func hierholzer[T comparable](grafo TDAGrafo.Grafo[T], v T, lista TDALista.Lista[T], visitados TDADict.Diccionario[Arista[T], bool]) {
	for iter := grafo.IterAdyacentes(v); iter.HayAlgoMas(); iter.Avanzar() {
		w := iter.VerActual()
		aristaActual := CrearArista(v, w, 1)
		if !visitados.Pertenece(aristaActual) {
			visitados.Guardar(aristaActual, true)
			hierholzer(grafo, w, lista, visitados)
		}
	}
	lista.InsertarPrimero(v)
}

func Hierholzer[T comparable](grafo TDAGrafo.Grafo[T]) TDALista.Lista[T] {
	// Precondición: El grafo adminte ciclos Eulerianos y no esta vacío
	lista := TDALista.CrearListaEnlazada[T]()
	visitados := TDADict.CrearHash[Arista[T], bool]()
	iter := grafo.IterVertices()
	if !iter.HayAlgoMas() {
		return lista
	}
	lista.InsertarUltimo(iter.VerActual())
	hierholzer(grafo, iter.VerActual(), lista, visitados)
	return lista
}
