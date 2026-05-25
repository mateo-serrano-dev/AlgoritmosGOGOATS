package database

import (
	"strings"
	Dict "tdas/diccionario"
	Modelos "tp2/models"
)

type Database struct {
	pacientes      Dict.Diccionario[string, *Modelos.Paciente]
	doctores       Dict.DiccionarioOrdenado[string, *Modelos.Doctor]
	especialidades Dict.Diccionario[string, *Modelos.Especialidad]
}

func CrearDatabase() *Database {
	db := new(Database)
	db.pacientes = Dict.CrearHash[string, *Modelos.Paciente]()
	db.doctores = Dict.CrearABB[string, *Modelos.Doctor](strings.Compare)
	db.especialidades = Dict.CrearHash[string, *Modelos.Especialidad]()
	return db
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

func (db *Database) agregarDoctor(nombre string, especialidad string) *Modelos.Doctor {
	d := Modelos.CrearDoctor(nombre, especialidad)
	db.doctores.Guardar(nombre, d)
	return d
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
