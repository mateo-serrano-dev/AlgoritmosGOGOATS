package diccionario

import TDAPila "tdas/pila"

const _ERR_NO_PERTENECE string = "La clave no pertenece al diccionario"
const _ERROR_ITER_TERMINO string = "El iterador termino de iterar"

type nodoAb[K comparable, V any] struct {
	izq   *nodoAb[K, V]
	der   *nodoAb[K, V]
	clave K
	valor V
}

type arbolBinarioBusqueda[K comparable, V any] struct {
	raiz     *nodoAb[K, V]
	comparar func(K, K) int
	cantidad int
}

func crearNodoAb[K comparable, V any](clave K, valor V) *nodoAb[K, V] {
	return &nodoAb[K, V]{
		izq:   nil,
		der:   nil,
		clave: clave,
		valor: valor,
	}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &arbolBinarioBusqueda[K, V]{
		raiz:     nil,
		comparar: funcion_cmp,
		cantidad: 0,
	}
}

func (abb *arbolBinarioBusqueda[K, V]) buscarPorClave(nodoActual *nodoAb[K, V], padre *nodoAb[K, V], clave K) (*nodoAb[K, V], *nodoAb[K, V]) {
	if nodoActual == nil {
		return padre, nodoActual
	}

	if abb.esMenor(nodoActual.clave, clave) {
		return abb.buscarPorClave(nodoActual.der, nodoActual, clave)
	}

	if abb.esMayor(nodoActual.clave, clave) {
		return abb.buscarPorClave(nodoActual.izq, nodoActual, clave)
	}

	return padre, nodoActual
}

func (abb *arbolBinarioBusqueda[K, V]) Guardar(clave K, valor V) {
	padre, hijo := abb.buscarPorClave(abb.raiz, nil, clave)
	if hijo != nil {
		// caso clave ya pertenece
		hijo.valor = valor
		return
	}
	nodoNuevo := crearNodoAb(clave, valor)
	if padre == nil {
		// caso arbol vacio
		abb.raiz = nodoNuevo
	} else if abb.esMayor(clave, padre.clave) {
		// caso clave nueva
		padre.der = nodoNuevo
	} else {
		padre.izq = nodoNuevo
	}
	abb.cantidad++
}

func (abb *arbolBinarioBusqueda[K, V]) Pertenece(clave K) bool {
	_, hijo := abb.buscarPorClave(abb.raiz, nil, clave)
	return hijo != nil
}

func (abb *arbolBinarioBusqueda[K, V]) Obtener(clave K) V {
	_, nodoEncontrado := abb.buscarPorClave(abb.raiz, nil, clave)
	if nodoEncontrado == nil {
		panic(_ERR_NO_PERTENECE)
	}
	return nodoEncontrado.valor
}

func (abb *arbolBinarioBusqueda[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *arbolBinarioBusqueda[K, V]) buscarMinimo(nodoActual, padre *nodoAb[K, V]) (*nodoAb[K, V], *nodoAb[K, V]) {
	for nodoActual.izq != nil {
		padre = nodoActual
		nodoActual = nodoActual.izq
	}
	return padre, nodoActual
}

func (abb *arbolBinarioBusqueda[K, V]) Borrar(clave K) V {
	padre, hijo := abb.buscarPorClave(abb.raiz, nil, clave)
	if hijo == nil {
		panic(_ERR_NO_PERTENECE)
	}

	borrado := hijo.valor

	// caso 2 hijos
	if hijo.der != nil && hijo.izq != nil {
		_, sucesor := abb.buscarMinimo(hijo.der, hijo)
		k := sucesor.clave
		v := abb.Borrar(k)
		hijo.clave = k
		hijo.valor = v
		return borrado
	}

	// casos de 0 o 1 hijo
	var reemplazo *nodoAb[K, V]
	if hijo.izq == nil {
		reemplazo = hijo.der
	} else {
		reemplazo = hijo.izq
	}

	if padre == nil {
		abb.raiz = reemplazo
	} else if padre.izq == hijo {
		padre.izq = reemplazo
	} else {
		padre.der = reemplazo
	}
	abb.cantidad--
	return borrado
}

// --- Comparacion ---
func (abb *arbolBinarioBusqueda[K, V]) esMayor(a K, b K) bool {
	return abb.comparar(a, b) > 0
}

func (abb *arbolBinarioBusqueda[K, V]) esMenor(a K, b K) bool {
	return abb.comparar(a, b) < 0
}

func (abb *arbolBinarioBusqueda[K, V]) esIgual(a K, b K) bool {
	return abb.comparar(a, b) == 0
}

// ---Iteradores---
type iteradorDiccionarioOrdenado[K comparable, V any] struct {
	abb   *arbolBinarioBusqueda[K, V]
	pila  TDAPila.Pila[*nodoAb[K, V]]
	desde *K
	hasta *K
}

func (abb *arbolBinarioBusqueda[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, visitar)
}

func (abb *arbolBinarioBusqueda[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *arbolBinarioBusqueda[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	pila := TDAPila.CrearPilaDinamica[*nodoAb[K, V]]()
	abb.apilarMenoresyActual(abb.raiz, desde, hasta, pila)
	for !pila.EstaVacia() {
		nodo := pila.Desapilar()
		if !visitar(nodo.clave, nodo.valor) {
			return
		}
		abb.apilarMenoresyActual(nodo.der, desde, hasta, pila)
	}
}

func (abb *arbolBinarioBusqueda[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iteradorDiccionarioOrdenado[K, V]{abb, TDAPila.CrearPilaDinamica[*nodoAb[K, V]](), desde, hasta}
	abb.apilarMenoresyActual(abb.raiz, desde, hasta, iter.pila)
	return iter
}

func (abb *arbolBinarioBusqueda[K, V]) apilarMenoresyActual(nodoActual *nodoAb[K, V], desde *K, hasta *K, pila TDAPila.Pila[*nodoAb[K, V]]) {
	for nodoActual != nil {
		if desde != nil && abb.esMenor(nodoActual.clave, *desde) {
			nodoActual = nodoActual.der
			continue
		}

		if hasta != nil && abb.esMayor(nodoActual.clave, *hasta) {
			nodoActual = nodoActual.izq
			continue
		}

		pila.Apilar(nodoActual)
		nodoActual = nodoActual.izq
	}
}

func (iter *iteradorDiccionarioOrdenado[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic(_ERROR_ITER_TERMINO)
	}
	nodo := iter.pila.Desapilar().der
	iter.abb.apilarMenoresyActual(nodo, iter.desde, iter.hasta, iter.pila)
}

func (iter *iteradorDiccionarioOrdenado[K, V]) HayAlgoMas() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic(_ERROR_ITER_TERMINO)
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().valor
}
