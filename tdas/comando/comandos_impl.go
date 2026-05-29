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
	Service.AtenderSiguiente(c.database, nombre)
}

func (c *InformeDoctores) Ejecutar(args []string) {
	if !c.cantidadArgumentosCorrecta(args) {
		return
	}

	argsPuntero := make([]*string, len(args))
	for i := range args {
		if args[i] == "" {
			argsPuntero[i] = nil
		} else {
			argsPuntero[i] = &args[i]
		}
	}

	inicio, fin := argsPuntero[0], argsPuntero[1]
	Service.InformeDoctores(c.database, inicio, fin)
}

func verificarArgumentos(len_args, esperado int, nombre string) bool {
	result := len_args == esperado
	if !result {
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
	comandoBase := &ComandoBase{database: db}
	dict.Guardar(Constantes.CMD_PEDIR_TURNO, &PedirTurno{*comandoBase})
	dict.Guardar(Constantes.CMD_ATENDER_SIGUIENTE, &AtenderPaciente{*comandoBase})
	dict.Guardar(Constantes.CMD_INFORME_DOCTOR, &InformeDoctores{*comandoBase})
	return dict
}
