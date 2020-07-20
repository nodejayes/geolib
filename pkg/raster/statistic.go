package raster

import (
	"math"
	"sort"
)

type Statistic struct {
	Mean      float64         `bson:"mean"`
	Median    float64         `bson:"median"`
	Cells     int             `bson:"cells"`
	Sum       float64         `bson:"sum"`
	Min       float64         `bson:"min"`
	Max       float64         `bson:"max"`
	Most      float64         `bson:"most"`
	Histogram map[float64]int `bson:"histogram"`
}

func calculateBandStatistic(band *Band, precision int) Statistic {
	var values []float64
	var mostValue float64

	min := math.MaxFloat64
	max := -math.MaxFloat64
	median := 0.0
	sum := 0.0
	cells := 0
	most := make(map[float64]int)
	hist := make(map[float64]int)

	generateBasicStats(band, &min, &max, &sum, &cells, &values, &most)
	generateMost(&mostValue, &most, &hist, precision)
	if cells < 1 {
		return Statistic{
			Mean:      0,
			Median:    0,
			Cells:     cells,
			Sum:       sum,
			Min:       0,
			Max:       0,
			Most:      mostValue,
			Histogram: hist,
		}
	}
	generateMedian(&median, &cells, &values)

	return Statistic{
		Min:       min,
		Max:       max,
		Sum:       sum,
		Cells:     cells,
		Mean:      sum / float64(cells),
		Median:    median,
		Most:      mostValue,
		Histogram: hist,
	}
}

func generateBasicStats(band *Band, min, max, sum *float64, cells *int, values *[]float64, most *map[float64]int) {
	for _, row := range band.Data {
		for _, value := range row {
			if value == band.NoData {
				continue
			}
			if value < *min {
				*min = value
			}
			if value > *max {
				*max = value
			}
			*sum = *sum + value
			*cells = *cells + 1
			*values = append(*values, value)
			(*most)[value] = (*most)[value] + 1
		}
	}
}

func generateMedian(median *float64, cells *int, values *[]float64) {
	sort.Float64s(*values)
	mid := (*cells - 1) / 2
	if mid%2 != 0 {
		*median = (*values)[mid]
	} else {
		*median = ((*values)[mid-1] + (*values)[mid]) / 2
	}
}

func generateMost(mostValue *float64, most *map[float64]int, hist *map[float64]int, precision int) {
	mostCounter := 0
	mult := 1
	for v := 0; v < precision; v++ {
		mult = mult * 10
	}
	for key, value := range *most {
		rounded := float64(int(key*float64(mult))) / float64(mult)
		(*hist)[rounded] = (*hist)[rounded] + 1
		if value > mostCounter {
			*mostValue = key
			mostCounter = value
		}
	}
}
