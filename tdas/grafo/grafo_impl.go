package grafo

import TDADict "tdas/diccionario"

type grafoImp[T comparable] struct {
	nodos    TDADict.Diccionario[T, TDADict.Diccionario[T, float64]]
	dirigido bool
}

func CrearGrafoPesado[T comparable](dirigido bool) GrafoPesado[T] {
	nodos := TDADict.CrearHash[T, TDADict.Diccionario[T, float64]]()
	g := grafoImp[T]{nodos, dirigido}
	return &grafoPesadoImp[T]{g}
}

func CrearGrafoNoPesado[T comparable](dirigido bool) GrafoNoPesado[T] {
	nodos := TDADict.CrearHash[T, TDADict.Diccionario[T, float64]]()
	g := grafoImp[T]{nodos, dirigido}
	return &grafoNoPesadoImp[T]{g}
}

func (g *grafoImp[T]) AgregarVertice(dato T) {
	vecinos := TDADict.CrearHash[T, float64]()
	g.nodos.Guardar(dato, vecinos)
}

func (g *grafoImp[T]) BorrarVertice(dato T) {
	g.nodos.Borrar(dato)
	for iter := g.nodos.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		_, vecinos := iter.VerActual()
		vecinos.Borrar(dato)
	}
}

func (g *grafoImp[T]) BorrarArista(desde, hasta T) {
	g.nodos.Obtener(desde).Borrar(hasta)
	if !g.dirigido {
		g.nodos.Obtener(hasta).Borrar(desde)
	}
}

func (g *grafoImp[T]) CantidadVertices() int {
	return g.nodos.Cantidad()
}

func (g *grafoImp[T]) CantidadAristas(dato T) int {
	return g.nodos.Obtener(dato).Cantidad()
}

func (g *grafoImp[T]) PerteneceVertice(dato T) bool {
	return g.nodos.Pertenece(dato)
}

// --- Pesado o No Pesado ---
type grafoPesadoImp[T comparable] struct {
	grafoImp[T]
}

type grafoNoPesadoImp[T comparable] struct {
	grafoImp[T]
}

func (g *grafoImp[T]) agregarArista(desde, hasta T, peso float64) {
	g.nodos.Obtener(desde).Guardar(hasta, peso)
	if !g.dirigido {
		g.nodos.Obtener(desde).Guardar(hasta, peso)
	}
}

func (g *grafoPesadoImp[T]) AgregarArista(desde, hasta T, peso float64) {
	g.agregarArista(desde, hasta, peso)
}

func (g *grafoPesadoImp[T]) Peso(desde, hasta T) float64 {
	return g.nodos.Obtener(desde).Obtener(hasta)
}

func (g *grafoNoPesadoImp[T]) AgregarArista(desde, hasta T) {
	g.agregarArista(desde, hasta, 1)
}

// --- Iterador Vertices ---
type iteradorVertices[T comparable] struct {
	iter TDADict.IterDiccionario[T, TDADict.Diccionario[T, float64]]
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

func (i *iteradorVertices[T]) VerActual() T {
	nodo, _ := i.iter.VerActual()
	return nodo
}

// --- Iterador adyacentes ---
type iteradorAdyacentes[T comparable] struct {
	iter TDADict.IterDiccionario[T, float64]
}

func (g *grafoImp[T]) IterAdyacentes(nodo T) IteradorVertices[T] {
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

func (i *iteradorAdyacentes[T]) VerActual() T {
	nodo, _ := i.iter.VerActual()
	return nodo
}
