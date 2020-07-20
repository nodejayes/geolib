package raster

// the Band Data of a Raster Band
type Band struct {
	BandHeader     uint8
	NoData         float64
	IsOffline      bool
	HasNoDataValue bool
	IsNoDataValue  bool
	Data           [][]float64
}

// fill the Band Data from a 1D Array
func (rb *Band) FillFrom1D(values []float64, width, height int) {
	for idx, value := range values {
		col := idx % height
		row := idx / height
		if col == 0 {
			rb.Data[row] = make([]float64, width)
		}
		rb.Data[row][col] = value
	}
}

// fill the Band Data from a 2D Array
func (rb *Band) Fill(data [][]float64) {
	rb.Data = data
}
