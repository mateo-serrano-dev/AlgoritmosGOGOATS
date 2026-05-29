package service

import (
	"fmt"
	"tp2/constantes"
	"tp2/database"
)

func AtenderSiguiente(db *database.Database, nombreDoctor string) {
	if !db.ExisteDoctor(nombreDoctor) {
		fmt.Printf(constantes.ENOENT_DOCTOR, nombreDoctor)
		return
	}

	doctor := db.ObtenerDoctor(nombreDoctor)
	especialidad := db.ObtenerEspecialidad(doctor.ObtenerEspecialidad())

	paciente := especialidad.DesencolarPaciente()
	if paciente == nil {
		fmt.Printf(constantes.SIN_PACIENTES)
		return
	}

	doctor.AumentarAtendidos()

	fmt.Printf(constantes.PACIENTE_ATENDIDO, paciente.ObtenerNombre())
	fmt.Printf(constantes.CANT_PACIENTES_ENCOLADOS, especialidad.CantidadEnEspera(), doctor.ObtenerEspecialidad())
}
