package logica

import (
	Cola "tdas/cola"
	ColaPrioridad "tdas/cola_prioridad"
)

type Especialidad struct {
	urgentes    Cola.Cola[*Paciente]
	no_urgentes ColaPrioridad.ColaPrioridad[*Paciente]
}

func CrearEspecialidad() *Especialidad {
	e := new(Especialidad)
	e.urgentes = Cola.CrearColaEnlazada[*Paciente]()
	e.no_urgentes = ColaPrioridad.CrearHeap(CompararPaciente)

	return e
}

func (e *Especialidad) DesencolarPaciente() *Paciente {
	var p *Paciente
	if !e.urgentes.EstaVacia() {
		p = e.urgentes.Desencolar()
	}

	if !e.no_urgentes.EstaVacia() {
		p = e.no_urgentes.Desencolar()
	}

	if p == nil {
		//Mensaje de error (TO DO)
	}

	return p
}

func (e *Especialidad) EnconlarPaciente(paciente *Paciente) {
	if paciente.urgencia == URGENTE {
		e.urgentes.Encolar(paciente)
	} else {
		e.no_urgentes.Encolar(paciente)
	}
}
