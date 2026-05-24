package logica

type Paciente struct {
	nombre          string
	año_inscripcion int
	urgencia        string
}

const AÑO_ACTUAL = 2026

const URGENTE = "URGENTE"
const REGULAR = "REGULAR"

func CrearPaciente(nombre string, año int, urgencia string) *Paciente {
	return &Paciente{nombre, año, urgencia}
}

func (p *Paciente) ObtenerInscripcion() int {
	return p.año_inscripcion
}

func (p *Paciente) ObtenerAntiguedad() int {
	return AÑO_ACTUAL - p.año_inscripcion
}

func (p *Paciente) ObtenerUrgencia() string {
	return p.urgencia
}

func CompararPaciente(a *Paciente, b *Paciente) int {
	return a.ObtenerAntiguedad() - b.ObtenerAntiguedad()
}
