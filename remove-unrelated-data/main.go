package main

import "fmt"

func removeUnrelatedData(mapData map[string]interface{}, key string) map[string]interface{} {
    delete(mapData, key)
    return mapData
}

func main() {
    data1 := map[string]interface{}{"name": "Edo", "age": 20, "address": "Jakarta"}
    fmt.Println(removeUnrelatedData(data1, "address")) // map[age:20 name:Edo]

    data2 := map[string]interface{}{"run": "lari", "jump": "loncat", "swim": "berenang"}
    fmt.Println(removeUnrelatedData(data2, "jump")) // map[run:lari swim:berenang]

    data3 := map[string]interface{}{"satu": "ichi", "dua": "ni", "tiga": "san", "empat": "yon", "lima": "go"}
    fmt.Println(removeUnrelatedData(data3, "tiga")) // map[dua:ni empat:yon lima:go satu:ichi]
}