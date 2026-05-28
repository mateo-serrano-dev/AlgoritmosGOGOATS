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

// ----Validar existencia en la DB----
func (db *Database) ExistePaciente(nombre string) bool {
	return db.pacientes.Pertenece(nombre)
}

func (db *Database) ExisteEspecialidad(especialidad string) bool {
	return db.especialidades.Pertenece(especialidad)
}

func (db *Database) ExisteDoctor(nombreDoctor string) bool {
	return db.doctores.Pertenece(nombreDoctor)
}

// ----Guardado de elementos en la DB
func (db *Database) AgregarPaciente(nombre string, año int) *Modelos.Paciente {
	p := Modelos.CrearPaciente(nombre, año)
	db.pacientes.Guardar(nombre, p)
	return p
}

func (db *Database) AgregarEspecialidad(nombre string) *Modelos.Especialidad {
	e := Modelos.CrearEspecialidad()
	db.especialidades.Guardar(nombre, e)
	return e
}

func (db *Database) AgregarDoctor(nombre string, especialidad string) *Modelos.Doctor {
	if !db.ExisteEspecialidad(especialidad) {
		db.AgregarEspecialidad(especialidad)
	}
	d := Modelos.CrearDoctor(nombre, especialidad)
	db.doctores.Guardar(nombre, d)
	return d
}

// ----Obtener elementos de la DB
func (db *Database) ObtenerDoctor(nombreDoctor string) *Modelos.Doctor {
	// precondicion: el doctor ya existe en la db
	return db.doctores.Obtener(nombreDoctor)
}

func (db *Database) ObtenerEspecialidad(nombreEspecialidad string) *Modelos.Especialidad {
	// precondicion: la especialidad ya existe en la db
	return db.especialidades.Obtener(nombreEspecialidad)
}

func (db *Database) ObtenerPaciente(nombrePaciente string) *Modelos.Paciente {
	// precondicion: el paciente ya existe en la db
	return db.pacientes.Obtener(nombrePaciente)
}
