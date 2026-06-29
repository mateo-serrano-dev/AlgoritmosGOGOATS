package service

import (
	"grafos_util"
	DB "tp3/database"
	Modelos "tp3/models"
)

func IrCmd(db *DB.Database, desde, hasta Modelos.Ciudad, ruta string) {
	padres, _ := grafos_util.Dijkstra(db.ObtenerGrafo(), desde)
}
