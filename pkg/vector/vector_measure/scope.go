package vector_measure

// Scope measure the Length of all 3 Edges from a Triangle
func Scope(t I3Facer) (float64, error) {
	lAB, err := EdgeLength(t, EdgeAB)
	lAC, err := EdgeLength(t, EdgeAC)
	lCB, err := EdgeLength(t, EdgeCB)
	if err != nil {
		return 0, err
	}
	return lAB + lAC + lCB, nil
}
