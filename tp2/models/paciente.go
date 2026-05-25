package models

import "tp2/constantes"

type Paciente struct {
	nombre          string
	año_inscripcion int
	urgencia        string
}

func CrearPaciente(nombre string, año int, urgencia string) *Paciente {
	return &Paciente{nombre, año, urgencia}
}

func (p *Paciente) ObtenerInscripcion() int {
	return p.año_inscripcion
}

func (p *Paciente) ObtenerAntiguedad() int {
	return constantes.ANIO_ACTUAL - p.año_inscripcion
}

func (p *Paciente) ObtenerUrgencia() string {
	return p.urgencia
}

func CompararPaciente(a *Paciente, b *Paciente) int {
	return a.ObtenerAntiguedad() - b.ObtenerAntiguedad()
}
