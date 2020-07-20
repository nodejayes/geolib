package polygon_tool

import (
	"encoding/json"
	"github.com/nodejayes/geolib/pkg/geometry"
	"github.com/nodejayes/geolib/pkg/reference_system"
	"io/ioutil"
	"os"
	"path"
)

type FillSpace struct{}

func (ctx *FillSpace) Fill(currentPoints [][]float64, otherPoints [][][]float64) {
	var lines [][][]float64
	for _, currentPoint := range currentPoints {
		lines = append(lines, ctx.buildConnections(currentPoint, otherPoints)...)
	}
	ctx.saveLines(lines)
}

// connect the current Point with all other points
func (ctx *FillSpace) buildConnections(currentPoint []float64, otherPoints [][][]float64) [][][]float64 {
	var lines [][][]float64
	for _, ring := range otherPoints {
		for _, point := range ring {
			/*
				xMovement, yMovement := ctx.getMovement(currentPoint, point)
				checkpoint := []float64{
					point[0], point[1],
				}
				if xMovement == "E" {
					checkpoint[0] = checkpoint[0] - 1
				} else {
					checkpoint[0] = checkpoint[0] + 1
				}
				if yMovement == "N" {
					checkpoint[1] = checkpoint[1] - 1
				} else {
					checkpoint[1] = checkpoint[1] + 1
				}
				if !detection.PointInPolygon(checkpoint, [][][]float64{
					ring,
				}, false) {
					continue
				}
			*/
			lines = append(lines, [][]float64{
				currentPoint, point,
			})
		}
	}
	return lines
}

func (ctx *FillSpace) getMovement(currentPoint []float64, point []float64) (string, string) {
	diffX := currentPoint[0] - point[0]
	diffY := currentPoint[1] - point[1]
	xMovement := ""
	yMovement := ""
	if diffX < 0 {
		xMovement = "W"
	} else {
		xMovement = "E"
	}
	if diffY < 0 {
		yMovement = "S"
	} else {
		yMovement = "N"
	}
	return xMovement, yMovement
}

func (ctx *FillSpace) saveLines(lines [][][]float64) {
	tmp := geometry.NewMultiLine(lines, reference_system.New(3857))
	stream, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	p := path.Join(os.TempDir(), "geolib")
	os.MkdirAll(p, 0755)
	ioutil.WriteFile(path.Join(p, "lines.json"), stream, 0755)
}
