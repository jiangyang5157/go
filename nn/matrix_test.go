package nn

import (
	"fmt"
	"testing"
)

func TestMatrix_String(t *testing.T) {
	mat := NewMatrix(4, 3)
	fmt.Println(mat.String())
}

func TestMatrix_Clone(t *testing.T) {
	mat := NewMatrix(4, 3)
	mat.Data[0][1] = 1
	matClone := mat.Clone()
	matClone.Data[0][1] = 2
	fmt.Printf("mat:\n%v\n", mat.String())
	fmt.Printf("Clone:\n%v\n", matClone.String())
}

func TestMatrix_Reshape(t *testing.T) {
	mat := NewMatrix(4, 3)
	mat.Data[0][1] = 1
	mat.Data[0][2] = 1
	mat.Data[1][0] = 1
	mat.Data[2][0] = 1
	mat.Data[3][0] = 1
	matReshape := mat.Reshape(1, 12)
	fmt.Printf("mat:\n%v\n", mat.String())
	fmt.Printf("Reshape:\n%v\n", matReshape.String())
}

func TestMatrix_T(t *testing.T) {
	mat := NewMatrix(4, 3)
	mat.Data[0][1] = 1
	matT := mat.T()
	matT.Data[0][1] = 2
	fmt.Printf("mat:\n%v\n", mat.String())
	fmt.Printf("matT:\n%v\n", matT.String())
}

func TestMatrix_Fill(t *testing.T) {
	mat := NewMatrix(4, 3)
	mat.Fill(2)
	fmt.Println(mat.String()) // 2...
}

func TestMatrix_FillArray(t *testing.T) {
	mat := NewMatrix(4, 3)
	mat.Data[0][1] = 1
	mat.Data[0][2] = 1
	mat.Data[1][0] = 1
	mat.Data[2][0] = 1
	mat.Data[3][0] = 1
	matReshape := mat.Reshape(1, 12)
	matFillArray := NewMatrix(3, 4)
	matFillArray.FillArray(matReshape.Data[0])
	fmt.Printf("Reshape:\n%v\n", matReshape.String())
	fmt.Printf("FillArray:\n%v\n", matFillArray.String())
}

func TestMatrix_FillRandGuass(t *testing.T) {
	RandSeed()
	mat := NewMatrix(4, 3)
	mat.FillRandGuass(0, 1)
	fmt.Printf("FillRandGuass:\n%v\n", mat.String())
}

func TestMatrix_Scale(t *testing.T) {
	mat := NewMatrix(4, 3)
	mat.Fill(2)
	mat = mat.Scale(3)
	fmt.Println(mat.String()) // 6...
}

func TestMatrix_ScalableAdd(t *testing.T) {
	mat := NewMatrix(4, 3)
	mat.Fill(2)
	mat2 := NewMatrix(4, 3)
	mat2.Fill(1)
	mat = mat.ScalableAdd(mat2, 3, 4)
	fmt.Println(mat.String()) // 10...
}

func TestMatrix_Add(t *testing.T) {
	mat := NewMatrix(4, 3)
	mat.Fill(2)
	mat2 := NewMatrix(4, 3)
	mat2.Fill(1)
	mat = mat.Add(mat2)
	fmt.Println(mat.String()) // 3...
}

func TestMatrix_Mul(t *testing.T) {
	mat := NewMatrix(4, 3)
	mat.Fill(2)
	mat2 := NewMatrix(4, 3)
	mat2.Fill(3)
	mat = mat.Mul(mat2)
	fmt.Println(mat.String()) // 6...
}

func TestMatrix_Dot(t *testing.T) {
	mat1 := NewMatrix(4, 3)
	mat1.Fill(2)
	mat2 := NewMatrix(3, 1)
	mat2.Fill(3)
	dot := mat1.Dot(mat2)
	fmt.Printf("mat1:\n%v\n", mat1.String()) // 2...
	fmt.Printf("mat2:\n%v\n", mat2.String()) // 3...
	fmt.Printf("dot:\n%v\n", dot.String())   // 18...
}
