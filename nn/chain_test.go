package nn

import (
	"fmt"
	"testing"
)

func TestChain_AND(t *testing.T) {
	RandSeed()
	input := NewMatrix(4, 2).FillArray([]float64{
		0.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
		1.0, 1.0,
	})
	expected := NewMatrix(4, 1).FillArray([]float64{
		0.0,
		0.0,
		0.0,
		1.0,
	})
	c := NewChain(2, 4, 1)
	for i := 0; i < 60000; i++ {
		err := c.Train(input, expected)
		if i%10000 == 0 {
			fmt.Printf("err:\n%v\n", err.String())
		}
	}

	prediction := c.Predict(NewMatrix(8, 2).FillArray([]float64{
		0.0, 0.0,
		1.0, 1.0,
		0.0, 1.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
	}))
	fmt.Printf("prediction:\n%v\n", prediction.String())
}

func TestChain_OR(t *testing.T) {
	RandSeed()
	input := NewMatrix(4, 2).FillArray([]float64{
		0.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
		1.0, 1.0,
	})
	expected := NewMatrix(4, 1).FillArray([]float64{
		0.0,
		1.0,
		1.0,
		1.0,
	})
	c := NewChain(2, 4, 1)
	for i := 0; i < 60000; i++ {
		err := c.Train(input, expected)
		if i%10000 == 0 {
			fmt.Printf("err:\n%v\n", err.String())
		}
	}

	prediction := c.Predict(NewMatrix(8, 2).FillArray([]float64{
		0.0, 0.0,
		1.0, 1.0,
		0.0, 1.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
	}))
	fmt.Printf("prediction:\n%v\n", prediction.String())
}

func TestChain_XOR(t *testing.T) {
	RandSeed()
	input := NewMatrix(4, 2).FillArray([]float64{
		0.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
		1.0, 1.0,
	})
	expected := NewMatrix(4, 1).FillArray([]float64{
		0.0,
		1.0,
		1.0,
		0.0,
	})
	c := NewChain(2, 4, 1)
	for i := 0; i < 60000; i++ {
		err := c.Train(input, expected)
		if i%10000 == 0 {
			fmt.Printf("err:\n%v\n", err.String())
		}
	}

	prediction := c.Predict(NewMatrix(8, 2).FillArray([]float64{
		0.0, 0.0,
		1.0, 1.0,
		0.0, 1.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
	}))
	fmt.Printf("prediction:\n%v\n", prediction.String())
}

func TestChain_Add(t *testing.T) {
	RandSeed()
	input := NewMatrix(4, 8).FillArray([]float64{
		0.0, 1.0, 0.0, 0.0,
		0.0, 1.0, 1.0, 1.0,

		0.0, 0.0, 1.0, 1.0,
		0.0, 0.0, 1.0, 1.0,

		0.0, 1.0, 1.0, 0.0,
		0.0, 0.0, 1.0, 0.0,

		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
	})
	expected := NewMatrix(4, 4).FillArray([]float64{
		1.0, 0.0, 1.0, 1.0,
		0.0, 1.0, 1.0, 0.0,
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
	})
	c := NewChain(8, 12, 4)
	for i := 0; i < 60000; i++ {
		err := c.Train(input, expected)
		if i%10000 == 0 {
			fmt.Printf("err:\n%v\n", err.String())
		}
	}

	prediction := c.Predict(NewMatrix(1, 8).FillArray([]float64{
		0.0, 1.0, 1.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
	}))
	fmt.Printf("prediction:\n%v\n", prediction.String())
}
