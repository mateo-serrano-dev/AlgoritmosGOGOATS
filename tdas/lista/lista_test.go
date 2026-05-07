package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ERR_LISTA_VACIA string = "La lista esta vacia"
const _ERR_ITER_TERMINO string = "El iterador termino de iterar"

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.VerPrimero() })
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.VerUltimo() })
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())

}

func TestInsertar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("Yo")
	require.Equal(t, "Yo", lista.VerPrimero())
	require.Equal(t, "Yo", lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())

	lista.InsertarPrimero("Soy")
	require.Equal(t, "Soy", lista.VerPrimero())
	require.Equal(t, "Yo", lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())

	lista.InsertarUltimo("La")
	require.Equal(t, "Soy", lista.VerPrimero())
	require.Equal(t, "La", lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, "Soy", lista.BorrarPrimero())
	require.Equal(t, "Yo", lista.VerPrimero())
	require.Equal(t, "La", lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, "Yo", lista.BorrarPrimero())
	require.Equal(t, "La", lista.VerPrimero())
	require.Equal(t, "La", lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, "La", lista.BorrarPrimero())
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.VerPrimero() })
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.VerUltimo() })
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())

	lista.InsertarPrimero("Verdad")
	lista.InsertarPrimero("Y")
	lista.InsertarPrimero("La vida")
	require.Equal(t, "La vida", lista.VerPrimero())
	require.Equal(t, "Verdad", lista.VerUltimo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
}

func TestVolumen(t *testing.T) {
	iteraciones := 10000
	lista := TDALista.CrearListaEnlazada[int]()

	for i := range iteraciones {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, iteraciones, lista.Largo())
	require.False(t, lista.EstaVacia())
	require.Equal(t, iteraciones-1, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())

	for i := range iteraciones {
		require.Equal(t, iteraciones-i-1, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.VerPrimero() })
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.VerUltimo() })
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.BorrarPrimero() })

	for i := range iteraciones {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, iteraciones, lista.Largo())
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, iteraciones-1, lista.VerUltimo())

	for i := range iteraciones {
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.VerPrimero() })
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.VerUltimo() })
	require.PanicsWithValue(t, _ERR_LISTA_VACIA, func() { _ = lista.BorrarPrimero() })
}

func TestInsertarListaVaciaConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.VerActual() })

	iter.Insertar(1)

	require.Equal(t, 1, iter.VerActual())
}

func TestInsertarElementoPrimeraIteracion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	require.Equal(t, 1, iter.VerActual())
	iter.Insertar(10)

	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 10, iter.VerActual())
	iter.Avanzar()
	require.Equal(t, 1, iter.VerActual())
}

func TestInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("a")
	lista.InsertarUltimo("c")
	lista.InsertarUltimo("d")

	iter := lista.Iterador()
	iter.Avanzar()
	require.Equal(t, "c", iter.VerActual())

	iter.Insertar("b")
	require.Equal(t, "b", iter.VerActual())
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, "a", lista.VerPrimero())
	require.Equal(t, "d", lista.VerUltimo())

	iter2 := lista.Iterador()
	require.Equal(t, "a", iter2.VerActual())
	iter2.Avanzar()
	require.Equal(t, "b", iter2.VerActual())
	iter2.Avanzar()
	require.Equal(t, "c", iter2.VerActual())
	iter2.Avanzar()
	require.Equal(t, "d", iter2.VerActual())
	iter2.Avanzar()

	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter2.Avanzar() })
}

func TestInsertarUltimaIteracion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	lista.InsertarUltimo(0.1)
	lista.InsertarUltimo(0.2)
	lista.InsertarUltimo(0.3)

	iter := lista.Iterador()
	iter.Avanzar()
	iter.Avanzar()
	iter.Avanzar()

	require.False(t, iter.HayAlgoMas())

	iter.Insertar(0.4)
	require.True(t, iter.HayAlgoMas())
	require.Equal(t, 0.4, lista.VerUltimo())
	require.Equal(t, 0.4, iter.VerActual())

	iter.Avanzar()
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.Avanzar() })
}

func TestBorrarElementoPrimeraIteracion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("a")
	lista.InsertarUltimo("b")
	lista.InsertarUltimo("c")

	iter := lista.Iterador()
	require.Equal(t, "a", iter.VerActual())

	require.Equal(t, "a", iter.Borrar())
	require.Equal(t, "b", iter.VerActual())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, "b", lista.VerPrimero())
	require.Equal(t, "c", lista.VerUltimo())

	iter.Avanzar()
	require.Equal(t, "c", iter.VerActual())
	iter.Avanzar()
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.VerActual() })
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.Avanzar() })
}

func TestBorrarElementoFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("a")
	lista.InsertarUltimo("b")
	lista.InsertarUltimo("c")

	iter := lista.Iterador()

	iter.Avanzar()
	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	require.Equal(t, "c", iter.Borrar())
	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.Borrar() })
	require.Equal(t, "b", lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.VerActual() })
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.Avanzar() })
}

func TestBorrarElementoMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("a")
	lista.InsertarUltimo("b")
	lista.InsertarUltimo("c")

	iter := lista.Iterador()

	iter.Avanzar()

	require.Equal(t, "b", iter.Borrar())
	require.Equal(t, "c", iter.VerActual())
	require.True(t, iter.HayAlgoMas())
	require.Equal(t, 2, lista.Largo())

	iter2 := lista.Iterador()
	require.Equal(t, "a", iter2.VerActual())
	iter2.Avanzar()
	require.Equal(t, "c", iter2.VerActual())
	iter2.Avanzar()
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter2.VerActual() })
}

func TestBorrarUnicoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("a")

	iter := lista.Iterador()
	require.Equal(t, "a", iter.Borrar())
	require.False(t, iter.HayAlgoMas())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.VerActual() })
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.Avanzar() })
	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.Borrar() })
}

func TestBorrarListaVaciaConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.PanicsWithValue(t, _ERR_ITER_TERMINO, func() { iter.Borrar() })
}

func TestIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	res := []int{}
	lista.Iterar(func(valor int) bool {
		res = append(res, valor)
		return true
	})
	resEsperado := []int{1, 2, 3, 4, 5}

	require.Equal(t, resEsperado, res)
}

func TestIteradorInternoCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	res := []int{}
	lista.Iterar(func(valor int) bool {
		res = append(res, valor)
		return valor != 3
	})
	resEsperado := []int{1, 2, 3}
	require.Equal(t, resEsperado, res)
}
