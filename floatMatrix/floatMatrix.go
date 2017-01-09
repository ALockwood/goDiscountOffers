package floatMatrix

import (
	"fmt"
)

type FloatVector []float64

type FloatMatrix struct {
	N int64
	A []float64
}

func NewMatrix(n int64) (m *FloatMatrix) {
	m = new(FloatMatrix)
	m.N = n
	m.A = make([]float64, n*n)
	return m
}

func (m FloatMatrix) Get(i int64, j int64) float64 {
	return m.A[i*m.N+j]
}

func (m FloatMatrix) Set(i int64, j int64, v float64) {
	m.A[i*m.N+j] = v
}

func (v FloatVector) Len() int64 {
	return int64(len(v))
}

func (v FloatVector) Print() {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%f ", v[i])
	}
	fmt.Print("\n")
}

func (m *FloatMatrix) Print() {
	var i, j int64
	for i = 0; i < m.N; i++ {
		for j = 0; j < m.N; j++ {
			fmt.Printf("%f ", m.Get(i, j))
		}
		fmt.Print("\n")
	}
}
