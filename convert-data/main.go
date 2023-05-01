package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	values := strings.Split(data, ",")
	age, err := strconv.Atoi(values[1])
	if err != nil {
		panic("Invalid age")
	}
	return User{Name: values[0], Age: age, Address: values[2]}
}

func main() {
	data := "Edo,20,Jakarta"
	result := ConvertData(data)
	fmt.Println("Input : " + data)
	fmt.Printf("Output: " + "%+v\n", result)

	data = "Budi,30,Bandung"
	result = ConvertData(data)
	fmt.Println("Input : " + data)
	fmt.Printf("Output : " + "%+v\n", result)
}
