package proj4

var projections = make(map[int]string)

func setupProjections() {
	projections[4326] = "EPSG:4326"
	projections[3857] = "EPSG:3857"
}

func GetProjection(epsgCode int) string {
	if len(projections) < 1 {
		setupProjections()
	}
	return projections[epsgCode]
}

func RegisterProjection(epsgCode int, definition string) {
	projections[epsgCode] = definition
}
