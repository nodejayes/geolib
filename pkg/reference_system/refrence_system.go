package reference_system

import (
	"bytes"
	"github.com/nodejayes/geolib/pkg/proj4"
	"log"
	"strconv"
	"strings"
)

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
	var err error
	tmp, err = proj4.ReprojectPoints(points, proj4.GetProjection(rf.GetSrId()), proj4.GetProjection(target))
	if err != nil {
		return nil, err
	}
	return tmp, nil
}
