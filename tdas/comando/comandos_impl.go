package comando

import (
	"fmt"
	Dict "tdas/diccionario"
	Constantes "tp2/constantes"
	DB "tp2/database"
	Service "tp2/service"
)

const _PERDIR_TURNO = "PERDIR_TURNO"
const _ATENDER_SIGUIENTE = "ATENDER_SIGUIENTE"
const _INFORME_DOCTOR = "INFORME"

const _PERDIR_TURNO_ARGS = 3
const _ATENDER_SIGUIENTE_ARGS = 1
const _INFORME_DOCTOR_ARGS = 2

type ComandoBase struct {
	database *DB.Database
}

type PedirTurno struct{ ComandoBase }
type AtenderPaciente struct{ ComandoBase }
type InformeDoctores struct{ ComandoBase }

func (c PedirTurno) Ejecutar(args []string) {
	if len(args) != _PERDIR_TURNO_ARGS {
		fmt.Printf(Constantes.ENOENT_PARAMS, _PERDIR_TURNO)
	}

	nombre, especialidad, urgencia := args[0], args[1], args[2]
	Service.PedirTurno(c.database, nombre, especialidad, urgencia)
}

func (c AtenderPaciente) Ejecutar(args []string) {
	if len(args) != _ATENDER_SIGUIENTE_ARGS {
		fmt.Printf(Constantes.ENOENT_PARAMS, _ATENDER_SIGUIENTE)
	}

	nombre := args[0]
	Service.AtenderSiguiente(nombre)
}

func (c InformeDoctores) Ejecutar(args []string) {
	if len(args) != _INFORME_DOCTOR_ARGS {
		fmt.Printf(Constantes.ENOENT_PARAMS, _INFORME_DOCTOR)
	}

	args_ptr := make([]*string, len(args))
	for i, arg := range args {
		if arg == "" {
			args_ptr[i] = nil
		} else {
			args_ptr[i] = &arg
		}
	}

	inicio, fin := args_ptr[0], args_ptr[1]
	Service.InformeDoctores(inicio, fin)
}

func ObtenerComandos() Dict.Diccionario[string, Comando] {
	dict := Dict.CrearHash[string, Comando]()
	dict.Guardar(_PERDIR_TURNO, new(PedirTurno))
	dict.Guardar(_ATENDER_SIGUIENTE, new(AtenderPaciente))
	dict.Guardar(_INFORME_DOCTOR, new(InformeDoctores))
	return dict
}
