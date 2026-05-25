package service

import (
	"fmt"
	"tp2/constantes"
	"tp2/database"
)

func PedirTurno(db *database.Database, nombrePaciente, especialidad, urgencia string) {
	hayError := false

	if !db.ExistePaciente(nombrePaciente) {
		hayError = true
		fmt.Printf(constantes.ENOENT_PACIENTE, nombrePaciente)
	}
	if !db.ExisteEspecialidad(especialidad) {
		hayError = true
		fmt.Printf(constantes.ENOENT_ESPECIALIDAD, especialidad)
	}
	if urgencia != constantes.URGENTE && urgencia != constantes.REGULAR {
		hayError = true
		fmt.Printf(constantes.ENOENT_URGENCIA, urgencia)
	}

	if !hayError {
		db.EncolarTurno(nombrePaciente, urgencia, especialidad)
		fmt.Printf(constantes.PACIENTE_ENCOLADO, nombrePaciente)
		fmt.Printf(constantes.CANT_PACIENTES_ENCOLADOS, db.ObtenerCantidadEnEspera(especialidad), especialidad)
	}
}
