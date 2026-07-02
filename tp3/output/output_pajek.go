package output

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	TDAGrafo "tdas/grafo"
	Constantes "tp3/constantes"
	Models "tp3/models"
)

func ExportarPajek(ruta string, g TDAGrafo.GrafoPesado[Models.Ciudad]) {
	archivo, err := os.Create(ruta)
	if err != nil {
		fmt.Printf(Constantes.ERR_EXPORTAR)
		return
	}
	defer archivo.Close()

	vertices := g.CantidadVertices()
	aristas := 0
	lineas := make([]string, 0)

	dataWriter := bufio.NewWriter(archivo)
	cantidadVertices := fmt.Sprintf("%s\n", strconv.Itoa(vertices))
	dataWriter.WriteString(cantidadVertices)

	for iter := g.IterVertices(); iter.HayAlgoMas(); iter.Avanzar() {
		ciudad := iter.VerActual()
		linea := fmt.Sprintf("%s,%f,%f\n", ciudad.Nombre(), ciudad.Latitud(), ciudad.Longitud())

		_, err := dataWriter.WriteString(linea)
		if err != nil {
			fmt.Printf(Constantes.ERR_EXPORTAR)
			return
		}

		for subIter := g.IterAdyacentes(ciudad); subIter.HayAlgoMas(); subIter.Avanzar() {
			destino := subIter.VerActual()
			if ciudad.Nombre() < destino.Nombre() {
				arista := fmt.Sprintf("%s,%s,%d\n", ciudad.Nombre(), destino.Nombre(), int(g.Peso(ciudad, destino)))
				lineas = append(lineas, arista)
				aristas++
			}
		}
	}

	cantidadAristas := fmt.Sprintf("%s\n", strconv.Itoa(aristas))
	_, err = dataWriter.WriteString(cantidadAristas)
	if err != nil {
		fmt.Printf(Constantes.ERR_EXPORTAR)
		return
	}

	for _, linea := range lineas {
		_, err = dataWriter.WriteString(linea)
		if err != nil {
			fmt.Printf(Constantes.ERR_EXPORTAR)
			return
		}
	}

	dataWriter.Flush()
}
