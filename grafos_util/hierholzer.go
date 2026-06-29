package grafos_util

import (
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDAPila "tdas/pila"
)

func hierholzer[T comparable](grafo TDAGrafo.Grafo[T], v T, pila TDAPila.Pila[T], visitados TDADict.Diccionario[Arista[T], bool]) {
	for iter := grafo.IterAdyacentes(v); iter.HayAlgoMas(); iter.Avanzar() {
		w := iter.VerActual()
		aristaActual := CrearArista(v, w, 1)
		if !visitados.Pertenece(aristaActual) {
			visitados.Guardar(aristaActual, true)
			hierholzer(grafo, w, pila, visitados)
		}
	}
	// Previo a apilar visitamos cada vértice adyacente de v
	// Visitando así cada arista del grafo para cada para vértices
	pila.Apilar(v)
}

func Hierholzer[T comparable](grafo TDAGrafo.Grafo[T]) TDAPila.Pila[T] {
	// Precondición: El grafo admite ciclos Eulerianos y no esta vacío
	pila := TDAPila.CrearPilaDinamica[T]()
	visitados := TDADict.CrearHash[Arista[T], bool]()
	iter := grafo.IterVertices()
	if !iter.HayAlgoMas() {
		return pila
	}
	pila.Apilar(iter.VerActual())
	hierholzer(grafo, iter.VerActual(), pila, visitados)
	return pila
}
