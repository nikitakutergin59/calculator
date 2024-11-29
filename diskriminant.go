package diskriminant

import (
	"errors"
	"math"
)

func CalculateDiscriminant(a, b, c float64) float64 {
	return (b * b) - 4*(a*c)
}

func CalculateRoots(a, b, discriminant float64) ([]float64, error) {
	if discriminant < 0 {
		return nil, errors.New("дискриминант меньше нуля, нет действительных корней")
	} else if discriminant == 0 {
		root := (-b) / (2 * a)
		return []float64{root}, nil
	} else {
		sqrtDiscriminant := math.Sqrt(discriminant)
		root1 := ((-b) + sqrtDiscriminant) / (2 * a)
		root2 := ((-b) - sqrtDiscriminant) / (2 * a)
		return []float64{root1, root2}, nil
	}
}
