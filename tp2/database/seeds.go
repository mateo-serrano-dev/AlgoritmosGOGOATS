package database

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	Constantes "tp2/constantes"
)

func LeerArchivo(ruta string, db *Database, f func(string, *Database)) {
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf(Constantes.ENOENT_ARCHIVO, ruta)
		return
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		f(linea, db)
	}
}

func ParsearPaciente(linea string, db *Database) {
	dividido := strings.Split(linea, ",")
	nombre, año := dividido[0], dividido[1]

	año_en_int, err := strconv.Atoi(año)

	if err != nil {
		fmt.Printf(Constantes.ENOENT_ANIO, año)
		return
	}

	db.agregarPaciente(nombre, año_en_int)
}

func ParsearDoctor(linea string, db *Database) {
	dividido := strings.Split(linea, ",")
	nombre, especialidad := dividido[0], dividido[1]

	if !db.especialidades.Pertenece(especialidad) {
		db.agregarEspecialidad(especialidad)
	}

	db.agregarDoctor(nombre, especialidad)
}
