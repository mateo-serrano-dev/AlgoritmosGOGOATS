package output

import (
	"strings"
	Models "tp3/models"
)

func ConstruirMensajeSalida(camino []Models.Ciudad) string {
	ciudades := make([]string, 0)
	for _, ciudad := range camino {
		ciudades = append(ciudades, ciudad.Nombre())
	}
	return strings.Join(ciudades, " -> ")
}
