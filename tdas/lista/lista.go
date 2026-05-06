package lista

type IteradorLista[T any] interface {
	// Devuelve el elemento en la posición actual del iterador
	// Lanza panic si el iterador terminó de iterar.
	VerActual() T

	// Devuelve true si quedan elementos por iterar.
	HayAlgoMas() bool

	// Mueve el puntero actual al siguiente elemento de la lista
	// Lanza panic si el iterador terminó de iterar.
	Avanzar()

	// Agrega un nuevo elemento en la posición actual del iterador
	// El nuevo elemento queda ubicado previo al elemento actual
	Insertar(T)

	// Elimina el elemento en la posición actual y lo devuelve
	// El iterador queda posicionado en el elemento siguiente al borrado
	// Lanza panic si el iterador terminó de iterar.
	Borrar() T
}

type Lista[T any] interface {
	// Devuelve true si la lista no tiene elementos
	EstaVacia() bool

	// Agrega un elemento al inicio de la lista
	InsertarPrimero(T)

	// Agrega un elemento al final de la Lista
	InsertarUltimo(T)

	// Borra y devuelve el primer elemento de la lista
	// Lanza un panic si la lista está vacía
	BorrarPrimero() T

	// Devuelve el primer elemento de la Lista
	// Lanza un panic si la lista está vacía
	VerPrimero() T

	// Devuelve el último elemento de la Lista
	// Lanza un panic si la lista está vacía
	VerUltimo() T

	// Devuelve la cantidad de elementos que almacena la lista
	Largo() int

	// Recorre todos los elementos de la lista aplicando la función visitar a cada elemento
	// Termina de iterar si visitar devuelve false o si llega al final de la lista
	Iterar(visitar func(T) bool)

	// Devuelve una instancia de IteradorLista para recorrer la lista de forma externa
	Iterador() IteradorLista[T]
}
