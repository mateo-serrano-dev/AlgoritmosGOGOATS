package grafos_util

import (
	"math"
	"slices"
	TDAHeap "tdas/cola_prioridad"
	TDADict "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

type verticeConDistancia[T comparable] struct {
	vertice    T
	distActual float64
}

func cmpDijkstra[T comparable](a, b verticeConDistancia[T]) int {
	if a.distActual < b.distActual {
		return -1
	} else if a.distActual > b.distActual {
		return 1
	}
	return 0
}

func Dijkstra[T comparable](grafo TDAGrafo.GrafoPesado[T], origen T) (TDADict.Diccionario[T, T], TDADict.Diccionario[T, float64]) {
	// Precondicion: origen debe pertencer al grafo y los pesos deben ser positivos
	dist := TDADict.CrearHash[T, float64]()
	padres := TDADict.CrearHash[T, T]()

	for iter := grafo.IterVertices(); iter.HayAlgoMas(); iter.Avanzar() {
		v := iter.VerActual()
		dist.Guardar(v, math.Inf(1))
	}

	dist.Guardar(origen, 0)
	padres.Guardar(origen, origen)
	q := TDAHeap.CrearHeap[verticeConDistancia[T]](cmpDijkstra)
	q.Encolar(verticeConDistancia[T]{
		vertice:    origen,
		distActual: 0,
	})

	for !q.EstaVacia() {
		v := q.Desencolar()
		if v.distActual > dist.Obtener(v.vertice) {
			continue
		}
		for iter := grafo.IterAdyacentes(v.vertice); iter.HayAlgoMas(); iter.Avanzar() {
			w := iter.VerActual()
			if dist.Obtener(v.vertice)+grafo.Peso(v.vertice, w) < dist.Obtener(w) {
				dist.Guardar(w, dist.Obtener(v.vertice)+grafo.Peso(v.vertice, w))
				padres.Guardar(w, v.vertice)
				q.Encolar(verticeConDistancia[T]{
					vertice:    w,
					distActual: dist.Obtener(w),
				})
			}
		}
	}
	return padres, dist
}

func ReconstruirCaminoDijkstra[T comparable](padres TDADict.Diccionario[T, T], origen, destino T) []T {
	res := make([]T, 0)
	// Reconstruimos el camino desde el destino al origen, siendo necesario invertir el arreglo final
	actual := destino
	for actual != origen {
		res = append(res, actual)
		actual = padres.Obtener(actual)
	}
	res = append(res, origen)
	slices.Reverse(res)
	return res
}
