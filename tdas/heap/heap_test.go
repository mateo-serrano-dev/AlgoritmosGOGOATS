package heap_test

import (
	"cmp"
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

/*
func TestDesencolarEncolarStrings(t *testing.T) {
	heap := TDAHeap.CrearHeap[string](string.Compare)
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
*/
