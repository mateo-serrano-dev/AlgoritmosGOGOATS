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

type HashMapAbierto[K comparable, V any] struct {
	datos    []TDALista.Lista[*claveValor[K, V]]
	cantidad int
}

type claveValor[K comparable, V any] struct {
	clave K
	valor V
}

func convertirABytes[K comparable](clave K) []byte {
	return fmt.Appendf(nil, "%v", clave)
}

func fnvHashing[K comparable](clave K, largo int) uint {
	h := _HASHING_INICIAL
	bytes := convertirABytes(clave)
	for _, b := range bytes {
		h *= _HASHING_MULTIPLIER
		h ^= uint(b)
	}
	return h % uint(largo)
}

func crearClaveValor[K comparable, V any](clave K, dato V) *claveValor[K, V] {
	return &claveValor[K, V]{clave, dato}
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &HashMapAbierto[K, V]{
		make([]TDALista.Lista[*claveValor[K, V]], _TAMAÑO_INICIAL),
		0}
}

// TODO: resizing
func (h *HashMapAbierto[K, V]) obtenerFactorCarga() float64 {
	return float64(h.cantidad) / float64(len(h.datos))
}

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

func (h *HashMapAbierto[K, V]) Guardar(clave K, dato V) {
	par := h.buscarClaveValor(clave)

	if par != nil {
		par.valor = dato
		return
	}

	i := fnvHashing(clave, len(h.datos))
	if h.datos[i] == nil {
		h.datos[i] = TDALista.CrearListaEnlazada[*claveValor[K, V]]()
	}

	par = crearClaveValor(clave, dato)
	h.datos[i].InsertarUltimo(par)
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
