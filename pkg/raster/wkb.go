package raster

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/nodejayes/geolib/pkg/geometry"
	"github.com/nodejayes/geolib/pkg/reference_system"
	"log"
	"math"
	"sync"
)

type Header struct {
	Version uint16
	Bands   uint16
	ScaleX  float64
	ScaleY  float64
	IpX     float64
	IpY     float64
	SkewX   float64
	SkewY   float64
	Srid    int32
	Width   uint16
	Height  uint16
}

type Raster struct {
	Endianness uint8
	Version    uint16
	ScaleX     float64
	ScaleY     float64
	IpX        float64
	MidIpX     float64
	IpY        float64
	MidIpY     float64
	SkewX      float64
	SkewY      float64
	Srid       int32
	Width      uint16
	Height     uint16
	Bands      []Band
	ValidTo    int64
	Workers    int
}

// adds new Band with Raster Data
func (r *Raster) AddBand(band Band) {
	r.Bands = append(r.Bands, band)
}

// get the Endianness ByteOrder
func (r *Raster) GetByteEndianness() binary.ByteOrder {
	if r.Endianness == 0 {
		return binary.BigEndian
	}
	return binary.LittleEndian
}

// convert the Raster to a Point Grid Feature Collection
func (r *Raster) AsFeatureCollection(bandNum int) *geometry.FeatureCollection {
	selectedBand := r.Bands[bandNum-1]
	targetLength := int(r.Height) * int(r.Width)
	xv := r.MidIpX
	yv := r.MidIpY
	fc := geometry.NewFeatureCollection(make([]*geometry.Feature, targetLength))

	for i := 0; i < targetLength; i++ {
		rowId := i / int(r.Width)
		colId := i % int(r.Width)
		xv = r.MidIpX + (3.0 * float64(colId))
		if colId == 0 {
			yv = r.MidIpY - (3.0 * float64(rowId))
		}

		v := selectedBand.Data[rowId][colId]
		idx := colId*int(r.Height) + rowId

		fc.Features[idx] = geometry.NewFeature(geometry.NewPoint([]float64{
			xv, yv,
		}, reference_system.New(int(r.Srid))), map[string]interface{}{
			"value": v,
		})
	}

	return fc
}

// convert the Raster to a byte Hex String to insert into Postgres with the
// st_rastfromhexwkb function
func (r *Raster) ToHex() []byte {
	byteRaster := r.ToByte()
	hexRaster := make([]byte, hex.EncodedLen(len(byteRaster)))
	hex.Encode(hexRaster, byteRaster)
	return hexRaster
}

func (r *Raster) ToByte() []byte {
	byteRaster, rasterConvertErr := WriteRasterWkb(*r)
	if rasterConvertErr != nil {
		log.Fatal(rasterConvertErr)
	}
	return byteRaster
}

// generate the Raster Statistics for all Bands
func (r *Raster) GetStatistics(precisions []int) []Statistic {
	result := make([]Statistic, len(r.Bands))
	in := make(chan int, len(r.Bands))
	var wg sync.WaitGroup

	for i := 0; i < r.Workers; i++ {
		wg.Add(1)
		go func() {
			for bandIdx := range in {
				if r.Bands[bandIdx].IsOffline {
					continue
				}
				result[bandIdx] = calculateBandStatistic(&r.Bands[bandIdx], precisions[bandIdx])
			}
			wg.Done()
		}()
	}
	for idx := range r.Bands {
		in <- idx
	}
	close(in)
	wg.Wait()
	return result
}

func (r *Raster) GetCoordinateFromIndexOnRaster(columnIdx, rowIdx int) (float64, float64) {
	startX := r.IpX + (r.ScaleX / 2.0)
	startY := r.IpY + (r.ScaleY / 2.0)
	return startX + (float64(columnIdx) * r.ScaleX),
		startY - (float64(rowIdx) * math.Abs(r.ScaleY))
}

func (r *Raster) GetIndexFromCoordinateOnRaster(band int, x, y float64) (int, int, error) {
	if x < r.MidIpX || y > r.MidIpY {
		return 0, 0, errors.New("outside")
	}
	idxX := (x - r.MidIpX) / math.Abs(r.ScaleX)
	idxY := (r.MidIpY - y) / math.Abs(r.ScaleY)
	if len(r.Bands[band].Data[0]) <= int(idxX) || len(r.Bands[band].Data) <= int(idxY) {
		return 0, 0, errors.New("outside")
	}
	return int(idxX),
		int(idxY),
		nil
}

func (r *Raster) GetMinMaxXY() (float64, float64, float64, float64) {
	return r.IpX, r.IpX + math.Floor(float64(r.Width)*r.ScaleX),
		r.IpY, r.IpY + math.Floor(float64(r.Height)*r.ScaleY)
}

func (r *Raster) BoundingBox() geometry.Polygon {
	return geometry.NewPolygon([][][]float64{
		{
			{r.IpX, r.IpY},
			{r.IpX + math.Floor(float64(r.Width)*r.ScaleX), r.IpY},
			{r.IpX + math.Floor(float64(r.Width)*r.ScaleX), r.IpY + math.Floor(float64(r.Height)*r.ScaleY)},
			{r.IpX, r.IpY + math.Floor(float64(r.Height)*r.ScaleY)},
			{r.IpX, r.IpY},
		},
	}, reference_system.New(int(r.Srid)))
}
