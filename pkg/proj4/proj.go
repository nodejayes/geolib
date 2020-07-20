package proj4

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

var vm *otto.Otto
var proj4code []byte

func ReprojectPoint(x, y, z float64, source, target string) (float64, float64, float64, error) {
	setup()
	c, err := executeTransformation(x, y, z, source, target)
	if err != nil {
		return 0, 0, 0, err
	}
	converted := <-c
	return converted[0], converted[1], converted[2], nil
}

func executeTransformation(x, y, z float64, source, target string) (chan []float64, error) {
	var returnChan = make(chan []float64, 1)

	setErr := vm.Set("sendResult", func(call otto.FunctionCall) otto.Value {
		nx, execErr := call.Argument(0).ToFloat()
		ny, execErr := call.Argument(1).ToFloat()
		nz, execErr := call.Argument(2).ToFloat()
		if execErr != nil {
			returnChan <- []float64{0, 0, 0}
		}
		returnChan <- []float64{nx, ny, nz}
		return otto.Value{}
	})
	setErr = vm.Set("val_x", x)
	setErr = vm.Set("val_y", y)
	setErr = vm.Set("val_z", z)
	setErr = vm.Set("val_source", source)
	setErr = vm.Set("val_target", target)
	if setErr != nil {
		return nil, setErr
	}

	_, programErr := vm.Run(string(proj4code) + " " +
		"var pr = proj4(val_source, val_target, [val_x, val_y, val_z]);" +
		"sendResult(pr[0], pr[1], pr[2]);")
	if programErr != nil {
		return nil, programErr
	}
	return returnChan, nil
}

func setup() {
	if vm == nil {
		vm = otto.New()
	}
	if proj4code == nil {
		readProj4Code()
	}
}

func readProj4Code() {
	var err error
	proj4code, err = ioutil.ReadFile("proj4.js")
	if err != nil {
		proj4code = nil
		panic(err.Error())
	}
}
