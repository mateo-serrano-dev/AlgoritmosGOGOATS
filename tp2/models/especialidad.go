package models

import (
	"fmt"
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
	var p *Paciente
	if !e.urgentes.EstaVacia() {
		p = e.urgentes.Desencolar()
	}

	if !e.noUrgentes.EstaVacia() {
		p = e.noUrgentes.Desencolar()
	}

	if p == nil {
		fmt.Printf(constantes.SIN_PACIENTES)
	}

	return p
}

func (e *Especialidad) EnconlarPaciente(paciente *Paciente, urgencia string) {
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
