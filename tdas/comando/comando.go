package comando

type Comando interface {
	//Ejecuta la funcion principal del comando
	Ejecutar([]string)

	//Verifica si la cantidad de argumentos para el comando es correcta.
	//Si no lo es, imprime un mensaje de error y devuelve false
	CantidadArgumentosCorrecta([]string) bool
}
