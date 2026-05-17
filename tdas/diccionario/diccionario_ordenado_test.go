package diccionario_test

import (
	"cmp"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ERR_NO_PERTENECE string = "La clave no pertenece al diccionario"

func TestDiccionarioOrdVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, _ERR_NO_PERTENECE, func() { dic.Obtener("A") })
	require.PanicsWithValue(t, _ERR_NO_PERTENECE, func() { dic.Borrar("A") })
}

func TestDiccionarioOrdClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](cmp.Compare)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElementOrd(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatoOrd(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioOrdBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	dic := TDADiccionario.CrearABB[int, string](cmp.Compare)
	dic.Guardar(5, "cinco")
	dic.Guardar(3, "tres")
	dic.Guardar(8, "ocho")
	dic.Guardar(1, "uno")
	dic.Guardar(4, "cuatro")
	dic.Guardar(7, "siete")
	dic.Guardar(10, "diez")
	require.Equal(t, 7, dic.Cantidad())

	require.True(t, dic.Pertenece(1))
	require.Equal(t, "uno", dic.Borrar(1))
	require.False(t, dic.Pertenece(1))
	require.Equal(t, 6, dic.Cantidad())

	require.True(t, dic.Pertenece(3))
	require.EqualValues(t, "tres", dic.Borrar(3))
	require.False(t, dic.Pertenece(3))
	require.True(t, dic.Pertenece(4))
	require.EqualValues(t, 5, dic.Cantidad())

	require.True(t, dic.Pertenece(8))
	require.EqualValues(t, "ocho", dic.Borrar(8))
	require.False(t, dic.Pertenece(8))

	require.True(t, dic.Pertenece(7))
	require.True(t, dic.Pertenece(10))
	require.EqualValues(t, "siete", dic.Obtener(7))
	require.EqualValues(t, "diez", dic.Obtener(10))
	require.EqualValues(t, 4, dic.Cantidad())
}

func TestReutlizacionDeBorradosOrd(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericasOrd(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](cmp.Compare)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestConClavesStructsOrd(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	obtener_valor_basico := func(bs basico) int {
		return bs.b + len(bs.a)
	}

	obtener_valor_avanzado := func(av avanzado) int {
		return av.w + obtener_valor_basico(av.x) + obtener_valor_basico(av.y) + len(av.z)
	}

	comparador := func(a avanzado, b avanzado) int {
		valor_a := obtener_valor_avanzado(a)
		valor_b := obtener_valor_avanzado(b)

		if valor_a > valor_b {
			return 1
		}

		if valor_a < valor_b {
			return -1
		}

		return 0
	}

	dic := TDADiccionario.CrearABB[avanzado, int](comparador)

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestClaveVaciaOrd(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNuloOrd(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](cmp.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestIteradorInternoClavesOrdenadas(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas una única vez de forma ordenada con el iterador interno")
	clave1 := "A"
	clave2 := "B"
	clave3 := "C"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		require.Equal(t, claves[*cantPtr], clave)
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
}

func TestIteradorInternoValoresOrdenados(t *testing.T) {
	t.Log("Valida que los datos sean recorridos correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoConCorte(t *testing.T) {
	t.Log("Valida que el iterador interno corte la ejecución inmediatamente al retornar false")
	dic := TDADiccionario.CrearABB[int, string](cmp.Compare)

	dic.Guardar(3, "tres")
	dic.Guardar(1, "uno")
	dic.Guardar(5, "cinco")
	dic.Guardar(2, "dos")
	dic.Guardar(4, "cuatro")

	contador := 0
	dic.Iterar(func(clave int, valor string) bool {
		contador++
		if clave == 3 {
			return false
		}
		return true
	})

	require.Equal(t, 3, contador)
}

func TestIteradorInternoRango(t *testing.T) {
	t.Log("Prueba el iterador interno por rango con diferentes acotaciones y condiciones de corte")
	dic := TDADiccionario.CrearABB[int, string](cmp.Compare)

	dic.Guardar(6, "seis")
	dic.Guardar(3, "tres")
	dic.Guardar(8, "ocho")
	dic.Guardar(2, "dos")
	dic.Guardar(5, "cinco")
	dic.Guardar(7, "siete")
	dic.Guardar(9, "nueve")
	dic.Guardar(1, "uno")
	dic.Guardar(4, "cuatro")
	dic.Guardar(10, "diez")

	desdeA := 4
	hastaA := 7
	rangoA := []int{4, 5, 6, 7}
	indiceA := 0

	dic.IterarRango(&desdeA, &hastaA, func(clave int, valor string) bool {
		require.Equal(t, rangoA[indiceA], clave)
		indiceA++
		return true
	})
	require.Equal(t, len(rangoA), indiceA)

	hastaB := 3
	rangoB := []int{1, 2, 3}
	indiceB := 0

	dic.IterarRango(nil, &hastaB, func(clave int, valor string) bool {
		require.Equal(t, rangoB[indiceB], clave)
		indiceB++
		return true
	})
	require.Equal(t, len(rangoB), indiceB)

	desdeC := 2
	hastaC := 9
	contadorC := 0

	dic.IterarRango(&desdeC, &hastaC, func(clave int, valor string) bool {
		contadorC++
		if clave == 5 {
			return false
		}
		return true
	})
	require.Equal(t, 4, contadorC)
}

func TestIteradorExternoCompleto(t *testing.T) {
	t.Log("Prueba el funcionamiento del iterador externo común recorriendo todo el árbol")
	dic := TDADiccionario.CrearABB[int, string](cmp.Compare)

	dic.Guardar(6, "seis")
	dic.Guardar(3, "tres")
	dic.Guardar(8, "ocho")
	dic.Guardar(2, "dos")
	dic.Guardar(5, "cinco")
	dic.Guardar(7, "siete")
	dic.Guardar(9, "nueve")
	dic.Guardar(1, "uno")
	dic.Guardar(4, "cuatro")
	dic.Guardar(10, "diez")

	require.Equal(t, 10, dic.Cantidad())

	iter := dic.Iterador()

	for i := 1; i <= 10; i++ {
		require.True(t, iter.HayAlgoMas())

		clave, _ := iter.VerActual()
		require.Equal(t, i, clave)
		iter.Avanzar()
	}

	require.False(t, iter.HayAlgoMas())

	require.Panics(t, func() { iter.VerActual() }, "Llamar a VerActual al final del iterador debe paniquear")
	require.Panics(t, func() { iter.Avanzar() }, "Llamar a Avanzar al final del iterador debe paniquear")

	require.EqualValues(t, 10, dic.Cantidad())
}
