package nn

type Chain struct {
	hidden *layer
	output *layer
}

type layer struct {
	weight *Matrix
	data   *Matrix
}

func NewChain(inN, hiddenDepth, outN int) *Chain {
	ret := new(Chain)
	ret.hidden = newLayer(inN, hiddenDepth)
	ret.output = newLayer(hiddenDepth, outN)
	return ret
}

func newLayer(m, n int) *layer {
	ret := new(layer)
	ret.weight = NewMatrix(m, n).FillRandGuass(0, 1)
	return ret
}

/*
https://www.youtube.com/watch?v=UJwK6jAStmg
z_2 = x * w_1
a_2 = f(z_2)
z_3 = a_2 * w_2
y = f(z_3)

x: input
w_1: hidden.weight
f(): Activation func
a_2: hidden.data
w_2: output.weight
y: output
*/
func (chain *Chain) forward(input *Matrix) *Matrix {
	hiddenZ := input.Dot(chain.hidden.weight)
	chain.hidden.data = hiddenZ.Map(func(x float64) float64 {
		return Sigmoid(x)
	})

	outputZ := chain.hidden.data.Dot(chain.output.weight)
	chain.output.data = outputZ.Map(func(x float64) float64 {
		return Sigmoid(x)
	})
	return chain.output.data
}

func (chain *Chain) backward(input, expected *Matrix) *Matrix {
	outputError := expected.Sub(chain.output.data)
	outputSigmoidDerivative := chain.output.data.Map(func(x float64) float64 {
		return SigmoidDerivative(x)
	})
	outputDelta := outputError.Mul(outputSigmoidDerivative)
	outputDw := chain.hidden.data.T().Dot(outputDelta)

	hiddenError := outputDelta.Dot(chain.output.weight.T())
	hiddenSigmoidDerivative := chain.hidden.data.Map(func(x float64) float64 {
		return SigmoidDerivative(x)
	})
	hiddenDelta := hiddenError.Mul(hiddenSigmoidDerivative)
	hiddenDw := input.T().Dot(hiddenDelta)

	chain.output.weight = chain.output.weight.Add(outputDw)
	chain.hidden.weight = chain.hidden.weight.Add(hiddenDw)
	return outputError
}

func (chain *Chain) Train(input, expected *Matrix) *Matrix {
	chain.forward(input)
	return chain.backward(input, expected)
}

func (chain *Chain) Predict(input *Matrix) *Matrix {
	return chain.forward(input)
}
