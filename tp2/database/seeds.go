package database

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	Constantes "tp2/constantes"
)

/*
	DOCTORES_SISTEMA = "%d doctor(es) en el sistema\n"
	INFORME_DOCTOR   = "%d: %s, especialidad %s, %d paciente(s) atendido(s)\n"

	ENOENT_CANT_PARAMS = "No se recibieron los 2 (dos) parametros: <archivo doctores> y <archivo pacientes>\n"
	ENOENT_ARCHIVO     = "No se pudo leer archivo %s\n"
	ENOENT_ANIO        = "Valor no numerico en campo de anio: %s\n"*/

func LeerArchivo(ruta string, db Database, f func(string, Database)) {
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

func ParsearPaciente(linea string, db Database) {
	dividido := strings.Split(linea, ",")
	nombre, año := dividido[0], dividido[1]

	año_en_int, err := strconv.Atoi(año)

	if err != nil {
		fmt.Printf(Constantes.ENOENT_ANIO, año)
		return
	}

	db.agregarPaciente(nombre, año_en_int)
}
