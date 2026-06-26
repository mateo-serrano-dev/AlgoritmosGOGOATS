package comando

import (
	"fmt"
	Dict "tdas/diccionario"
	Constantes "tp3/constantes"
	DB "tp3/database"
	Service "tp3/service"
)

type ComandoBase struct {
	database *DB.Database
}

type IrCmd struct{ ComandoBase }
type ItinerarioCmd struct{ ComandoBase }
type ViajeCmd struct{ ComandoBase }
type ReducirCaminosCmd struct{ ComandoBase }

func (c *IrCmd) Ejecutar(args []string) {
	if !c.cantidadArgumentosCorrecta(args) {
		return
	}

	ciudad1 := c.database.ObtenerCiudad(args[0])
	ciudad2 := c.database.ObtenerCiudad(args[1])
	ruta := args[2]
	Service.IrCmd(c.database, ciudad1, ciudad2, ruta)
}

func (c *ItinerarioCmd) Ejecutar(args []string) {
	if !c.cantidadArgumentosCorrecta(args) {
		return
	}

	ruta := args[0]
	Service.ItinerarioCmd(c.database, ruta)
}

func (c *ViajeCmd) Ejecutar(args []string) {
	if !c.cantidadArgumentosCorrecta(args) {
		return
	}

	origen := c.database.ObtenerCiudad(args[0])
	Service.ViajeCmd(c.database, origen, args[1])
}

func (c *ReducirCaminosCmd) Ejecutar(args []string) {
	if !c.cantidadArgumentosCorrecta(args) {
		return
	}

	Service.ReducirCaminosCmd(c.database, args[0])
}

func verificarArgumentos(len_args, esperado int) bool {
	result := len_args == esperado
	if !result {
		fmt.Printf(Constantes.ERR_ARGS)
	}
	return result
}

func (c *IrCmd) cantidadArgumentosCorrecta(args []string) bool {
	return verificarArgumentos(len(args), Constantes.ARGS_IR)
}

func (c *ItinerarioCmd) cantidadArgumentosCorrecta(args []string) bool {
	return verificarArgumentos(len(args), Constantes.ARGS_ITINERARIO)
}

func (c *ViajeCmd) cantidadArgumentosCorrecta(args []string) bool {
	return verificarArgumentos(len(args), Constantes.ARGS_VIAJE)
}

func (c *ReducirCaminosCmd) cantidadArgumentosCorrecta(args []string) bool {
	return verificarArgumentos(len(args), Constantes.ARGS_REDUCIR)
}

// Devuelve un diccionario con los comandos posibles en el CLI
func ObtenerComandos(db *DB.Database) Dict.Diccionario[string, Comando] {
	dict := Dict.CrearHash[string, Comando]()
	comandoBase := &ComandoBase{database: db}
	dict.Guardar(Constantes.CMD_IR, &IrCmd{*comandoBase})
	dict.Guardar(Constantes.CMD_ITINERARIO, &ItinerarioCmd{*comandoBase})
	dict.Guardar(Constantes.CMD_VIAJE, &ViajeCmd{*comandoBase})
	dict.Guardar(Constantes.CMD_REDUCIR, &ReducirCaminosCmd{*comandoBase})
	return dict
}
