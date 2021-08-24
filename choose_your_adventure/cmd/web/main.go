package main

import (
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

	story, err := choose_your_adventure.JsonStory(f)
	
	fmt.Println("%+v\n", story)
}
