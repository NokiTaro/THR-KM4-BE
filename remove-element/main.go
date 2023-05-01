package main

import "fmt"

func removeElement(array []interface{}, position string) []interface{} {
    if position == "left" {
        return array[1:]
    } else if position == "right" {
        return array[:len(array)-1]
    }
    return array
}

func main() {
    arr := []interface{}{1, 2, 3, 4, 5}
    fmt.Println(removeElement(arr, "left"))  // [2 3 4 5]
    fmt.Println(removeElement(arr, "right")) // [1 2 3 4]

    strArr := []interface{}{"Edo", "Budi", "Joko", "Tono"}
    fmt.Println(removeElement(strArr, "left"))  // [Budi Joko Tono]
    fmt.Println(removeElement(strArr, "right")) // [Edo Budi Joko]
}



