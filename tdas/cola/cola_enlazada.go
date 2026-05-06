package cola

const _ERR_COLA_VACIA_ = "La cola esta vacia"

type nodo[T any] struct {
	valor     T
	siguiente *nodo[T]
}

func crearNodo[T any](t T) *nodo[T] {
	return &nodo[T]{t, nil}
}

type colaEnlazada[T any] struct {
	inicio *nodo[T]
	ultimo *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.inicio == nil && c.ultimo == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic(_ERR_COLA_VACIA_)
	}
	return c.inicio.valor
}

func (c *colaEnlazada[T]) Encolar(t T) {
	var nuevoNodo *nodo[T] = crearNodo(t)

	if c.EstaVacia() {
		c.inicio = nuevoNodo
	} else {
		c.ultimo.siguiente = nuevoNodo
	}

	c.ultimo = nuevoNodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	var valor T = c.VerPrimero()
	c.inicio = c.inicio.siguiente

	if c.inicio == nil {
		c.ultimo = nil
	}

	return valor
}
