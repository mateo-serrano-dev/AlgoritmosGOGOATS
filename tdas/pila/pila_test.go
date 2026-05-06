package pila_test

import (
	"strings"
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ERR_PILA_VACIA = "La pila esta vacia"

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _ERR_PILA_VACIA, func() { _ = pila.VerTope() })
	require.PanicsWithValue(t, _ERR_PILA_VACIA, func() { _ = pila.Desapilar() })
}

func TestApilarDesapilar(t *testing.T) {
	elementos := strings.Split("Que tu fuerza comience donde termina la mia", " ")
	pila := TDAPila.CrearPilaDinamica[string]()

	for _, e := range elementos {
		pila.Apilar(e)
	}

	require.False(t, pila.EstaVacia())
	require.Equal(t, "mia", pila.VerTope())

	require.Equal(t, "mia", pila.Desapilar())
	require.Equal(t, "la", pila.Desapilar())
	require.Equal(t, "termina", pila.Desapilar())
	require.Equal(t, "donde", pila.Desapilar())
	require.False(t, pila.EstaVacia())

	pila.Apilar("hola mundovich")
	pila.Apilar("god")
	require.False(t, pila.EstaVacia())

	require.Equal(t, "god", pila.Desapilar())
	require.Equal(t, "hola mundovich", pila.Desapilar())
	require.Equal(t, "comience", pila.Desapilar())
	require.False(t, pila.EstaVacia())

	require.Equal(t, "fuerza", pila.Desapilar())
	require.Equal(t, "tu", pila.Desapilar())
	require.Equal(t, "Que", pila.Desapilar())

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _ERR_PILA_VACIA, func() { _ = pila.VerTope() })
	require.PanicsWithValue(t, _ERR_PILA_VACIA, func() { _ = pila.Desapilar() })

	pila.Apilar("boquita el mas grande")
	require.False(t, pila.EstaVacia())
	require.Equal(t, "boquita el mas grande", pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _ERR_PILA_VACIA, func() { _ = pila.VerTope() })
	require.PanicsWithValue(t, _ERR_PILA_VACIA, func() { _ = pila.Desapilar() })
}

func TestApilarDesapilarVolumen(t *testing.T) {
	var iteraciones int = 1000
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := range iteraciones {
		pila.Apilar(i)
		require.False(t, pila.EstaVacia())
		require.Equal(t, i, pila.VerTope())
	}
	require.False(t, pila.EstaVacia())

	for i := range iteraciones {
		require.False(t, pila.EstaVacia())
		valorEsperado := iteraciones - (i + 1)
		require.Equal(t, valorEsperado, pila.VerTope())
		require.Equal(t, valorEsperado, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, _ERR_PILA_VACIA, func() { _ = pila.VerTope() })
	require.PanicsWithValue(t, _ERR_PILA_VACIA, func() { _ = pila.Desapilar() })
}
