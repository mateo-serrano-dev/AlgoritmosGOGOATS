package diccionario

import TDAPila "tdas/pila"

const _ERR_NO_PERTENECE string = "La clave no pertenece al diccionario"

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

func (abb *arbolBinarioBusqueda[K, V]) guardar(nodoActual *nodoAb[K, V], clave K, valor V) *nodoAb[K, V] {
	//TODO reutilizar la busqueda por clave para guardar

	if nodoActual == nil {
		abb.cantidad++
		return crearNodoAb(clave, valor)
	}

	comparacion := abb.comparar(nodoActual.clave, clave)
	if comparacion < 0 {
		nodoActual.der = abb.guardar(nodoActual.der, clave, valor)
	} else if comparacion > 0 {
		nodoActual.izq = abb.guardar(nodoActual.izq, clave, valor)
	} else {
		nodoActual.valor = valor // Pisamos el valor anterior
	}

	return nodoActual
}

func (abb *arbolBinarioBusqueda[K, V]) Guardar(clave K, valor V) {
	abb.raiz = abb.guardar(abb.raiz, clave, valor)
}

/*
func (abb *arbolBinarioBusqueda[K, V]) Guardarmejora(clave K, valor V) {
	padre, hijo := abb.buscarPorClave(abb.raiz, nil, clave)

	if hijo != nil {
		hijo.valor = valor
	}

	abb.cantidad++

	if abb.esMayor(padre.clave, clave) {
		padre.izq
	}
}*/

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

func buscarMinimo[K comparable, V any](nodoActual *nodoAb[K, V]) *nodoAb[K, V] {
	for nodoActual.izq != nil {
		nodoActual = nodoActual.izq
	}
	return nodoActual
}

func (abb *arbolBinarioBusqueda[K, V]) borrarNodo(nodoABorrar *nodoAb[K, V]) (*nodoAb[K, V], V) {
	borrado := nodoABorrar.valor

	// Casos base (0 o 1 hijo)
	if nodoABorrar.izq == nil {
		abb.cantidad--
		return nodoABorrar.der, borrado // si nodoABorrar.der es nil, se pisa el nodo que queremos borrar con nil
	}
	if nodoABorrar.der == nil {
		abb.cantidad--
		return nodoABorrar.izq, borrado
	}

	// Reemplazamos con el menor de sus hijos derechos
	nodoReemplazo := buscarMinimo(nodoABorrar.der)

	// Este llamado llega al caso base (no es necesario hacer abb.cantidad--)
	nodoABorrar.der, _ = abb.borrar(nodoABorrar.der, nodoReemplazo.clave)

	// Pisamos los datos del nodo que queremos borrar
	nodoABorrar.clave, nodoABorrar.valor = nodoReemplazo.clave, nodoReemplazo.valor
	return nodoABorrar, borrado
}

func (abb *arbolBinarioBusqueda[K, V]) borrar(nodoActual *nodoAb[K, V], clave K) (*nodoAb[K, V], V) {
	//TODO se puede utilizar la funcion de buscar para hacer esto.

	if nodoActual == nil {
		panic(_ERR_NO_PERTENECE)
	}
	var valor V
	comparacion := abb.comparar(nodoActual.clave, clave)
	if comparacion < 0 {
		nodoActual.der, valor = abb.borrar(nodoActual.der, clave)
		return nodoActual, valor
	} else if comparacion > 0 {
		nodoActual.izq, valor = abb.borrar(nodoActual.izq, clave)
		return nodoActual, valor
	}
	return abb.borrarNodo(nodoActual)
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
		visitar(nodo.clave, nodo.valor)
		abb.apilarMenoresyActual(nodo.der, desde, hasta, pila)
	}
}

func (abb *arbolBinarioBusqueda[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := iteradorDiccionarioOrdenado[K, V]{abb, TDAPila.CrearPilaDinamica[*nodoAb[K, V]](), desde, hasta}
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

func (iter iteradorDiccionarioOrdenado[K, V]) Avanzar() {
	nodo := iter.pila.Desapilar().der
	iter.abb.apilarMenoresyActual(nodo, iter.desde, iter.hasta, iter.pila)
}

func (iter iteradorDiccionarioOrdenado[K, V]) HayAlgoMas() bool {
	return !iter.pila.EstaVacia()
}

func (iter iteradorDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	return iter.pila.VerTope().clave, iter.pila.VerTope().valor
}
