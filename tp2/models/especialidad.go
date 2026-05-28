package models

import (
	Cola "tdas/cola"
	ColaPrioridad "tdas/cola_prioridad"
	"tp2/constantes"
)

type Especialidad struct {
	urgentes   Cola.Cola[*Paciente]
	noUrgentes ColaPrioridad.ColaPrioridad[*Paciente]
	cantidad   int
}

func CrearEspecialidad() *Especialidad {
	return &Especialidad{
		urgentes:   Cola.CrearColaEnlazada[*Paciente](),
		noUrgentes: ColaPrioridad.CrearHeap(CompararPaciente),
	}
}

func (e *Especialidad) DesencolarPaciente() *Paciente {
	// puede devolver nil, se maneja ese caso especificamente donde corresponda
	var p *Paciente
	if !e.urgentes.EstaVacia() {
		p = e.urgentes.Desencolar()
	}

	if !e.noUrgentes.EstaVacia() {
		p = e.noUrgentes.Desencolar()
	}

	return p
}

func (e *Especialidad) EncolarPaciente(paciente *Paciente, urgencia string) {
	if urgencia == constantes.URGENTE {
		e.urgentes.Encolar(paciente)
	} else {
		e.noUrgentes.Encolar(paciente)
	}
	e.cantidad++
}

func (e *Especialidad) CantidadEnEspera() int {
	return e.cantidad
}
