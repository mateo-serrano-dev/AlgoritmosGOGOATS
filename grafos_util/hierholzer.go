package grafos_util

import (
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDAPila "tdas/pila"
)

func hierholzer[T comparable](grafo TDAGrafo.GrafoPesado[T], v T, pila TDAPila.Pila[T], visitados TDADict.Diccionario[Arista[T], bool]) {
	for iter := grafo.IterAdyacentes(v); iter.HayAlgoMas(); iter.Avanzar() {
		w := iter.VerActual()
		peso := grafo.Peso(v, w)
		aristaActual := CrearArista(v, w, peso)
		aristaInversa := CrearArista(w, v, peso) // Necesaria para grafos no dirigidos
		if !visitados.Pertenece(aristaActual) {
			visitados.Guardar(aristaActual, true)
			visitados.Guardar(aristaInversa, true)
			hierholzer(grafo, w, pila, visitados)
		}
	}
	// Previo a apilar visitamos cada vértice adyacente de v
	// Visitando así cada arista del grafo para cada par vértices
	pila.Apilar(v)
}

func Hierholzer[T comparable](grafo TDAGrafo.GrafoPesado[T], origen T) TDAPila.Pila[T] {
	// Precondición: El grafo admite ciclos Eulerianos y no esta vacío
	// Funciona para grafos pesados o pesados, si es necesario se pasa 1 en el peso de todas las aristas
	pila := TDAPila.CrearPilaDinamica[T]()
	visitados := TDADict.CrearHash[Arista[T], bool]()
	hierholzer(grafo, origen, pila, visitados)
	return pila
}

func ReconstruirCaminoHierholzer[T comparable](grafo TDAGrafo.GrafoPesado[T], pila TDAPila.Pila[T]) ([]T, float64) {
	res := make([]T, 0)
	var pesoTotal float64
	if pila.EstaVacia() {
		return res, 0
	}

	v := pila.Desapilar()
	res = append(res, v)

	for !pila.EstaVacia() {
		w := pila.Desapilar()
		res = append(res, w)
		pesoTotal += grafo.Peso(v, w)
		v = w
	}
	return res, pesoTotal
}
