package database

import (
	Dict "tdas/diccionario"
	Logica "tp2/logica"
)

type Database struct {
	pacientes Dict.Diccionario[string, *Logica.Paciente]
	//doctores  Dict.DiccionarioOrdenado[string, *Doctor]
	especialidades Dict.Diccionario[string, *Logica.Especialidad]
}

func (db *Database) AgregarPaciente(nombre string, año int, urgencia string) *Logica.Paciente {
	p := Logica.CrearPaciente(nombre, año, urgencia)
	db.pacientes.Guardar(nombre, p)
	return p
}

func (db *Database) AgregarEspecialidad(nombre string) *Logica.Especialidad {
	e := Logica.CrearEspecialidad()
	db.especialidades.Guardar(nombre, e)
	return e
}

func (db *Database) EncolarTurno(nombre string, año int, urgencia string, especialidad string) {
	var p *Logica.Paciente
	if !db.pacientes.Pertenece(nombre) {
		p = db.AgregarPaciente(nombre, año, urgencia)
	} else {
		p = db.pacientes.Obtener(nombre)
	}

	var e *Logica.Especialidad
	if !db.especialidades.Pertenece(especialidad) {
		e = db.AgregarEspecialidad(especialidad)
	} else {
		e = db.especialidades.Obtener(especialidad)
	}

	e.EnconlarPaciente(p)
}
