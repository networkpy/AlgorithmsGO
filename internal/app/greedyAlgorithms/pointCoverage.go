package greedyalgorithms

import (
	"sort"
)

// Вводятся точки в виде слайса,
// необходимо найти минимальное количество отрезков,
// которыми можно покрыть эти точки
func PointCoverage(power []float64) [][2]float64 {
	sort.Float64s(power)
	var result [][2]float64
	i := 0
	for i < len(power) {
		section := [2]float64{power[i], power[i] + 1.0}
		result = append(result, section)
		i++
		for i < len(power) && power[i] <= result[len(result)-1][1] {
			i++
		}
	}
	return result
}
