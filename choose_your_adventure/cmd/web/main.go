package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/S-ayanide/GoExperiments/choose_your_adventure"
)

func main(){
	filename := flag.String("file","gopher.json","the JSON file with the Choose Your Adventure Story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story choose_your_adventure.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Println("%+v\n", story)
}
