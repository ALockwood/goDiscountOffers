// Copyright 2014 clypd, inc.
//
// see /LICENSE file for more information
//

package munkres

import (
	"math"

	"github.com/alockwood/goDiscountOffers/floatMatrix"
)

const (
	Unset mark = iota
	Starred
	Primed
	zero64 = int64(0)
)

type mark int

type Context struct {
	m          *floatMatrix.FloatMatrix
	rowCovered []bool
	colCovered []bool
	marked     []mark
	z0row      int64
	z0column   int64
	rowPath    []int64
	colPath    []int64
}

type Step interface {
	Compute(*Context) (Step, bool)
}

type Step1 struct{}
type Step2 struct{}
type Step3 struct{}
type Step4 struct{}
type Step5 struct{}
type Step6 struct{}

type RowCol struct {
	Row, Col int64
}

func newContext(m *floatMatrix.FloatMatrix) *Context {
	ctx := Context{
		m: &floatMatrix.FloatMatrix{
			A: make([]float64, m.N*m.N),
			N: m.N,
		},
		rowPath: make([]int64, 2*m.N),
		colPath: make([]int64, 2*m.N),
		marked:  make([]mark, m.N*m.N),
	}
	copy(ctx.m.A, m.A)
	clearCovers(&ctx)
	return &ctx
}

func min(a ...float64) float64 {
	min := math.MaxFloat64
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func (Step1) Compute(ctx *Context) (Step, bool) {
	n := ctx.m.N
	for i := zero64; i < n; i++ {
		row := ctx.m.A[i*n : (i+1)*n]
		minval := min(row...)
		for idx := range row {
			row[idx] -= minval
		}
	}
	return Step2{}, false
}

func clearCovers(ctx *Context) {
	n := ctx.m.N
	ctx.rowCovered = make([]bool, n)
	ctx.colCovered = make([]bool, n)
}

func (Step2) Compute(ctx *Context) (Step, bool) {
	n := ctx.m.N
	for i := zero64; i < n; i++ {
		rowStart := i * n
		for j := zero64; j < n; j++ {
			pos := rowStart + j
			if (ctx.m.A[pos] == 0) &&
				!ctx.colCovered[j] && !ctx.rowCovered[i] {
				ctx.marked[pos] = Starred
				ctx.colCovered[j] = true
				ctx.rowCovered[i] = true
			}
		}
	}
	clearCovers(ctx)
	return Step3{}, false
}

func (Step3) Compute(ctx *Context) (Step, bool) {
	n := ctx.m.N
	count := zero64
	for i := zero64; i < n; i++ {
		rowStart := i * n
		for j := zero64; j < n; j++ {
			pos := rowStart + j
			if ctx.marked[pos] == Starred {
				ctx.colCovered[j] = true
				count++
			}
		}
	}
	if count >= n {
		return nil, true
	}

	return Step4{}, false
}

func findAZero(ctx *Context) (int64, int64) {
	row := int64(-1)
	col := int64(-1)
	n := ctx.m.N
Loop:
	for i := zero64; i < n; i++ {
		rowStart := i * n
		for j := zero64; j < n; j++ {
			if (ctx.m.A[rowStart+j] == 0) &&
				!ctx.rowCovered[i] && !ctx.colCovered[j] {
				row = i
				col = j
				break Loop
			}
		}
	}
	return row, col
}

func findStarInRow(ctx *Context, row int64) int64 {
	n := ctx.m.N
	for j := zero64; j < n; j++ {
		if ctx.marked[row*n+j] == Starred {
			return j
		}
	}
	return -1
}

func (Step4) Compute(ctx *Context) (Step, bool) {
	starCol := int64(-1)
	for {
		row, col := findAZero(ctx)
		if row < 0 {
			return Step6{}, false
		}
		n := ctx.m.N
		pos := row*n + col
		ctx.marked[pos] = Primed
		starCol = findStarInRow(ctx, row)
		if starCol >= 0 {
			col = starCol
			ctx.rowCovered[row] = true
			ctx.colCovered[col] = false
		} else {
			ctx.z0row = row
			ctx.z0column = col
			break
		}
	}
	return Step5{}, false
}

func findStarInCol(ctx *Context, col int64) int64 {
	n := ctx.m.N
	for i := zero64; i < n; i++ {
		if ctx.marked[i*n+col] == Starred {
			return i
		}
	}
	return -1
}

func findPrimeInRow(ctx *Context, row int64) int64 {
	n := ctx.m.N
	for j := zero64; j < n; j++ {
		if ctx.marked[row*n+j] == Primed {
			return j
		}
	}
	return -1
}

func convertPath(ctx *Context, count int) {
	n := ctx.m.N
	for i := 0; i < count+1; i++ {
		r, c := ctx.rowPath[i], ctx.colPath[i]
		offset := r*n + c
		if ctx.marked[offset] == Starred {
			ctx.marked[offset] = Unset
		} else {
			ctx.marked[offset] = Starred
		}
	}
}

func erasePrimes(ctx *Context) {
	n := ctx.m.N
	for i := zero64; i < n; i++ {
		rowStart := i * n
		for j := zero64; j < n; j++ {
			if ctx.marked[rowStart+j] == Primed {
				ctx.marked[rowStart+j] = Unset
			}
		}
	}
}

func (Step5) Compute(ctx *Context) (Step, bool) {
	count := 0
	ctx.rowPath[count] = ctx.z0row
	ctx.colPath[count] = ctx.z0column
	var done bool
	for !done {
		row := findStarInCol(ctx, ctx.colPath[count])
		if row >= 0 {
			count++
			ctx.rowPath[count] = row
			ctx.colPath[count] = ctx.colPath[count-1]
		} else {
			done = true
		}

		if !done {
			col := findPrimeInRow(ctx, ctx.rowPath[count])
			count++
			ctx.rowPath[count] = ctx.rowPath[count-1]
			ctx.colPath[count] = col
		}
	}
	convertPath(ctx, count)
	clearCovers(ctx)
	erasePrimes(ctx)
	return Step3{}, false
}

func findSmallest(ctx *Context) float64 {
	n := ctx.m.N
	minval := math.MaxFloat64
	for i := zero64; i < n; i++ {
		rowStart := i * n
		for j := zero64; j < n; j++ {
			if (!ctx.rowCovered[i]) && (!ctx.colCovered[j]) {
				a := ctx.m.A[rowStart+j]
				if minval > a {
					minval = a
				}
			}
		}
	}
	return minval
}

func (Step6) Compute(ctx *Context) (Step, bool) {
	n := ctx.m.N
	minval := findSmallest(ctx)
	for i := zero64; i < n; i++ {
		rowStart := i * n
		for j := zero64; j < n; j++ {
			if ctx.rowCovered[i] {
				ctx.m.A[rowStart+j] += minval
			}
			if !ctx.colCovered[j] {
				ctx.m.A[rowStart+j] -= minval
			}
		}
	}
	return Step4{}, false
}

func ComputeMunkres(m *floatMatrix.FloatMatrix) []RowCol {
	ctx := newContext(m)

	var step Step
	step = Step1{}
	for {
		nextStep, done := step.Compute(ctx)

		if done {
			break
		}
		step = nextStep
	}
	results := []RowCol{}
	n := m.N
	for i := zero64; i < n; i++ {
		rowStart := i * n
		for j := zero64; j < n; j++ {
			if ctx.marked[rowStart+j] == Starred {
				results = append(results, RowCol{i, j})
			}
		}
	}
	return results
}
