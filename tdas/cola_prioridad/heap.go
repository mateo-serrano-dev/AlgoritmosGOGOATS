package cola_prioridad

const _TAMAÑO_INICIAL int = 16
const _PROPORCION_MINIMA int = 4
const _PROPORCION_REDIMENSION int = 2
const _ERROR_VACIA = "La cola esta vacia"

type heap[T any] struct {
	arr      []T
	comparar func(T, T) int
	cantidad int
}

// Funcion de comparar (a, b)
// negativo si a < b
// 0 si a == b
// positivo si a > b

// --- Creacion ---
func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{make([]T, _TAMAÑO_INICIAL), funcion_cmp, 0}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	nuevoArreglo := make([]T, len(arreglo))
	copy(nuevoArreglo, arreglo)
	heap := heap[T]{nuevoArreglo, funcion_cmp, len(nuevoArreglo)}
	heap.heapify()
	return &heap
}

// --- Primitivas ---
func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elemento T) {
	if len(h.arr) <= h.cantidad {
		h.redimensionar(len(h.arr) * _PROPORCION_REDIMENSION)
	}

	h.arr[h.cantidad] = elemento
	h.upheap(h.cantidad)
	h.cantidad++
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic(_ERROR_VACIA)
	}

	return h.arr[0]
}

func (h *heap[T]) Desencolar() T {
	result := h.VerMax()
	h.swap(&h.arr[0], &h.arr[h.cantidad-1])
	h.cantidad--
	h.downheap(0)

	if len(h.arr) > _TAMAÑO_INICIAL && h.cantidad <= len(h.arr)/_PROPORCION_MINIMA {
		h.redimensionar(len(h.arr) / _PROPORCION_REDIMENSION)
	}

	return result
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

// --- Acciones internas ---

func (h *heap[T]) upheap(i int) {
	for i > 0 {
		padre := padre(i)
		if h.comparar(h.arr[i], h.arr[padre]) <= 0 {
			break
		}
		h.swap(&h.arr[i], &h.arr[padre])
		i = padre
	}
}

func (h *heap[T]) downheap(i int) {
	for posIzq := izq(i); posIzq < h.cantidad; posIzq = izq(i) {
		posDer := der(i)
		mayor := posIzq

		if posDer < h.cantidad && h.comparar(h.arr[posDer], h.arr[posIzq]) > 0 {
			mayor = posDer
		}

		if h.comparar(h.arr[mayor], h.arr[i]) <= 0 {
			break
		}

		h.swap(&h.arr[i], &h.arr[mayor])

		i = mayor
	}
}

func (h *heap[T]) redimensionar(nuevaCapacidad int) {
	if nuevaCapacidad == 0 {
		nuevaCapacidad = _TAMAÑO_INICIAL
	}
	nuevoArreglo := make([]T, nuevaCapacidad)
	copy(nuevoArreglo, h.arr)
	h.arr = nuevoArreglo
}

// --- Miscelaneo ---
func (h *heap[T]) heapify() {
	for i := (len(h.arr) / 2) - 1; i >= 0; i-- {
		h.downheap(i)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	if len(elementos) < 2 {
		return
	}
	heap := &heap[T]{elementos, funcion_cmp, len(elementos)}
	heap.heapify()
	for heap.cantidad > 1 {
		heap.swap(&heap.arr[0], &heap.arr[heap.cantidad-1])
		heap.cantidad--
		heap.downheap(0)
	}
}

func (h *heap[T]) swap(x *T, y *T) {
	*x, *y = *y, *x
}

func padre(i int) int {
	return (i - 1) / 2
}

func izq(i int) int {
	return 2*i + 1
}

func der(i int) int {
	return izq(i) + 1
}
