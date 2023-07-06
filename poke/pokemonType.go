package poke

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"unsafe"
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

	var p Pokemon
	json.Unmarshal([]byte(stringBody), &p)

	fmt.Println("Names of Ghost Pokemons")

	fmt.Println("Front Default Sprites of Ghost Pokemons")

	for i := 0; i < len(p.Pokemon); i++ {

		a := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%s", p.Pokemon[i].Pokemon.Name)

		resp, _ := http.Get(a)

		body, _ := io.ReadAll(resp.Body)

		stringBody := string(body)

		var s Pokemon2
		json.Unmarshal([]byte(stringBody), &s)

		fmt.Println(s.Sprites.FrontDefault)

		resp2, _ := http.Get(s.Sprites.FrontDefault)

		body2, _ := io.ReadAll(resp2.Body)

		body2byte := *(*[]byte)(unsafe.Pointer(&body2))

		fileName := "images/" + p.Pokemon[i].Pokemon.Name + ".png"

		file, _ := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		body2Reader := bytes.NewReader(body2byte)
		_, err = io.Copy(file, body2Reader)
		if err != nil {
			fmt.Println(err)

		}

	}

}
