package grafos_util

import (
	TDAHeap "tdas/cola_prioridad"
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func cmpPrim[T comparable](a, b Arista[T]) int {
	if a.peso < b.peso {
		return 1
	} else if a.peso > b.peso {
		return -1
	}
	return 0
}

func Prim[T comparable](grafo TDAGrafo.GrafoPesado[T]) (TDAGrafo.GrafoPesado[T], float64) {
	// Precondicion: El grafo recibido debe ser no dirigido, conexo y pesado

	// Buscamos un vértice aleatorio
	iter := grafo.IterVertices()
	if !iter.HayAlgoMas() {
		return nil, 0
	}
	v := iter.VerActual()

	visitados := TDADict.CrearHash[T, bool]() // Usamos un hashmap como hashset
	visitados.Guardar(v, true)

	q := TDAHeap.CrearHeap[Arista[T]](cmpPrim)
	for iterAdyacentes := grafo.IterAdyacentes(v); iterAdyacentes.HayAlgoMas(); iterAdyacentes.Avanzar() {
		w := iterAdyacentes.VerActual()
		q.Encolar(CrearArista(v, w, grafo.Peso(v, w)))
	}

	arbol := TDAGrafo.CrearGrafoPesado[T](false)
	arbol.AgregarVertice(v)

	var pesoTotal float64

	for !q.EstaVacia() {
		x := q.Desencolar()
		if visitados.Pertenece(x.w) {
			continue
		}
		arbol.AgregarVertice(x.w)
		arbol.AgregarArista(x.v, x.w, x.peso)
		pesoTotal += x.peso

		visitados.Guardar(x.w, true)
		for iterAdyacentes := grafo.IterAdyacentes(x.w); iterAdyacentes.HayAlgoMas(); iterAdyacentes.Avanzar() {
			if !visitados.Pertenece(iterAdyacentes.VerActual()) {
				q.Encolar(CrearArista(x.w, iterAdyacentes.VerActual(), grafo.Peso(x.w, iterAdyacentes.VerActual())))
			}
		}
	}
	return arbol, pesoTotal
}
