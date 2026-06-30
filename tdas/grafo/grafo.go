package grafo

type Grafo[T comparable] interface {
	AgregarVertice(T)
	BorrarVertice(T)
	BorrarArista(T, T)
	CantidadVertices() int
	CantidadAristas(T) int
	PerteneceVertice(T) bool

	IterVertices() IteradorVertices[T]
	IterAdyacentes(T) IteradorVertices[T]
}

type GrafoPesado[T comparable] interface {
	Grafo[T]

	AgregarArista(T, T, float64)
	Peso(T, T) float64
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
