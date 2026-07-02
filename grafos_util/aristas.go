package grafos_util

type Arista[T comparable] struct {
	v    T
	w    T
	peso float64
}

func CrearArista[T comparable](v, w T, peso float64) Arista[T] {
	return Arista[T]{
		v:    v,
		w:    w,
		peso: peso,
	}
}
