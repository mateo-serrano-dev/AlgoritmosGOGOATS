package diccionario

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

// type diccionarioOrdenado[K comparable, V any] struct {
// abb[K, V]
//
// }

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

func (abb *arbolBinarioBusqueda[K, V]) buscarPorClave(nodoActual *nodoAb[K, V], clave K) *nodoAb[K, V] {
	if nodoActual == nil {
		return nodoActual
	}

	comparacion := abb.comparar(nodoActual.clave, clave)

	if comparacion < 0 {
		return abb.buscarPorClave(nodoActual.der, clave)
	} else if comparacion > 0 {
		return abb.buscarPorClave(nodoActual.izq, clave)
	} else {
		return nodoActual
	}
}

func (abb *arbolBinarioBusqueda[K, V]) guardar(nodoActual *nodoAb[K, V], clave K, valor V) *nodoAb[K, V] {
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

func (abb *arbolBinarioBusqueda[K, V]) Pertenece(clave K) bool {
	return abb.buscarPorClave(abb.raiz, clave) != nil
}

func (abb *arbolBinarioBusqueda[K, V]) Obtener(clave K) V {
	nodoEncontrado := abb.buscarPorClave(abb.raiz, clave)
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
