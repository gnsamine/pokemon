package main

import (
	"example.com/gnsamine/pokemon/poke"
)

func main() {
	// resp, err := http.Get("https://pokeapi.co/api/v2/type/ghost/")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// stringBody := string(body)
	// fmt.Println(stringBody)

	// we use "var data map[string]interface{}"
	//bc values can be of any type with interface{}

	//fmt.Println(poke.SampleData)

	poke.GetNames()

}
