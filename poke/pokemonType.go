package poke

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Pokemon []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon"`
}

type Pokemon2 struct {
	Sprites struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`
}

//func GetNames() {
// jsonData := SampleData

// var p Pokemon
// json.Unmarshal([]byte(jsonData), &p)

// fmt.Println("Names of Ghost Pokemons")
// for i := 0; i < len(p.Pokemon); i++ {
// 	fmt.Println(p.Pokemon[i].Pokemon.Name)
// }
//}

func GetSprites() {
	resp, err := http.Get("https://pokeapi.co/api/v2/type/ghost/")
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	stringBody := string(body)
	//fmt.Println(stringBody)

	var p Pokemon
	json.Unmarshal([]byte(stringBody), &p)

	fmt.Println("Names of Ghost Pokemons")

	for i := 0; i < len(p.Pokemon); i++ {
		fmt.Println(p.Pokemon[i].Pokemon.Name)

	}

	fmt.Println("Front DEfoult Sprites of Ghost Pokemons")

	for i := 0; i < len(p.Pokemon); i++ {

		a := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%s", p.Pokemon[i].Pokemon.Name)
		//fmt.Println(a)

		resp, _ := http.Get(a)

		body, _ := io.ReadAll(resp.Body)

		stringBody := string(body)

		var s Pokemon2
		json.Unmarshal([]byte(stringBody), &s)

		fmt.Println(s.Sprites.FrontDefault)

	}

}

func GetPNG() {
	resp, err := http.Get("https://pokeapi.co/api/v2/type/ghost/")
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
}
