package reference_system

import (
	"bytes"
	"errors"
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
	for _, p := range points {
		switch len(p) {
		case 3:
			tX, tY, tZ, projErr := proj4.ReprojectPoint(p[0], p[1], p[2], proj4.GetProjection(rf.GetSrId()), proj4.GetProjection(target))
			if projErr != nil {
				return nil, projErr
			}
			tmp = append(tmp, []float64{tX, tY, tZ})
			break
		case 2:
			tX, tY, _, projErr := proj4.ReprojectPoint(p[0], p[1], 0.0, proj4.GetProjection(rf.GetSrId()), proj4.GetProjection(target))
			if projErr != nil {
				return nil, projErr
			}
			tmp = append(tmp, []float64{tX, tY})
			break
		default:
			return nil, errors.New("invalid Point Length " + strconv.FormatInt(int64(len(p)), 10))
		}
	}
	return tmp, nil
}
