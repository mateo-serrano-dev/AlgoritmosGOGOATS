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

func ImportarPajek(ruta string, db DB.Database) {
	resultado := TDAGrafo.CrearGrafoPesado[Models.Ciudad](false)

	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf(Constantes.NO_EXISTE_ARCHIVO, ruta)
		return
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
		parsearArista(resultado, linea, db)
	}

	if err = s.Err(); err != nil {
		panic(Constantes.ARCHIVO_PAJEK_ERR)
	}

	db.CargarGrafo(resultado)
}

func ExportarPajek(ruta string, g TDAGrafo.GrafoPesado[Models.Ciudad]) {
	archivo, err := os.Create(ruta)
	if err != nil {
		fmt.Printf(Constantes.ERR_EXPORTAR)
	}
	defer archivo.Close()

	vertices := g.CantidadVertices()
	aristas := 0
	lineas := make([]string, vertices)

	dataWriter := bufio.NewWriter(archivo)
	dataWriter.WriteString(strconv.Itoa(vertices))

	for iter := g.IterVertices(); iter.HayAlgoMas(); iter.Avanzar() {
		ciudad := iter.VerActual()
		linea := fmt.Sprintf("%s,%f,%f\n", ciudad.Nombre(), ciudad.Latitud(), ciudad.Longitud())

		_, err := dataWriter.WriteString(linea)
		if err != nil {
			fmt.Printf(Constantes.ERR_EXPORTAR)
		}

		aristas += g.CantidadAristas(ciudad)
		for subIter := g.IterAdyacentes(ciudad); iter.HayAlgoMas(); iter.VerActual() {
			destino := subIter.VerActual()
			arista := fmt.Sprintf("%s,%s,%d\n", ciudad.Nombre(), destino.Nombre(), g.Peso(ciudad, destino))
			lineas = append(lineas, arista)
		}
	}

	_, err = dataWriter.WriteString(strconv.Itoa(aristas))
	if err != nil {
		fmt.Printf(Constantes.ERR_EXPORTAR)
	}

	for _, linea := range lineas {
		_, err = dataWriter.WriteString(linea)
		if err != nil {
			fmt.Printf(Constantes.ERR_EXPORTAR)
		}
	}

	dataWriter.Flush()
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

func parsearInt(s string) int {
	resultado, err := strconv.Atoi(s)
	if err != nil {
		panic(Constantes.ARCHIVO_PAJEK_ERR)
	}
	return resultado
}

func parsearFloat(s string) float64 {
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
	lat := parsearFloat(latitud)
	long := parsearFloat(longitud)

	ciudad := Models.CrearCiudad(nombre, lat, long)
	g.AgregarVertice(ciudad)

	db.RegistrarCiudad(nombre, ciudad)
}

func parsearArista(g TDAGrafo.GrafoPesado[Models.Ciudad], linea string, db DB.Database) {
	linea = strings.TrimSpace(linea)
	dividido := strings.Split(linea, ",")

	if len(dividido) != 3 {
		panic(Constantes.ARCHIVO_PAJEK_ERR)
	}
	ciudad1, ciudad2, distancia := dividido[0], dividido[1], dividido[2]
	g.AgregarArista(db.ObtenerCiudad(ciudad1), db.ObtenerCiudad(ciudad2), parsearInt(distancia))
}
