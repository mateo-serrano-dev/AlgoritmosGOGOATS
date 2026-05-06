package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const _HASHING_INICIAL uint = 14695981039346656037
const _HASHING_MULTIPLIER uint = 1099511628211
const _TAMAÑO_INICIAL int = 16
const _NOT_ON_HASHMAP string = "La clave no pertenece al diccionario"
const _ERR_ITER_TERMINO string = "El iterador termino de iterar"

const _TECHO_CARGA float64 = 3
const _PISO_CARGA float64 = 0.25
const _PROPORCION_REDIMENSION int = 2

type HashMapAbierto[K comparable, V any] struct {
	datos    []TDALista.Lista[*claveValor[K, V]]
	cantidad int
}

type claveValor[K comparable, V any] struct {
	clave K
	valor V
}

func crearClaveValor[K comparable, V any](clave K, dato V) *claveValor[K, V] {
	return &claveValor[K, V]{clave, dato}
}

// ----Primitivas----
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &HashMapAbierto[K, V]{
		make([]TDALista.Lista[*claveValor[K, V]], _TAMAÑO_INICIAL), //Las posiciones del array apuntan a nil
		0}
}

func (h *HashMapAbierto[K, V]) Guardar(clave K, dato V) {
	h.verificarRedimension()
	par := h.buscarClaveValor(clave)

	if par != nil {
		par.valor = dato
		return
	}

	par = crearClaveValor(clave, dato)
	h.colocarValor(par, h.datos)
	h.cantidad += 1 //Solo si no pisa el valor
}

func (h *HashMapAbierto[K, V]) Pertenece(clave K) bool {
	return h.buscarClaveValor(clave) != nil
}

func (h *HashMapAbierto[K, V]) Obtener(clave K) V {
	par := h.buscarClaveValor(clave)
	if par == nil {
		panic(_NOT_ON_HASHMAP)
	}

	return par.valor
}

func (h *HashMapAbierto[K, V]) Borrar(clave K) V {
	h.verificarRedimension()
	i := fnvHashing(clave, len(h.datos))
	lista := h.datos[i]

	if lista == nil {
		panic(_NOT_ON_HASHMAP)
	}

	iterador := lista.Iterador()
	for iterador.HayAlgoMas() {
		actual := iterador.VerActual()

		if actual.clave == clave {
			h.cantidad--
			return iterador.Borrar().valor
		}

		iterador.Avanzar()
	}

	panic(_NOT_ON_HASHMAP)
}

func (h *HashMapAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

// ----Hashing----
func fnvHashing[K comparable](clave K, largo int) uint {
	h := _HASHING_INICIAL
	bytes := convertirABytes(clave)
	for _, b := range bytes {
		h ^= uint(b) //Potenciando antes de multiplicar se obtiene una mejor distribucion
		h *= _HASHING_MULTIPLIER
	}
	return h % uint(largo)
}

func convertirABytes[K comparable](clave K) []byte {
	return fmt.Appendf(nil, "%v", clave)
}

// ----Redimension----

func (h *HashMapAbierto[K, V]) obtenerFactorCarga() float64 {
	return float64(h.cantidad) / float64(len(h.datos))
}

func (h *HashMapAbierto[K, V]) redimensionar(nuevo_largo int) {
	nueva_tabla := make([]TDALista.Lista[*claveValor[K, V]], nuevo_largo)

	for _, lista := range h.datos {
		if lista == nil {
			continue
		}

		iterador := lista.Iterador()
		for iterador.HayAlgoMas() {
			par := iterador.VerActual()
			h.colocarValor(par, nueva_tabla)
			iterador.Avanzar()
		}
	}

	h.datos = nueva_tabla
}

func (h *HashMapAbierto[K, V]) verificarRedimension() {
	factor := h.obtenerFactorCarga()

	if factor >= _TECHO_CARGA {
		h.redimensionar(len(h.datos) * _PROPORCION_REDIMENSION)
	}

	if factor <= _PISO_CARGA && len(h.datos) > _TAMAÑO_INICIAL {
		h.redimensionar(len(h.datos) / _PROPORCION_REDIMENSION)
	}
}

// ----Patrones reutilizables----

func (h *HashMapAbierto[K, V]) buscarClaveValor(clave K) *claveValor[K, V] {
	i := fnvHashing(clave, len(h.datos))
	lista := h.datos[i]

	if lista == nil {
		return nil
	}

	iterador := lista.Iterador()
	for iterador.HayAlgoMas() {
		actual := iterador.VerActual()

		if actual.clave == clave {
			return actual
		}

		iterador.Avanzar()
	}

	return nil
}

func (h *HashMapAbierto[K, V]) colocarValor(par *claveValor[K, V], tabla []TDALista.Lista[*claveValor[K, V]]) {
	i := fnvHashing(par.clave, len(tabla))

	if tabla[i] == nil {
		tabla[i] = TDALista.CrearListaEnlazada[*claveValor[K, V]]()
	}
	tabla[i].InsertarPrimero(par)
}
