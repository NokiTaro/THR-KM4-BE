package main

import "fmt"

func howManyElements(data []interface{}) int {
	return len(data)
}

func main() {
	slice1 := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(howManyElements(slice1))

	slice2 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(howManyElements(slice2))

	slice3 := []interface{}{1, 1, 1, 5, 5, 5}
	fmt.Println(howManyElements(slice3))

	slice4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(slice4))

	slice5 := []interface{}{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(slice5))

	slice6 := []interface{}{true, false, true, false, true, false}
	fmt.Println(howManyElements(slice6))
}