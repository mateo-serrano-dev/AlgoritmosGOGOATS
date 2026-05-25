package models

type Doctor struct {
	nombre       string
	especialidad string
	atendidos    int
}

func CrearDoctor(nombre, especialidad string) *Doctor {
	return &Doctor{
		nombre:       nombre,
		especialidad: especialidad,
		atendidos:    0,
	}
}

func (d *Doctor) AtenderPaciente() {
	d.atendidos++
}

func (d *Doctor) CantidadAtendidos() int {
	return d.atendidos
}
