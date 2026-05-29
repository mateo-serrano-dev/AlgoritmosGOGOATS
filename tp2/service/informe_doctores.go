package service

import (
	"fmt"
	"tp2/constantes"
	"tp2/database"
)

func InformeDoctores(db *database.Database, inicio, fin *string) {
	sliceInfoDoctores := db.ObtenerDoctoresPorRango(inicio, fin)
	fmt.Printf(constantes.DOCTORES_SISTEMA, len(sliceInfoDoctores))
	for i, doctor := range sliceInfoDoctores {
		fmt.Printf(
			constantes.INFORME_DOCTOR,
			i+1,
			doctor.ObtenerNombre(),
			doctor.ObtenerEspecialidad(),
			doctor.CantidadAtendidos(),
		)
	}
}
