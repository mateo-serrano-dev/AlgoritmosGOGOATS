package parsing

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAGrafo "tdas/grafo"
	Constantes "tp3/constantes"
	DB "tp3/database"
	Models "tp3/models"
)

func ImportarPajek(ruta string, db DB.Database) TDAGrafo.Grafo[Models.Ciudad] {
	resultado := TDAGrafo.CrearGrafo[Models.Ciudad](false, true)

	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf(Constantes.NO_EXISTE_ARCHIVO, ruta)
		return nil
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)

	contador := proximoNumero(s)

	for contador > 0 {
		s.Scan()
		linea := s.Text()
		parsearVertice(resultado, linea, db)
		contador--
	}

	s.Scan() //Salteo el numero de aristas

	for s.Scan() {
		linea := s.Text()
		parsearArista(resultado, linea)
	}

	if err = s.Err(); err != nil {
		return nil
	}
}

func proximoNumero(s *bufio.Scanner) int {
	s.Scan()
	linea := s.Text()
	res, err := strconv.Atoi(linea)
	if err != nil {
		panic(Constantes.ARCHIVO_PAJEK_ERR)
	}

	return res
}

func parseFloat(s string) float64 {
	resultado, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(Constantes.ARCHIVO_PAJEK_ERR)
	}
	return resultado
}

func parsearVertice(g TDAGrafo.Grafo[Models.Ciudad], linea string, db DB.Database) {
	linea = strings.TrimSpace(linea)
	dividido := strings.Split(linea, ",")

	if len(dividido) != 3 {
		panic(Constantes.ARCHIVO_PAJEK_ERR)
	}
	nombre, latitud, longitud := dividido[0], dividido[1], dividido[2]
	lat := parseFloat(latitud)
	long := parseFloat(longitud)

	ciudad := Models.CrearCiudad(nombre, lat, long)
	g.AgregarVertice(ciudad)
}

func parsearArista(g TDAGrafo.Grafo[Models.Ciudad], linea string) {
	linea = strings.TrimSpace(linea)
	dividido := strings.Split(linea, ",")

	if len(dividido) != 3 {
		panic(Constantes.ARCHIVO_PAJEK_ERR)
	}
	ciudad1, ciudad2, distancia := dividido[0], dividido[1], dividido[2]
}
