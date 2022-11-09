package main

import (
	"fmt"
)

func main() {
	inputdariuser()
}

func inputdariuser() {
	fmt.Print("Masukkan sebuah kata: ")
	var input string
	fmt.Scanf("%f", &input)
	output := input
	fmt.Println(output)
}

func _lessfor() {
	//var x = []int64{1, 2, 3, 4}
	// for i, y := range x {
	// 	fmt.Println('%d %d\n', i, y)
	// }
}

func _lessfor2() {
	var x = []int64{1, 2, 3, 4, 10, 123, 423, 5245, 34657}

	for _, y := range x {
		if y%2 == 0 {
			fmt.Printf("%d adalah genap \n", y)
		} else {
			fmt.Printf("%d adalah ganjil \n", y)
		}
	}
}

func _array1() {
	var x [5]int
	x[4] = 10

	fmt.Println(x)

	//var y = new([]int)
	var y []int
	y = append(y, 10)

	for i := 0; i < 5; i++ {
		y = append(y, 1)
	}
	fmt.Println(y)
}

func _maps() {
	var x = make(map[string]int)
	x["test"] = 1
	delete(x, "test")
	fmt.Println(x["test"])
}

func _maps2() {
	var elements = make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["H"] = "Hydrogen"
	elements["H"] = "Hydrogen"
}

func _maps3() {
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

func average(xs []float64) (float64, int) {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs)), len(xs)
}

func average2(xs []float64) (average float64, ordered []float64) {
	total := 0.0
	//temp := 0.0
	for _, v := range xs {
		total += v
		// if v > float32(temp) {
		// 	ordered = append(ordered, v)
		// }
	}
	count := len(xs)
	average = total / float64(count)
	return
}
