package heap_test

import (
	"cmp"
	"math/rand"
	"strings"
	TDAHeap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ERROR_VACIA = "La cola esta vacia"

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmp.Compare)
	require.EqualValues(t, true, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.VerMax() })
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.Desencolar() })
}

func TestDesencolarEncolarInts(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmp.Compare)
	require.EqualValues(t, true, heap.EstaVacia())
	heap.Encolar(3)
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())
	heap.Encolar(1)
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())
	heap.Encolar(11)
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 11, heap.VerMax())
	require.EqualValues(t, 11, heap.Desencolar())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())

	heap.Encolar(11)
	require.EqualValues(t, 11, heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())
	heap.Encolar(13)
	require.EqualValues(t, 13, heap.VerMax())
	require.EqualValues(t, 4, heap.Cantidad())
	heap.Encolar(111)
	require.EqualValues(t, 111, heap.VerMax())
	require.EqualValues(t, 5, heap.Cantidad())
	heap.Encolar(121)
	require.EqualValues(t, 121, heap.VerMax())
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, false, heap.EstaVacia())

	require.EqualValues(t, 121, heap.Desencolar())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 111, heap.Desencolar())
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 13, heap.Desencolar())
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 11, heap.Desencolar())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 3, heap.Desencolar())
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 1, heap.Desencolar())
	require.EqualValues(t, 0, heap.Cantidad())

	require.EqualValues(t, true, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.VerMax() })
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.Desencolar() })
}

func TestElementosRepetidos(t *testing.T) {
	heap := TDAHeap.CrearHeap[string](strings.Compare)
	heap.Encolar("Hola")
	heap.Encolar("Hola")
	heap.Encolar("Hola")
	heap.Encolar("Hola")

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, "Hola", heap.VerMax())

	require.EqualValues(t, "Hola", heap.Desencolar())
	require.EqualValues(t, 3, heap.Cantidad())

	require.EqualValues(t, "Hola", heap.Desencolar())
	require.EqualValues(t, 2, heap.Cantidad())

	require.EqualValues(t, "Hola", heap.Desencolar())
	require.EqualValues(t, 1, heap.Cantidad())

	require.EqualValues(t, "Hola", heap.Desencolar())
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.VerMax() })
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.Desencolar() })
}

func TestCrearHeapArr(t *testing.T) {
	arr := []float64{0.1, 0.6, 0.3, 0.2, 0.123, 0.45}
	heap := TDAHeap.CrearHeapArr[float64](arr, cmp.Compare)

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, 0.6, heap.VerMax())
	require.EqualValues(t, 0.6, heap.Desencolar())

	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 0.45, heap.VerMax())
	require.EqualValues(t, 0.45, heap.Desencolar())

	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 0.3, heap.VerMax())
	require.EqualValues(t, 0.3, heap.Desencolar())

	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 0.2, heap.VerMax())
	require.EqualValues(t, 0.2, heap.Desencolar())

	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 0.123, heap.VerMax())
	require.EqualValues(t, 0.123, heap.Desencolar())

	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 0.1, heap.VerMax())
	require.EqualValues(t, 0.1, heap.Desencolar())

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.VerMax() })
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.Desencolar() })

	require.EqualValues(t, []float64{0.1, 0.6, 0.3, 0.2, 0.123, 0.45}, arr)
}

func TestHeapVolumen(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmp.Compare)
	cantidadElementos := 10000
	elemDesordenados := rand.Perm(cantidadElementos)
	for _, elem := range elemDesordenados {
		heap.Encolar(elem)
	}
	require.EqualValues(t, cantidadElementos, heap.Cantidad())

	for j := cantidadElementos - 1; j >= 0; j-- {
		require.EqualValues(t, j, heap.VerMax())
		require.EqualValues(t, j, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.VerMax() })
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.Desencolar() })
}

func TestDesencolarEncolar(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmp.Compare)
	require.EqualValues(t, true, heap.EstaVacia())
	heap.Encolar(3)
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())
	heap.Encolar(1)
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())
	heap.Encolar(11)
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 11, heap.VerMax())
	require.EqualValues(t, 11, heap.Desencolar())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())

	heap.Encolar(11)
	require.EqualValues(t, 11, heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())
	heap.Encolar(13)
	require.EqualValues(t, 13, heap.VerMax())
	require.EqualValues(t, 4, heap.Cantidad())
	heap.Encolar(111)
	require.EqualValues(t, 111, heap.VerMax())
	require.EqualValues(t, 5, heap.Cantidad())
	heap.Encolar(121)
	require.EqualValues(t, 121, heap.VerMax())
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, false, heap.EstaVacia())

	require.EqualValues(t, 121, heap.Desencolar())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 111, heap.Desencolar())
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 13, heap.Desencolar())
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 11, heap.Desencolar())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 3, heap.Desencolar())
	require.EqualValues(t, false, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 1, heap.Desencolar())
	require.EqualValues(t, 0, heap.Cantidad())

	require.EqualValues(t, true, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.VerMax() })
	require.PanicsWithValue(t, _ERROR_VACIA, func() { heap.Desencolar() })
}

func TestHeapSortArrVacioYUnElem(t *testing.T) {
	arrVacio := []int{}
	TDAHeap.HeapSort[int](arrVacio, cmp.Compare)
	require.EqualValues(t, []int{}, arrVacio)

	arrUnElem := []string{"cola de prioridad"}
	TDAHeap.HeapSort[string](arrUnElem, strings.Compare)
	require.EqualValues(t, []string{"cola de prioridad"}, arrUnElem)
}

func TestHeapSortArregloDesordenado(t *testing.T) {
	arr := []int{54, 12, 89, 3, 45, 21, 76, 9, 67, 34}

	TDAHeap.HeapSort[int](arr, cmp.Compare)

	ordenado := []int{3, 9, 12, 21, 34, 45, 54, 67, 76, 89}
	require.EqualValues(t, ordenado, arr)
}

func TestHeapSortArregloYaOrdenado(t *testing.T) {
	arrFloats := []float64{1.12, 2.5, 5.89, 7.01, 12.4, 15.34, 22.8, 45.0, 67.21, 99.9}

	TDAHeap.HeapSort[float64](arrFloats, cmp.Compare)

	esperado := []float64{1.12, 2.5, 5.89, 7.01, 12.4, 15.34, 22.8, 45.0, 67.21, 99.9}
	require.EqualValues(t, esperado, arrFloats)
}

func TestHeapSortVolumen(t *testing.T) {
	t.Log("Probamos que HeapSort ordene correctamente un volumen grande de elementos desordenados")

	const CANTIDAD_ELEMENTOS = 5000

	arr := rand.Perm(CANTIDAD_ELEMENTOS)

	TDAHeap.HeapSort[int](arr, cmp.Compare[int])

	for j := range CANTIDAD_ELEMENTOS - 1 {
		require.True(t, arr[j] <= arr[j+1])
	}
}
