package nn

import (
	"bytes"
	"fmt"
)

type Matrix struct {
	M, N int
	Data [][]float64
}

func NewMatrix(m, n int) *Matrix {
	ret := new(Matrix)
	ret.M = m
	ret.N = n
	ret.Data = make([][]float64, ret.M)
	for i := 0; i < ret.M; i++ {
		ret.Data[i] = make([]float64, ret.N)
	}
	return ret
}

func (mat *Matrix) String() string {
	var retBuff bytes.Buffer
	for _, row := range mat.Data {
		retBuff.WriteString(fmt.Sprintln(row))
	}
	return retBuff.String()
}

func (mat *Matrix) Clone() *Matrix {
	ret := NewMatrix(mat.M, mat.N)
	for i := 0; i < mat.M; i++ {
		for j := 0; j < mat.N; j++ {
			ret.Data[i][j] = mat.Data[i][j]
		}
	}
	return ret
}

// Transpose
func (mat *Matrix) T() *Matrix {
	ret := NewMatrix(mat.N, mat.M)
	for i := 0; i < mat.M; i++ {
		for j := 0; j < mat.N; j++ {
			ret.Data[j][i] = mat.Data[i][j]
		}
	}
	return ret
}

func (mat *Matrix) Reshape(m, n int) *Matrix {
	if m*n != mat.M*mat.N {
		return nil
	}
	ret := NewMatrix(m, n)
	for i := 0; i < mat.M; i++ {
		for j := 0; j < mat.N; j++ {
			k := i*mat.N + j
			ret.Data[k/n][k%n] = mat.Data[i][j]
		}
	}
	return ret
}

func (mat *Matrix) Fill(value float64) *Matrix {
	for i := 0; i < mat.M; i++ {
		for j := 0; j < mat.N; j++ {
			mat.Data[i][j] = value
		}
	}
	return mat
}

func (mat *Matrix) FillArray(data []float64) *Matrix {
	if len(data) != mat.M*mat.N {
		return nil
	}
	for i := 0; i < mat.M; i++ {
		for j := 0; j < mat.N; j++ {
			mat.Data[i][j] = data[i*mat.N+j]
		}
	}
	return mat
}

func (mat *Matrix) FillRandGuass(mu, std float64) *Matrix {
	for i := 0; i < mat.M; i++ {
		for j := 0; j < mat.N; j++ {
			mat.Data[i][j] = RandGuass(mu, std)
		}
	}
	return mat
}

func (x *Matrix) Dot(y *Matrix) *Matrix {
	if x.N != y.M {
		return nil
	}
	ret := NewMatrix(x.M, y.N)
	for i := 0; i < x.M; i++ {
		for j := 0; j < y.N; j++ {
			var sum float64
			for k := 0; k < x.N; k++ {
				sum += x.Data[i][k] * y.Data[k][j]
			}
			ret.Data[i][j] = sum
		}
	}
	return ret
}

func (x *Matrix) Mul(y *Matrix) *Matrix {
	if x.M != y.M && x.N != y.N {
		return nil
	}
	ret := NewMatrix(x.M, x.N)
	for i := 0; i < x.M; i++ {
		for j := 0; j < x.N; j++ {
			ret.Data[i][j] = x.Data[i][j] * y.Data[i][j]
		}
	}
	return ret
}

func (x *Matrix) Scale(s float64) *Matrix {
	ret := NewMatrix(x.M, x.N)
	for i := 0; i < x.M; i++ {
		for j := 0; j < x.N; j++ {
			ret.Data[i][j] = x.Data[i][j] * s
		}
	}
	return ret
}

func (x *Matrix) Add(y *Matrix) *Matrix {
	if x.M != y.M && x.N != y.N {
		return nil
	}
	ret := NewMatrix(x.M, x.N)
	for i := 0; i < x.M; i++ {
		for j := 0; j < x.N; j++ {
			ret.Data[i][j] = x.Data[i][j] + y.Data[i][j]
		}
	}
	return ret
}

func (x *Matrix) ScalableAdd(y *Matrix, s1, s2 float64) *Matrix {
	if x.M != y.M && x.N != y.N {
		return nil
	}
	ret := NewMatrix(x.M, x.N)
	for i := 0; i < x.M; i++ {
		for j := 0; j < x.N; j++ {
			ret.Data[i][j] = x.Data[i][j]*s1 + y.Data[i][j]*s2
		}
	}
	return ret
}

func (x *Matrix) Sub(y *Matrix) *Matrix {
	if x.M != y.M && x.N != y.N {
		return nil
	}
	ret := NewMatrix(x.M, x.N)
	for i := 0; i < x.M; i++ {
		for j := 0; j < x.N; j++ {
			ret.Data[i][j] = x.Data[i][j] - y.Data[i][j]
		}
	}
	return ret
}

func (x *Matrix) ScalableSub(y *Matrix, s1, s2 float64) *Matrix {
	if x.M != y.M && x.N != y.N {
		return nil
	}
	ret := NewMatrix(x.M, x.N)
	for i := 0; i < x.M; i++ {
		for j := 0; j < x.N; j++ {
			ret.Data[i][j] = x.Data[i][j]*s1 - y.Data[i][j]*s2
		}
	}
	return ret
}

func (x *Matrix) Map(f func(x float64) float64) *Matrix {
	ret := NewMatrix(x.M, x.N)
	for i := 0; i < x.M; i++ {
		for j := 0; j < x.N; j++ {
			ret.Data[i][j] = f(x.Data[i][j])
		}
	}
	return ret
}
