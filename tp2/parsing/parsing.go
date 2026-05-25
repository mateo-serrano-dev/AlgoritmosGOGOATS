package parsing

import (
	"fmt"
	"strings"
	Comando "tdas/comando"
	Dict "tdas/diccionario"
	Constantes "tp2/constantes"
)

func ParsearComando(linea string, comandos Dict.Diccionario[string, Comando.Comando]) (Comando.Comando, []string) {
	dividido := strings.Split(linea, ":")
	nombre, args := dividido[0], dividido[1]

	if !comandos.Pertenece(nombre) {
		fmt.Printf(Constantes.ENOENT_CMD, nombre)
	}

	return comandos.Obtener(nombre), strings.Split(args, ",")
}
