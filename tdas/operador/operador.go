package operador

type Operador interface {
	TieneAsociatividadDerecha() bool
	Prioridad() int
}
