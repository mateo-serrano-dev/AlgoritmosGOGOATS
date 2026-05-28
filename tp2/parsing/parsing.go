package parsing

import (
	"fmt"
	"strconv"
	"strings"
	Comando "tdas/comando"
	Dict "tdas/diccionario"
	Constantes "tp2/constantes"
	DB "tp2/database"
)

func ParsearComando(linea string, comandos Dict.Diccionario[string, Comando.Comando]) (Comando.Comando, []string, bool) {
	linea = strings.TrimSpace(linea)
	dividido := strings.SplitN(linea, ":", 2)
	if len(dividido) < 2 {
		fmt.Printf(Constantes.ENOENT_FORMATO, linea)
		return nil, nil, true
	}
	nombre, args := dividido[0], dividido[1]

	if !comandos.Pertenece(nombre) {
		fmt.Printf(Constantes.ENOENT_CMD, nombre)
		return nil, nil, true
	}

	var argsFinales []string
	if args != "" {
		argsFinales = strings.Split(args, ",")
	} else {
		argsFinales = []string{}
	}

	return comandos.Obtener(nombre), argsFinales, false
}

func ParsearPaciente(linea string, db *DB.Database) {
	dividido := strings.Split(linea, ",")
	nombre, año := dividido[0], dividido[1]

	año_en_int, err := strconv.Atoi(año)

	if err != nil {
		fmt.Printf(Constantes.ENOENT_ANIO, año)
		return
	}

	db.AgregarPaciente(nombre, año_en_int)
}

func ParsearDoctor(linea string, db *DB.Database) {
	dividido := strings.Split(linea, ",")
	nombre, especialidad := dividido[0], dividido[1]

	if !db.ExisteEspecialidad(especialidad) {
		db.AgregarEspecialidad(especialidad)
	}

	db.AgregarDoctor(nombre, especialidad)
}
