package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupedTemperatures := groupTemperatures(temperatures, 10)

	for group, values := range groupedTemperatures {
		fmt.Printf("%d: %v\n", group, values)
	}
}

func groupTemperatures(temperatures []float64, step float64) map[int][]float64 {
	groupedTemperatures := make(map[int][]float64)

	for _, temp := range temperatures {
		group := int(temp/step) * int(step) // Определяем группу
		groupedTemperatures[group] = append(groupedTemperatures[group], temp)
	}

	return groupedTemperatures
}
