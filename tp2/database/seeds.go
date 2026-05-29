package database

import (
	"bufio"
	"fmt"
	"os"
	Constantes "tp2/constantes"
)

func LeerArchivo(ruta string, db *Database, f func(string, *Database)) {
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf(Constantes.ENOENT_ARCHIVO, ruta)
		return
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		f(linea, db)
	}
	if err := s.Err(); err != nil {
		return
	}
}
