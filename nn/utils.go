package nn

import (
	"math"
	"math/rand"
	"time"
)

const (
	EPSILON float64 = 1e-9
)

var (
	guassCache      float64 = 0.0
	guassCacheReady bool    = false
)

func RandSeed() {
	rand.Seed(time.Now().Unix())
}

func LikeZero(x float64) bool {
	return math.Abs(x) < EPSILON
}

func LikeInf(x float64) bool {
	return x >= math.Inf(1) || x <= math.Inf(-1)
}

//https://www.google.co.nz/search?q=Exp(-x)&oq=Exp(-x)&aqs=chrome..69i57j0l5.415j0j7&sourceid=chrome&ie=UTF-8#q=1%2F(1%2BExp(-x))
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

//https://www.google.co.nz/webhp?sourceid=chrome-instant&ion=1&espv=2&ie=UTF-8#q=x*(1-x)
func SigmoidDerivative(x float64) float64 {
	return x * (1.0 - x)
}

//https://www.google.co.nz/search?q=tanh(x)&oq=tanh(x)&aqs=chrome..69i57j0l5.438j0j9&sourceid=chrome&ie=UTF-8
func Tanh(x float64) float64 {
	return math.Tanh(x)
}

//https://www.google.co.nz/search?q=1+-+x+*+x&oq=1+-+x+*+x&aqs=chrome..69i57.398j0j7&sourceid=chrome&ie=UTF-8
func TanhDerivative(x float64) float64 {
	return 1 - x*x
}

//https://en.wikipedia.org/wiki/Marsaglia_polar_method
func RandStdGuass() float64 {
	if guassCacheReady {
		guassCacheReady = false
		return guassCache
	}
	var u, v, s float64
	for s >= 1.0 || s == 0.0 {
		u = 2.0*rand.Float64() - 1.0
		v = 2.0*rand.Float64() - 1.0
		s = u*u + v*v
	}
	s = math.Sqrt(-2.0 * math.Log(s) / s)
	guassCache = v * s
	guassCacheReady = true
	return u * s
}

func RandGuass(mean, stdDev float64) float64 {
	return mean + RandStdGuass()*stdDev
}
