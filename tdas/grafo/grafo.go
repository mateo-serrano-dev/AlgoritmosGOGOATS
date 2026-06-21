package grafo

type Grafo[T comparable] interface {
	AgregarVertice(T)
	BorrarVertice(T)
	BorrarArista(T, T)
	CantidadVertices() int
	CantidadAristas(T) int

	IterVertices() IteradorVertices[T]
	IterAdyacentes(T) IteradorVertices[T]
}

type GrafoPesado[T comparable] interface {
	Grafo[T]

	AgregarArista(T, T, int)
	Peso(T, T) int
}

type GrafoNoPesado[T comparable] interface {
	Grafo[T]
	AgregarArista(T, T)
}

type IteradorVertices[T comparable] interface {
	Avanzar()
	VerActual() T
	HayAlgoMas() bool
}
