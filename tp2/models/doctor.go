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

func (d *Doctor) AumentarAtendidos() {
	d.atendidos++
}

func (d *Doctor) CantidadAtendidos() int {
	return d.atendidos
}

func (d *Doctor) ObtenerEspecialidad() string {
	return d.especialidad
}

func (d *Doctor) ObtenerNombre() string {
	return d.nombre
}
