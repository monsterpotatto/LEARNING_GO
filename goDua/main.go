package main

import "fmt"

func main() {
	fmt.Println(Average([]float64{1, 2, 3}))
}
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func GanjilGenap(angka int) string {
	if angka%2 == 0 {
		return "genap"
	}
	return "ganjil"
}
