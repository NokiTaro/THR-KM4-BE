package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		return append([]int{newData}, data...)
	} else if position == "down" {
		return append(data, newData)
	} else {
		panic("Position should be either 'up' or 'down'")
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	newData := 6
	position := "up"
	result := AddElement(data, newData, position)
	fmt.Println("Input : ", result)

	position = "down"
	result = AddElement(data, newData, position)
	fmt.Println("Output : ",result)
}