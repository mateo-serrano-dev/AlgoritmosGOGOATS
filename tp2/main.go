package tp2

import (
	"bufio"
	"fmt"
	"os"
	Constantes "tp2/constantes"
	Database "tp2/database"
)

const CANTIDAD_ARGUMENTOS = 3

func main() {
	db := Database.CrearDatabase()

	if len(os.Args) != CANTIDAD_ARGUMENTOS {
		fmt.Printf(Constantes.ENOENT_CANT_PARAMS)
		return
	}

	csvPacientes, csvDoctores := os.Args[1], os.Args[2]

	Database.LeerArchivo(csvPacientes, db, Database.ParsearPaciente)
	Database.LeerArchivo(csvDoctores, db, Database.ParsearDoctor)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
	}
}
