package output

import (
	"fmt"
	"os"
	"strings"
	TDADict "tdas/diccionario"
	Constantes "tp3/constantes"
	Modelos "tp3/models"
)

func obtenerCoordsString(ciudad Modelos.Ciudad) string {
	return fmt.Sprintf("%f, %f", ciudad.Longitud(), ciudad.Latitud())
}

func ExportarKml(camino []Modelos.Ciudad, ruta string) error {
	var archivo strings.Builder
	archivo.WriteString(Constantes.KML_HEADER)

	visitados := TDADict.CrearHash[Modelos.Ciudad, bool]()
	coordenadas := make([]string, len(camino))

	for i, ciudad := range camino {
		coordStr := obtenerCoordsString(ciudad)
		coordenadas[i] = coordStr

		if !visitados.Pertenece(ciudad) {
			placemark := fmt.Sprintf(Constantes.KML_PLACEMARK, ciudad.Nombre(), coordStr)
			archivo.WriteString(placemark)
			visitados.Guardar(ciudad, true)
		}
	}

	for i := 1; i < len(coordenadas); i++ {
		linea := fmt.Sprintf("%s %s", coordenadas[i-1], coordenadas[i])
		linestring := fmt.Sprintf(Constantes.KML_LINESTRING, linea)
		archivo.WriteString(linestring)
	}
	archivo.WriteString(Constantes.KML_FOOTER)

	return os.WriteFile(ruta, []byte(archivo.String()), Constantes.CODIGO_DE_PERMISO)
}
