package tp2

import (
	"bufio"
	"fmt"
	"os"
	Comando "tdas/comando"
	Constantes "tp2/constantes"
	Database "tp2/database"
	Parsing "tp2/parsing"
)

const _CANTIDAD_ARGUMENTOS = 3

func main() {
	db := Database.CrearDatabase()

	if len(os.Args) != _CANTIDAD_ARGUMENTOS {
		fmt.Printf(Constantes.ENOENT_CANT_PARAMS)
		return
	}

	csvPacientes, csvDoctores := os.Args[1], os.Args[2]

	Database.LeerArchivo(csvPacientes, db, Parsing.ParsearPaciente)
	Database.LeerArchivo(csvDoctores, db, Parsing.ParsearDoctor)

	comandos := Comando.ObtenerComandos(db)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		linea := s.Text()
		if linea == "" {
			break
		}

		comando, args, hayError := Parsing.ParsearComando(linea, comandos)
		if hayError {
			continue
		}
		comando.Ejecutar(args)
	}
	if err := s.Err(); err != nil {
		return
	}
}
