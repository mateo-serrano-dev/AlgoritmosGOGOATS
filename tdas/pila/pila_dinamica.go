package pila

const _MIN_SIZE = 16
const _ERR_PILA_VACIA = "La pila esta vacia"
const _REDIMENSION_PROPORCION = 2

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {

	return &pilaDinamica[T]{make([]T, _MIN_SIZE), 0}
}

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic(_ERR_PILA_VACIA)
	}

	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(t T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(p.cantidad * _REDIMENSION_PROPORCION)
	}

	p.datos[p.cantidad] = t
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	var valor T = p.VerTope()

	p.cantidad--
	if p.cantidad <= len(p.datos)/(_REDIMENSION_PROPORCION*2) && p.cantidad/_REDIMENSION_PROPORCION >= _MIN_SIZE {
		p.redimensionar(len(p.datos) / _REDIMENSION_PROPORCION)
	}

	return valor
}

func (p *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}
