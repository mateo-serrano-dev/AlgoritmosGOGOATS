package tp3

const NO_EXISTE_ARCHIVO = "No existe el archivo %s\n"
const ARCHIVO_PAJEK_ERR = "Mal formato de pajek\n"
const ERR_EXPORTAR = "Hubo un error al exportar\n"

const ERR_ARGS = "Cantidad de parámetros incorrecta\n"
const ARGS_VIAJE = 2
const ARGS_ITINERARIO = 1
const ARGS_IR = 3
const ARGS_REDUCIR = 1

const CMD_IR = "ir"
const CMD_VIAJE = "viaje"
const CMD_ITINERARIO = "itinerario"
const CMD_REDUCIR = "reducir_caminos"

const ERR_RECORRIDO_NO_ENCONTRADO = "No se encontro recorrido"

const KML_HEADER = `<?xml version="1.0" encoding="UTF-8"?>
<kml xmlns="http://earth.google.com/kml/2.1">
    <Document>
        <name>Camino Mínimo</name>
        <description>Recorrido generado por el algoritmo de Dijkstra.</description>
`
const KML_PLACEMARK = `        <Placemark>
            <name>%s</name>
            <Point>
                <coordinates>%s</coordinates>
            </Point>
        </Placemark>
`
const KML_LINESTRING = `        <Placemark>
            <LineString>
                <coordinates>%s</coordinates>
            </LineString>
        </Placemark>
    </Document>
</kml>
`

const CODIGO_DE_PERMISO = 0644

const TIEMPO_TOTAL = "Tiempo total: %d\n"
