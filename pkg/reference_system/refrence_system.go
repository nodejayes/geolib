package reference_system

import (
	"bytes"
	"github.com/everystreet/go-proj/cproj"
	"github.com/everystreet/go-proj/proj"
	"log"
	"strconv"
	"strings"
)

type Coordinate struct {
	X float64
	Y float64
}

func (c Coordinate) PutCoordinate(coord *cproj.PJ_COORD) {
	panic("implement me")
}

func (c Coordinate) FromCoordinate(coord cproj.PJ_COORD) {
	panic("implement me")
}

type ReferenceSystemProperties struct {
	Name string `json:"name"`
}

type ReferenceSystem struct {
	Type       string                    `json:"type"`
	Properties ReferenceSystemProperties `json:"properties"`
}

func New(epsgCode int) ReferenceSystem {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(systemHeader)
	buf.WriteString(strconv.FormatInt(int64(epsgCode), 10))
	return ReferenceSystem{
		Type: systemType,
		Properties: ReferenceSystemProperties{
			Name: buf.String(),
		},
	}
}

func (rf *ReferenceSystem) GetSrId() int {
	srid, err := strconv.ParseInt(strings.ReplaceAll(rf.Properties.Name, systemHeader, ""), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(srid)
}

func (rf *ReferenceSystem) TransformPoints(target int, points [][]float64) ([][]float64, error) {
	var tmp [][]float64
	for _, p := range points {
		xy := Coordinate{
			X: p[0],
			Y: p[1],
		}
		err := proj.CRSToCRS(
			proj.CRS("EPSG:"+strconv.FormatInt(int64(rf.GetSrId()), 10)),
			proj.CRS("EPSG:"+strconv.FormatInt(int64(target), 10)),
			proj.TransformForward(&xy),
		)
		if err != nil {
			return nil, err
		}
		tmp = append(tmp, []float64{xy.X, xy.Y})
	}

	return tmp, nil
}
