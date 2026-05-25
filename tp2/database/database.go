package database

import (
	Dict "tdas/diccionario"
	Modelos "tp2/models"
)

type Database struct {
	pacientes Dict.Diccionario[string, *Modelos.Paciente]
	//doctores  Dict.DiccionarioOrdenado[string, *Doctor]
	especialidades Dict.Diccionario[string, *Modelos.Especialidad]
}

func (db *Database) ExistePaciente(nombre string) bool {
	return db.pacientes.Pertenece(nombre)
}

func (db *Database) ExisteEspecialidad(especialidad string) bool {
	return db.especialidades.Pertenece(especialidad)
}

func (db *Database) agregarPaciente(nombre string, año int) *Modelos.Paciente {
	p := Modelos.CrearPaciente(nombre, año)
	db.pacientes.Guardar(nombre, p)
	return p
}

func (db *Database) agregarEspecialidad(nombre string) *Modelos.Especialidad {
	e := Modelos.CrearEspecialidad()
	db.especialidades.Guardar(nombre, e)
	return e
}

func (db *Database) EncolarTurno(nombre string, año int, urgencia string, especialidad string) {
	// precondicion: el paciente y la especialidad ya existen en la db
	p := db.pacientes.Obtener(nombre)
	e := db.especialidades.Obtener(especialidad)

	e.EnconlarPaciente(p, urgencia)
}

func (db *Database) ObtenerCantidadEnEspera(especialidad string) int {
	// precondicion: la especialidad ya existe en la db
	e := db.especialidades.Obtener(especialidad)
	return e.CantidadEnEspera()
}
