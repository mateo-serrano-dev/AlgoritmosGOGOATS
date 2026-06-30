package output

import (
	"fmt"
	"os"
	"strings"
	Constantes "tp3/constantes"
	Models "tp3/models"
)

func ExportarKml(camino []Models.Ciudad, ruta string) error {
	var archivo strings.Builder

	archivo.WriteString(Constantes.KML_HEADER)

	coordenadas := make([]string, 0)
	for _, ciudad := range camino {
		coordStr := fmt.Sprintf("%f, %f", ciudad.Longitud(), ciudad.Latitud())
		coordenadas = append(coordenadas, coordStr)
		placemark := fmt.Sprintf(Constantes.KML_PLACEMARK, ciudad.Nombre(), coordStr)
		archivo.WriteString(placemark)
	}
	coords := strings.Join(coordenadas, " ")
	linestring := fmt.Sprintf(Constantes.KML_LINESTRING, coords)
	archivo.WriteString(linestring)
	return os.WriteFile(ruta, []byte(archivo.String()), Constantes.CODIGO_DE_PERMISO)
}
