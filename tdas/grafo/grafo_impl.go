package grafo

import TDADict "tdas/diccionario"

type grafoImp[T any] struct {
	nodos    TDADict.Diccionario[Nodo[T], TDADict.Diccionario[Nodo[T], int]]
	dirigido bool
}

func CrearGrafo[T any](dirigido, pesado bool) Grafo[T] {
	nodos := TDADict.CrearHash[Nodo[T], TDADict.Diccionario[Nodo[T], int]]()
	g := grafoImp[T]{nodos, dirigido}
	if pesado {
		return &grafoPesadoImp[T]{g}
	} else {
		return &grafoNoPesadoImp[T]{g}
	}
}

func (g *grafoImp[T]) AgregarVertice(dato T) {
	nodo := nodoImp[T]{dato}
	vecinos := TDADict.CrearHash[Nodo[T], int]()
	g.nodos.Guardar(nodo, vecinos)
}

func (g *grafoImp[T]) BorrarVertice(nodo Nodo[T]) {
	g.nodos.Borrar(nodo)
	for iter := g.nodos.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		_, vecinos := iter.VerActual()
		vecinos.Borrar(nodo)
	}
}

func (g *grafoImp[T]) BorrarArista(desde, hasta Nodo[T]) {
	g.nodos.Obtener(desde).Borrar(hasta)
	if !g.dirigido {
		g.nodos.Obtener(hasta).Borrar(desde)
	}
}

// --- Pesado o No Pesado ---
type grafoPesadoImp[T any] struct {
	grafoImp[T]
}

type grafoNoPesadoImp[T any] struct {
	grafoImp[T]
}

func (g *grafoImp[T]) agregarArista(desde, hasta Nodo[T], peso int) {
	g.nodos.Obtener(desde).Guardar(hasta, peso)
	if !g.dirigido {
		g.nodos.Obtener(desde).Guardar(hasta, peso)
	}
}

func (g *grafoPesadoImp[T]) AgregarArista(desde, hasta Nodo[T], peso int) {
	g.agregarArista(desde, hasta, peso)
}

func (g *grafoPesadoImp[T]) Peso(desde, hasta Nodo[T]) int {
	return g.nodos.Obtener(desde).Obtener(hasta)
}

func (g *grafoNoPesadoImp[T]) AgregarArista(desde, hasta Nodo[T]) {
	g.agregarArista(desde, hasta, 1)
}

// --- Nodo ---
type nodoImp[T any] struct {
	dato T
}

func (n nodoImp[T]) Dato() T {
	return n.dato
}

// --- Iterador Vertices ---
type iteradorVertices[T any] struct {
	iter TDADict.IterDiccionario[Nodo[T], TDADict.Diccionario[Nodo[T], int]]
}

func (g *grafoImp[T]) IterVertices() IteradorVertices[T] {
	i := new(iteradorVertices[T])
	i.iter = g.nodos.Iterador()
	return i
}

func (i *iteradorVertices[T]) HayAlgoMas() bool {
	return i.iter.HayAlgoMas()
}

func (i *iteradorVertices[T]) Avanzar() {
	i.iter.Avanzar()
}

func (i *iteradorVertices[T]) VerActual() Nodo[T] {
	nodo, _ := i.iter.VerActual()
	return nodo
}

// --- Iterador adyacentes ---
type iteradorAdyacentes[T any] struct {
	iter TDADict.IterDiccionario[Nodo[T], int]
}

func (g *grafoImp[T]) IterAdyacentes(nodo Nodo[T]) IteradorVertices[T] {
	i := new(iteradorAdyacentes[T])
	i.iter = g.nodos.Obtener(nodo).Iterador()
	return i
}

func (i *iteradorAdyacentes[T]) HayAlgoMas() bool {
	return i.iter.HayAlgoMas()
}

func (i *iteradorAdyacentes[T]) Avanzar() {
	i.iter.Avanzar()
}

func (i *iteradorAdyacentes[T]) VerActual() Nodo[T] {
	nodo, _ := i.iter.VerActual()
	return nodo
}
