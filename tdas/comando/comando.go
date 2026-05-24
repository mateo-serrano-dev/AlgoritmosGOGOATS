package comando

type Comando interface {
	Ejecutar([]string)
}
