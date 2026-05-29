package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

type hashMapAbierto[K comparable, V any] struct {
	datos    []TDALista.Lista[*claveValor[K, V]]
	cantidad int
}

type claveValor[K comparable, V any] struct {
	clave K
	valor V
}

type iterDiccionario[K comparable, V any] struct {
	hash      *hashMapAbierto[K, V]
	actual    int
	iterLista TDALista.IteradorLista[*claveValor[K, V]]
}

func crearClaveValor[K comparable, V any](clave K, dato V) *claveValor[K, V] {
	return &claveValor[K, V]{clave, dato}
}

// ----Primitivas----
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashMapAbierto[K, V]{
		make([]TDALista.Lista[*claveValor[K, V]], _TAMAÑO_INICIAL), //Las posiciones del array apuntan a nil
		0}
}

func (h *hashMapAbierto[K, V]) Guardar(clave K, dato V) {
	par := h.buscarClaveValor(clave)

	if par != nil {
		par.valor = dato
		return
	}

	par = crearClaveValor(clave, dato)
	h.colocarValor(par, h.datos)
	h.cantidad++ //Solo si no pisa el valor

	if h.obtenerFactorCarga() >= _TECHO_CARGA {
		h.redimensionar(len(h.datos) * _PROPORCION_REDIMENSION)
	}
}

func (h *hashMapAbierto[K, V]) Pertenece(clave K) bool {
	return h.buscarClaveValor(clave) != nil
}

func (h *hashMapAbierto[K, V]) Obtener(clave K) V {
	par := h.buscarClaveValor(clave)
	if par == nil {
		panic(_ERR_NO_PERTENECE)
	}

	return par.valor
}

func (h *hashMapAbierto[K, V]) Borrar(clave K) V {
	i := fnvHashing(clave, len(h.datos))
	lista := h.datos[i]

	if lista == nil {
		panic(_ERR_NO_PERTENECE)
	}

	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		actual := iter.VerActual()

		if actual.clave == clave {
			h.cantidad--
			dato := iter.Borrar().valor
			if lista.EstaVacia() {
				h.datos[i] = nil
			}
			if h.obtenerFactorCarga() <= _PISO_CARGA && len(h.datos) > _TAMAÑO_INICIAL {
				h.redimensionar(len(h.datos) / _PROPORCION_REDIMENSION)
			}
			return dato
		}
	}

	panic(_ERR_NO_PERTENECE)
}

func (h *hashMapAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

// ----Hashing----
func fnvHashing[K comparable](clave K, largo int) uint {
	h := _HASHING_INICIAL
	bytes := convertirABytes(clave)
	for _, b := range bytes {
		h ^= uint(b) //Potenciando antes de multiplicar se obtiene una mejor distribucion
		h *= _HASHING_MULTIPLICADOR
	}
	return h % uint(largo)
}

func convertirABytes[K comparable](clave K) []byte {
	return fmt.Appendf(nil, "%v", clave)
}

// ----Redimension----

func (h *hashMapAbierto[K, V]) obtenerFactorCarga() float64 {
	return float64(h.cantidad) / float64(len(h.datos))
}

func (h *hashMapAbierto[K, V]) redimensionar(nuevo_largo int) {
	nuevaTabla := make([]TDALista.Lista[*claveValor[K, V]], nuevo_largo)

	for _, lista := range h.datos {
		if lista == nil {
			continue
		}

		for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
			par := iter.VerActual()
			h.colocarValor(par, nuevaTabla)
		}
	}

	h.datos = nuevaTabla
}

// ----Patrones reutilizables----

func (h *hashMapAbierto[K, V]) buscarClaveValor(clave K) *claveValor[K, V] {
	i := fnvHashing(clave, len(h.datos))
	lista := h.datos[i]

	if lista == nil {
		return nil
	}

	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		actual := iter.VerActual()

		if actual.clave == clave {
			return actual
		}

	}

	return nil
}

func (h *hashMapAbierto[K, V]) colocarValor(par *claveValor[K, V], tabla []TDALista.Lista[*claveValor[K, V]]) {
	i := fnvHashing(par.clave, len(tabla))

	if tabla[i] == nil {
		tabla[i] = TDALista.CrearListaEnlazada[*claveValor[K, V]]()
	}
	tabla[i].InsertarPrimero(par)
}

// ----Iteradores----
func (h *hashMapAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	for i, lista := range h.datos {
		if lista != nil {
			return &iterDiccionario[K, V]{h, i, lista.Iterador()}
		}
	}
	return &iterDiccionario[K, V]{h, len(h.datos), nil} // diccionario vacio
}

func (i *iterDiccionario[K, V]) HayAlgoMas() bool {
	if i.iterLista == nil {
		return false
	}
	return i.iterLista.HayAlgoMas()
}

func (i *iterDiccionario[K, V]) Avanzar() {
	if !i.HayAlgoMas() {
		panic(_ERR_ITER_TERMINO)
	}
	i.iterLista.Avanzar()

	if !i.iterLista.HayAlgoMas() {
		i.actual++
		for i.actual < len(i.hash.datos) {
			if i.hash.datos[i.actual] != nil {
				i.iterLista = i.hash.datos[i.actual].Iterador()
				return
			}
			i.actual++
		}
		i.iterLista = nil
	}
}

func (i *iterDiccionario[K, V]) VerActual() (K, V) {
	if !i.HayAlgoMas() {
		panic(_ERR_ITER_TERMINO)
	}
	par := i.iterLista.VerActual()
	return par.clave, par.valor
}

func (h *hashMapAbierto[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	for _, lista := range h.datos {
		if lista == nil {
			continue
		}
		for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
			par := iter.VerActual()
			if !visitar(par.clave, par.valor) {
				return
			}
		}
	}
}
