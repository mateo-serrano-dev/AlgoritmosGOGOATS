package comando

import (
	DB "tp2/database"
)

type ComandoBase struct {
	database *DB.Database
}

type PedirTurno struct{ ComandoBase }
type AtenderPaciente struct{ ComandoBase }
type InformeDoctores struct{ ComandoBase }

func (c PedirTurno) Ejecutar() {

}

func (c AtenderPaciente) Ejecutar() {

}

func (c InformeDoctores) Ejecutar() {

}
