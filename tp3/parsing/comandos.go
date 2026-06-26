package parsing

import (
	"strings"
	CMD "tdas/comando"
	TDADict "tdas/diccionario"
)

func ParsearComando(linea string, comandos TDADict.Diccionario[string, CMD.Comando]) (CMD.Comando, []string, bool) {
	linea = strings.TrimSpace(linea)
	dividido := strings.Split(linea, " ")

	nombre, args := dividido[0], dividido[1:]
	if !comandos.Pertenece(nombre) {
		return nil, nil, true
	}

	return comandos.Obtener(nombre), args, false
}
