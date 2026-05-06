package cola_test

import (
	"strings"
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ERR_COLA_VACIA = "La cola esta vacia"

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _ERR_COLA_VACIA, func() { _ = cola.VerPrimero() })
	require.PanicsWithValue(t, _ERR_COLA_VACIA, func() { _ = cola.Desencolar() })
}

func TestEncolarDesencolarStrings(t *testing.T) {
	elementos := strings.Split("Todo lo puedo en Cristo que me fortalece", " ")
	cola := TDACola.CrearColaEnlazada[string]()

	for _, elemento := range elementos {
		print(elemento)
		cola.Encolar(elemento)
		require.False(t, cola.EstaVacia())
		require.Equal(t, "Todo", cola.VerPrimero())
	}

	require.Equal(t, "Todo", cola.VerPrimero())
	require.Equal(t, "Todo", cola.Desencolar())
	require.Equal(t, "lo", cola.Desencolar())
	require.Equal(t, "puedo", cola.Desencolar())
	require.False(t, cola.EstaVacia())
	cola.Encolar("hola mundeanos")
	require.False(t, cola.EstaVacia())
	require.Equal(t, "en", cola.Desencolar())
	require.Equal(t, "Cristo", cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.Equal(t, "que", cola.Desencolar())
	require.Equal(t, "me", cola.Desencolar())
	require.Equal(t, "fortalece", cola.VerPrimero())
	require.Equal(t, "fortalece", cola.Desencolar())
	require.Equal(t, "hola mundeanos", cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _ERR_COLA_VACIA, func() { _ = cola.VerPrimero() })
	require.PanicsWithValue(t, _ERR_COLA_VACIA, func() { _ = cola.Desencolar() })

	cola.Encolar("planeta")
	cola.Encolar("vegeta")
	require.False(t, cola.EstaVacia())
	require.Equal(t, "planeta", cola.Desencolar())
	require.Equal(t, "vegeta", cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestApilarDesapilarVolumen(t *testing.T) {
	var iteraciones int = 1000
	cola := TDACola.CrearColaEnlazada[int]()

	for i := range iteraciones {
		cola.Encolar(i)
		require.False(t, cola.EstaVacia())
		require.Equal(t, 0, cola.VerPrimero())
	}

	for i := range iteraciones {
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())

		if !cola.EstaVacia() {
			require.Equal(t, i+1, cola.VerPrimero())
		}
	}

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, _ERR_COLA_VACIA, func() { _ = cola.VerPrimero() })
	require.PanicsWithValue(t, _ERR_COLA_VACIA, func() { _ = cola.Desencolar() })
}
