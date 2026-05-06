package operador

const PRIORIDAD_0 = 0
const PRIORIDAD_1 = 1
const PRIORIDAD_2 = 2
const SIN_PRIORIDAD = -1

type Suma struct{}
type Resta struct{}
type Multiplicacion struct{}
type Division struct{}
type Potenciacion struct{}
type ParentesisAbierto struct{}
type ParentesisCerrado struct{}

func (_ Suma) TieneAsociatividadDerecha() bool { return false }
func (_ Suma) Prioridad() int                  { return PRIORIDAD_0 }

func (_ Resta) TieneAsociatividadDerecha() bool { return false }
func (_ Resta) Prioridad() int                  { return PRIORIDAD_0 }

func (_ Multiplicacion) TieneAsociatividadDerecha() bool { return false }
func (_ Multiplicacion) Prioridad() int                  { return PRIORIDAD_1 }

func (_ Division) TieneAsociatividadDerecha() bool { return false }
func (_ Division) Prioridad() int                  { return PRIORIDAD_1 }

func (_ Potenciacion) TieneAsociatividadDerecha() bool { return true }
func (_ Potenciacion) Prioridad() int                  { return PRIORIDAD_2 }

func (_ ParentesisAbierto) TieneAsociatividadDerecha() bool { return false }
func (_ ParentesisAbierto) Prioridad() int                  { return SIN_PRIORIDAD }

func (_ ParentesisCerrado) TieneAsociatividadDerecha() bool { return false }
func (_ ParentesisCerrado) Prioridad() int                  { return SIN_PRIORIDAD }
