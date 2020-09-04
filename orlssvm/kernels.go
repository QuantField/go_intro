package main

import "gonum.org/v1/gonum/mat"

type Kernel interface {
	GetType() string
	Evaluate(x, xt *mat.Matrix) *mat.Matrix
	SetParams(params ...float64)
	GetParams() []float64
}

type Linear struct {
	kernType string
}

func (l Linear) GetType() {
	return l.kernType

}
