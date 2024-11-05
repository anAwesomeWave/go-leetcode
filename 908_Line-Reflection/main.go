package main

import (
	"fmt"
	"math"
)

func IsReflected(points [][]int) bool {
	//https://www.lintcode.com/problem/908/description

	// Write your code here
	// 1. понять где такая линия может быть -> xMin, xMax O(n)
	// 2. для каждой точки проверить, есть ли симм. такая что,
	// она по другую сторону и с такой же y коорд.
	// map[int]map[int]struct{} -> x coord -> set of y coords

	// space O(n)
	// time O(n)
	if len(points) == 0 {
		return true
	}
	xMin, xMax := points[0][0], points[0][0]
	mapOfPoints := map[int]map[int]int{}
	for _, v := range points {
		if v[0] > xMax {
			xMax = v[0]
		}
		if v[0] < xMin {
			xMin = v[0]
		}
		if _, ok := mapOfPoints[v[0]]; !ok {
			mapOfPoints[v[0]] = make(map[int]int)
		}
		mapOfPoints[v[0]][v[1]] += 1
	}

	linePos := float64(xMax+xMin) / 2.0
	for _, v := range points {
		// works in all cases. no matter left or right point.
		oppCoord := v[0] + int(math.Round((linePos-float64(v[0]))*2))
		if opposite := mapOfPoints[oppCoord][v[1]]; opposite > 0 {
			mapOfPoints[oppCoord][v[1]] -= 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsReflected([][]int{
		{0, 10},
		{5, 1},
		{6, 1},
		{11, 10},
	}))
}
