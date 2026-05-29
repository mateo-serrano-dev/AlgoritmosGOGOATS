package diccionario

import TDAPila "tdas/pila"

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

func (abb *arbolBinarioBusqueda[K, V]) buscarMinimo(nodoActual *nodoAb[K, V]) *nodoAb[K, V] {
	for nodoActual.izq != nil {
		nodoActual = nodoActual.izq
	}
	return nodoActual
}

func (abb *arbolBinarioBusqueda[K, V]) borrar(nodo *nodoAb[K, V], clave K) (*nodoAb[K, V], V) {
	if nodo == nil {
		panic(_ERR_NO_PERTENECE)
	}

	comp := abb.comparar(nodo.clave, clave)
	var dato V
	if comp < 0 {
		nodo.der, dato = abb.borrar(nodo.der, clave)
		return nodo, dato
	} else if comp > 0 {
		nodo.izq, dato = abb.borrar(nodo.izq, clave)
		return nodo, dato
	}

	borrado := nodo.valor

	if nodo.izq == nil {
		abb.cantidad--
		return nodo.der, borrado
	}
	if nodo.der == nil {
		abb.cantidad--
		return nodo.izq, borrado
	}

	sucesor := abb.buscarMinimo(nodo.der)
	nodo.clave, nodo.valor = sucesor.clave, sucesor.valor
	nodo.der, _ = abb.borrar(nodo.der, sucesor.clave)

	return nodo, borrado
}

func (abb *arbolBinarioBusqueda[K, V]) Borrar(clave K) V {
	nuevaRaiz, dato := abb.borrar(abb.raiz, clave)
	abb.raiz = nuevaRaiz
	return dato
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
		panic(_ERR_ITER_TERMINO)
	}
	nodo := iter.pila.Desapilar().der
	iter.abb.apilarMenoresyActual(nodo, iter.desde, iter.hasta, iter.pila)
}

func (iter *iteradorDiccionarioOrdenado[K, V]) HayAlgoMas() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic(_ERR_ITER_TERMINO)
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().valor
}
