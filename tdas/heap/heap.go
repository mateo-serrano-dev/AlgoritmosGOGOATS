package heap

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
	var nuevo_arreglo []T
	copy(nuevo_arreglo, arreglo)
	heap := heap[T]{nuevo_arreglo, funcion_cmp, len(nuevo_arreglo)}

	for i := len(arreglo) - 1; i >= 0; i-- {
		heap.downheap(i)
	}

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
	if h.cantidad <= 0 {
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
	for true {
		padre := padre(i)
		if padre < 0 {
			return
		}

		if h.comparar(h.arr[i], h.arr[padre]) > 0 {
			h.swap(&h.arr[i], &h.arr[padre])
			i = padre
			continue
		}

		return
	}
}

func (h *heap[T]) downheap(i int) {
	for true {
		var mayor int
		var menor int
		izq := izq(i)
		der := der(i)

		if h.comparar(h.arr[der], h.arr[izq]) < 0 {
			mayor = izq
			menor = der
		} else {
			mayor = der
			menor = izq
		}

		if mayor < h.cantidad && h.comparar(h.arr[i], h.arr[mayor]) < 0 {
			h.swap(&h.arr[i], &h.arr[mayor])
			i = mayor
			continue
		} else if menor < h.cantidad && h.comparar(h.arr[i], h.arr[menor]) < 0 {
			h.swap(&h.arr[i], &h.arr[menor])
			i = menor
			continue
		}

		break
	}
}

func (h *heap[T]) redimensionar(dimension int) {
	nuevo_arreglo := make([]T, dimension)
	for i := 0; i <= dimension; i++ {
		nuevo_arreglo[i] = h.arr[i]
	}
	h.arr = nuevo_arreglo
}

// --- Miscelaneo ---
func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {

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
