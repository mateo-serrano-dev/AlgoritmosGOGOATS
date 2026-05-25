package models

import "tp2/constantes"

type Paciente struct {
	nombre          string
	año_inscripcion int
}

/*type Turno struct {
	paciente *Paciente
	urgencia string
}*/

func CrearPaciente(nombre string, año int) *Paciente {
	return &Paciente{nombre, año}
}

func (p *Paciente) ObtenerInscripcion() int {
	return p.año_inscripcion
}

func (p *Paciente) ObtenerAntiguedad() int {
	return constantes.ANIO_ACTUAL - p.año_inscripcion
}

/*func (t *Turno) ObtenerUrgencia() string {
	return t.urgencia
}*/

func CompararPaciente(a *Paciente, b *Paciente) int {
	return a.ObtenerAntiguedad() - b.ObtenerAntiguedad()
}
