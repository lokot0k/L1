package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupedTemps := getGroupedTemps(temps)
	fmt.Println("Группировка температур:")
	for key, values := range groupedTemps {
		isFirst := true
		out := ""
		for temp := range values {
			if !isFirst {
				out += fmt.Sprintf(", %.1f", temp)
			} else {
				out += fmt.Sprintf("%.1f", temp)
			}
			isFirst = false
		}
		fmt.Printf("%d: %s\n", key, out)
	}
}

func getTempLowerBound(temp float64) int {
	if temp >= 0 {
		return int(math.Floor(temp/10)) * 10
	} else {
		return int(math.Ceil(temp/10)) * 10
	}
}

func getGroupedTemps(temps []float64) map[int]map[float64]struct{} {
	groupedTemps := make(map[int]map[float64]struct{}) // используем set, так как температуры могут дублироваться
	for _, temp := range temps {
		key := getTempLowerBound(temp)
		if groupedTemps[key] == nil {
			groupedTemps[key] = make(map[float64]struct{})
		}
		groupedTemps[key][temp] = struct{}{}
	}
	return groupedTemps
}
