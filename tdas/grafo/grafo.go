package grafo

type Grafo[T any] interface {
	AgregarVertice(T)
	BorrarVertice(Nodo[T])
	BorrarArista(Nodo[T], Nodo[T])

	IterVertices() IteradorVertices[T]
	IterAdyacentes(Nodo[T]) IteradorVertices[T]
}

type GrafoPesado[T any] interface {
	Grafo[T]

	AgregarArista(Nodo[T], Nodo[T], int)
	Peso(Nodo[T], Nodo[T]) int
}

type GrafoNoPesado[T any] interface {
	Grafo[T]
	AgregarArista(Nodo[T], Nodo[T])
}

type Nodo[T any] interface {
	Dato() T
}

type IteradorVertices[T any] interface {
	Avanzar()
	VerActual() Nodo[T]
	HayAlgoMas() bool
}
