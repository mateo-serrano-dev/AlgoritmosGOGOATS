package grafo

type Grafo[T comparable] interface {

	//Se añade un vértice con un dato T. Si el dato ya existe en el grafo, no se añade ningun nodo nuevo.
	AgregarVertice(T)

	//Se borra el nodo de dato T. Si el nodo no existe previamente, lanzará un panic.
	BorrarVertice(T)

	//Se borra la arista desde el primer dato hacia el segundo. Si el grafo es no dirigido, también
	//se borrará la arista desde el segundo hacia el primero. Si alguno de los dos nodos no existe,
	//se lanzará un panic.
	BorrarArista(T, T)

	//Devuelve en tiempo constante la cantidad de nodos almacenada en el grafo.
	CantidadVertices() int

	//Devuelve en tiempo constante la cantidad de aristas salientes de un nodo.
	CantidadAristas(T) int

	//Verifica si el dato T existe en el grafo.
	PerteneceVertice(T) bool

	//Devuelve un iterador que permite iterar por todos los vertices que
	//existen en el grafo.
	IterVertices() IteradorVertices[T]

	//Devuelve un iterador que permite iterar todos los nodos adyacentes
	//al nodo de dato T. Si el nodo no existe, lanza un panic.
	IterAdyacentes(T) IteradorVertices[T]
}

type GrafoPesado[T comparable] interface {
	Grafo[T]

	//Agrega una arista desde el primer nodo hasta el segundo con un peso indicado. Si es no dirigido, también
	//agrega una arista desde el segundo hacia el primero. Si alguno de los dos datos no
	//existe, lanza un panic.
	AgregarArista(T, T, float64)

	//Devuelve en peso de la arista desde el primer dato hacia el segundo. Si alguno de
	//los dos datos no existe, lanza un panic.
	Peso(T, T) float64
}

type GrafoNoPesado[T comparable] interface {
	Grafo[T]

	//Agrega una arista desde el primer nodo hasta el segundo. Si es no dirigido, también
	//agrega una arista desde el segundo hacia el primero. Si alguno de los dos datos no
	//existe, lanza un panic.
	AgregarArista(T, T)
}

type IteradorVertices[T comparable] interface {
	Avanzar()
	VerActual() T
	HayAlgoMas() bool
}
