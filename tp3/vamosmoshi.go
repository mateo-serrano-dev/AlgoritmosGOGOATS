package main

import (
	"bufio"
	"fmt"
	"os"
	Comando "tdas/comando"
	Constantes "tp3/constantes"
	DB "tp3/database"
	Parsing "tp3/parsing"
)

const _CANTIDAD_ARGUMENTOS = 2

func main() {
	db := DB.CrearDatabase()

	if len(os.Args) != _CANTIDAD_ARGUMENTOS {
		fmt.Printf(Constantes.ERR_ARGS)
		return
	}

	ciudadesPajek := os.Args[1]
	Parsing.ImportarPajek(ciudadesPajek, db)

	comandos := Comando.ObtenerComandos(db)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		linea := s.Text()
		if linea == "" {
			break
		}

		comando, args, hayError := Parsing.ParsearComando(linea, comandos)
		if hayError {
			fmt.Print(Constantes.ERR_CMD)
			continue
		}
		comando.Ejecutar(args)
	}
	if err := s.Err(); err != nil {
		return
	}
}
