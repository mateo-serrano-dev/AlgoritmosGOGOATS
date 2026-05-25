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

func ParsearComando(linea string, comandos Dict.Diccionario[string, Comando.Comando]) (Comando.Comando, []string) {
	dividido := strings.Split(linea, ":")
	nombre, args := dividido[0], dividido[1]

	if !comandos.Pertenece(nombre) {
		fmt.Printf(Constantes.ENOENT_CMD, nombre)
	}

	return comandos.Obtener(nombre), strings.Split(args, ",")
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
