package comando

import (
	"fmt"
	Dict "tdas/diccionario"
	Constantes "tp2/constantes"
	DB "tp2/database"
	Service "tp2/service"
)

type ComandoBase struct {
	database *DB.Database
}

type PedirTurno struct{ ComandoBase }
type AtenderPaciente struct{ ComandoBase }
type InformeDoctores struct{ ComandoBase }

func (c *PedirTurno) Ejecutar(args []string) {
	if !c.cantidadArgumentosCorrecta(args) {
		return
	}

	nombre, especialidad, urgencia := args[0], args[1], args[2]
	Service.PedirTurno(c.database, nombre, especialidad, urgencia)
}

func (c *AtenderPaciente) Ejecutar(args []string) {
	if !c.cantidadArgumentosCorrecta(args) {
		return
	}

	nombre := args[0]
	Service.AtenderSiguiente(nombre)
}

func (c *InformeDoctores) Ejecutar(args []string) {
	if !c.cantidadArgumentosCorrecta(args) {
		return
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

func verificarArgumentos(len_args, esperado int, nombre string) bool {
	result := len_args == esperado
	if result {
		fmt.Printf(Constantes.ENOENT_PARAMS, nombre)
	}
	return result
}

// cantidadArgumentosCorrecta es un método privado que solo se usa internamente para validar rapidamente los args ingresados
func (c *PedirTurno) cantidadArgumentosCorrecta(args []string) bool {
	return verificarArgumentos(len(args), Constantes.PEDIR_TURNO_ARGS, Constantes.CMD_PEDIR_TURNO)
}

func (c *AtenderPaciente) cantidadArgumentosCorrecta(args []string) bool {
	return verificarArgumentos(len(args), Constantes.ATENDER_SIGUIENTE_ARGS, Constantes.CMD_ATENDER_SIGUIENTE)
}

func (c *InformeDoctores) cantidadArgumentosCorrecta(args []string) bool {
	return verificarArgumentos(len(args), Constantes.INFORME_DOCTOR_ARGS, Constantes.CMD_INFORME_DOCTOR)
}

// Devuelve un diccionario con los comandos posibles en el CLI
func ObtenerComandos(db *DB.Database) Dict.Diccionario[string, Comando] {
	dict := Dict.CrearHash[string, Comando]()
	dict.Guardar(Constantes.CMD_PEDIR_TURNO, new(PedirTurno))
	dict.Guardar(Constantes.CMD_ATENDER_SIGUIENTE, new(AtenderPaciente))
	dict.Guardar(Constantes.CMD_INFORME_DOCTOR, new(InformeDoctores))
	return dict
}
