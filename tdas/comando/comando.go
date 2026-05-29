package comando

type Comando interface {
	//Ejecuta la funcion principal del comando
	Ejecutar([]string)
}
