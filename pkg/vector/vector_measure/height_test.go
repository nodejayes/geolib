package vector_measure

import (
	"fmt"
	"testing"
)

var HeightTest = []struct {
	name           string
	triangle       TestTriangle
	expectedHeight float64
	expectedErr    string
	pointKey       PointKey
}{
	{
		name: "Height for Point A",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		pointKey:       Alpha,
		expectedHeight: 3.130495169777232,
	},
	{
		name: "Height for Point B",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		pointKey:       Beta,
		expectedHeight: 1.9414506856750204,
	},
	{
		name: "Height for Point C",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		pointKey:       Gamma,
		expectedHeight: 1.3728129465925845,
	},
	{
		name: "invalid Point Key",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		pointKey:    "",
		expectedErr: "invalid Point Key  only support Alpha, Beta or Gamma",
	},
	{
		name: "Point A Negative",
		triangle: TestTriangle{
			A: []float64{
				386130,
				5692254,
			},
			B: []float64{
				386408.967053963,
				5692913.35517462,
			},
			C: []float64{
				386106.870525168,
				5692632.47844243,
			},
		},
		pointKey:       Alpha,
		expectedHeight: 292.9317229151599,
	},
	{
		name: "Point A Positive",
		triangle: TestTriangle{
			A: []float64{
				386130,
				5692254,
			},
			B: []float64{
				386106.870525168,
				5692632.47844243,
			},
			C: []float64{
				386211.636432768,
				5692238.19060928,
			},
		},
		pointKey:       Alpha,
		expectedHeight: 74.83891952375997,
	},
	{
		name: "Point A Negative1",
		triangle: TestTriangle{
			A: []float64{
				386130,
				5692254,
			},
			B: []float64{
				386211.636432768,
				5692238.19060928,
			},
			C: []float64{
				386259.794025127,
				5692052.3088027,
			},
		},
		pointKey:       Alpha,
		expectedHeight: 75.06238598507127,
	},
	{
		name: "Point A Negative2",
		triangle: TestTriangle{
			A: []float64{
				386130,
				5692254,
			},
			B: []float64{
				386211.636432768,
				5692238.19060928,
			},
			C: []float64{
				386259.794025127,
				5692052.3088027,
			},
		},
		pointKey:       Alpha,
		expectedHeight: 75.06238598507127,
	},
}

func TestHeightFrom(t *testing.T) {
	for _, tt := range HeightTest {
		t.Run(tt.name, func(t *testing.T) {
			val, err := HeightFrom(tt.triangle, tt.pointKey)
			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					if err == nil {
						t.Errorf("Expect err to be %v: nil", tt.expectedErr)
						return
					} else {
						t.Errorf("Expect err to be %v: %v", tt.expectedErr, err.Error())
						return
					}
				}
			} else {
				if err != nil {
					t.Errorf("Expect err to be nil: %v", err.Error())
					return
				}
			}
			if val != tt.expectedHeight {
				t.Errorf("Expect height to be %v: %v", tt.expectedHeight, val)
			}
		})
	}
}

func ExampleHeightFrom() {
	t := TestTriangle{
		A: []float64{0, 0},
		B: []float64{5, 1},
		C: []float64{3, 2},
	}
	val, _ := HeightFrom(t, Alpha)
	fmt.Printf("%v", val)
	// Output: 3.130495169777232
}

var ShortestDistanceTest = []struct {
	name           string
	triangle       TestTriangle
	expectedHeight float64
	expectedErr    string
	pointKey       PointKey
}{
	{
		name: "Point A Negative",
		triangle: TestTriangle{
			A: []float64{
				386130,
				5692254,
			},
			B: []float64{
				386408.967053963,
				5692913.35517462,
			},
			C: []float64{
				386106.870525168,
				5692632.47844243,
			},
		},
		pointKey:       Alpha,
		expectedHeight: 379.18452498739,
	},
	{
		name: "Point A Positive",
		triangle: TestTriangle{
			A: []float64{
				386130,
				5692254,
			},
			B: []float64{
				386106.870525168,
				5692632.47844243,
			},
			C: []float64{
				386211.636432768,
				5692238.19060928,
			},
		},
		pointKey:       Alpha,
		expectedHeight: 74.83891952375997,
	},
	{
		name: "Point A Negative1",
		triangle: TestTriangle{
			A: []float64{
				386130,
				5692254,
			},
			B: []float64{
				386211.636432768,
				5692238.19060928,
			},
			C: []float64{
				386259.794025127,
				5692052.3088027,
			},
		},
		pointKey:       Alpha,
		expectedHeight: 83.15313577986396,
	},
	{
		name: "Point A Negative2",
		triangle: TestTriangle{
			A: []float64{
				386130,
				5692254,
			},
			B: []float64{
				386211.636432768,
				5692238.19060928,
			},
			C: []float64{
				386259.794025127,
				5692052.3088027,
			},
		},
		pointKey:       Alpha,
		expectedHeight: 83.15313577986396,
	},
	{
		name: "invalid Point Key",
		triangle: TestTriangle{
			A: []float64{0, 0},
			B: []float64{5, 1},
			C: []float64{3, 2},
		},
		pointKey:    "",
		expectedErr: "invalid Point Key  only support Alpha, Beta or Gamma",
	},
}

func TestShortestDistance(t *testing.T) {
	for _, tt := range ShortestDistanceTest {
		t.Run(tt.name, func(t *testing.T) {
			val, err := ShortestDistance(tt.triangle, tt.pointKey)
			if tt.expectedErr != "" {
				if err == nil || err.Error() != tt.expectedErr {
					if err == nil {
						t.Errorf("Expect err to be %v: nil", tt.expectedErr)
						return
					} else {
						t.Errorf("Expect err to be %v: %v", tt.expectedErr, err.Error())
						return
					}
				}
			} else {
				if err != nil {
					t.Errorf("Expect err to be nil: %v", err.Error())
					return
				}
			}
			if val != tt.expectedHeight {
				t.Errorf("Expect height to be %v: %v", tt.expectedHeight, val)
			}
		})
	}
}

func ExampleShortestDistance() {
	t := TestTriangle{
		A: []float64{
			386130,
			5692254,
		},
		B: []float64{
			386408.967053963,
			5692913.35517462,
		},
		C: []float64{
			386106.870525168,
			5692632.47844243,
		},
	}
	val, _ := ShortestDistance(t, Alpha)
	fmt.Printf("%v", val)
	// Output: 379.18452498739
}
