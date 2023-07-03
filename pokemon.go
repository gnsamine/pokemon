package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Name struct {
	Name string `json:"name"`
}

func main() {
	resp, err := http.Get("https://pokeapi.co/api/v2/type/ghost/")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

}

type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
