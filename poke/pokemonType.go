package poke

import (
	"encoding/json"
	"fmt"
)

type Pokemon struct {
	Pokemon []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon"`

	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

func GetNames() {
	jsonData := SampleData

	var p Pokemon
	json.Unmarshal([]byte(jsonData), &p)

	fmt.Println("Names of Ghost Pokemons")
	for i := 0; i < len(p.Pokemon); i++ {
		fmt.Println(p.Pokemon[i].Pokemon.Name)
	}
}

// func GetSprites() {
// 	resp, err := http.Get("https://pokeapi.co/api/v2/type/ghost/")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	stringBody := string(body)
// 	fmt.Println(stringBody)

// 	// we use "var data map[string]interface{}"
// 	//bc values can be of any type with interface{}

// }
