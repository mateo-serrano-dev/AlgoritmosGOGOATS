package parsing

import (
	"strings"
	CMD "tdas/comando"
	TDADict "tdas/diccionario"
)

func ParsearComando(linea string, comandos TDADict.Diccionario[string, CMD.Comando]) (CMD.Comando, []string, bool) {
	linea = strings.TrimSpace(linea)
	dividido := strings.SplitN(linea, " ", 2)

	nombre, args := dividido[0], strings.Split(dividido[1], ", ")

	if !comandos.Pertenece(nombre) {
		return nil, nil, true
	}

	return comandos.Obtener(nombre), args, false
}
