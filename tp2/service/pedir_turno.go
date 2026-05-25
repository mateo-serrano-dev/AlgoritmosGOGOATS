package service

import (
	"fmt"
	"tp2/constantes"
	"tp2/database"
)

func PedirTurno(db *database.Database, nombrePaciente, especialidad, urgencia string, año int) {
	error := false

	if !db.ExistePaciente(nombrePaciente) {
		error = true
		fmt.Printf(constantes.ENOENT_PACIENTE, nombrePaciente)
	}
	if !db.ExisteEspecialidad(especialidad) {
		error = true
		fmt.Printf(constantes.ENOENT_ESPECIALIDAD, especialidad)
	}
	if urgencia != constantes.URGENTE && urgencia != constantes.REGULAR {
		error = true
		fmt.Printf(constantes.ENOENT_URGENCIA, urgencia)
	}

	if !error {
		db.EncolarTurno(nombrePaciente, año, urgencia, especialidad)
	}
}
