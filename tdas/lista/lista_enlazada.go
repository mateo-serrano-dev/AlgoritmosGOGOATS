package lista

const _ERR_LISTA_VACIA string = "La lista esta vacia"
const _ERR_ITER_TERMINO string = "El iterador termino de iterar"

type nodo[T any] struct {
	valor     T
	siguiente *nodo[T]
}

type listaEnlazada[T any] struct {
	inicio *nodo[T]
	ultimo *nodo[T]
	largo  int
}

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodo[T]
	anterior *nodo[T]
}

func crearNodo[T any](t T) *nodo[T] {
	return &nodo[T]{t, nil}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.inicio == nil || l.ultimo == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodo(dato)
	nuevoNodo.siguiente = l.inicio
	l.inicio = nuevoNodo
	if l.EstaVacia() {
		l.ultimo = nuevoNodo
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodo(dato)
	if l.EstaVacia() {
		l.inicio = nuevoNodo
	} else {
		l.ultimo.siguiente = nuevoNodo
	}
	l.ultimo = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	dato := l.VerPrimero()
	l.inicio = l.inicio.siguiente
	if l.inicio == nil {
		l.ultimo = nil
	}
	l.largo--
	return dato
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic(_ERR_LISTA_VACIA)
	}

	return l.inicio.valor
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic(_ERR_LISTA_VACIA)
	}

	return l.ultimo.valor
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := l.inicio; actual != nil; actual = actual.siguiente {
		if !visitar(actual.valor) {
			return
		}
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{
		lista:    l,
		anterior: nil,
		actual:   l.inicio, // nil si la lista esta vacía
	}
}

func (i *iteradorLista[T]) HayAlgoMas() bool {
	return i.actual != nil
}

func (i *iteradorLista[T]) VerActual() T {
	if !i.HayAlgoMas() {
		panic(_ERR_ITER_TERMINO)
	}
	return i.actual.valor
}

func (i *iteradorLista[T]) Avanzar() {
	if !i.HayAlgoMas() {
		panic(_ERR_ITER_TERMINO)
	}
	i.anterior = i.actual
	i.actual = i.actual.siguiente
}

func (i *iteradorLista[T]) Insertar(dato T) {
	nodoNuevo := crearNodo(dato)
	if i.anterior == nil {
		nodoNuevo.siguiente = i.lista.inicio
		i.lista.inicio = nodoNuevo
	} else {
		nodoNuevo.siguiente = i.actual
		i.anterior.siguiente = nodoNuevo
	}

	if nodoNuevo.siguiente == nil {
		i.lista.ultimo = nodoNuevo
	}

	i.actual = nodoNuevo
	i.lista.largo++
}

func (i *iteradorLista[T]) Borrar() T {
	if !i.HayAlgoMas() {
		panic(_ERR_ITER_TERMINO)
	}
	dato := i.actual.valor
	if i.anterior == nil {
		i.lista.inicio = i.actual.siguiente
	} else {
		i.anterior.siguiente = i.actual.siguiente
	}

	if i.actual.siguiente == nil {
		i.lista.ultimo = i.anterior
	}
	i.actual = i.actual.siguiente
	i.lista.largo -= 1
	return dato
}
