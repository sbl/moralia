package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Aphorism is exactly that by T.W. Adorno.
type Aphorism struct {
	Title string
	Text  string
}

// Aphorisms is the collection found in "Minima Moralia".
type Aphorisms []Aphorism

func main() {
	rand.Seed(time.Now().Unix())

	var as Aphorisms
	if err := json.Unmarshal([]byte(source), &as); err != nil {
		fmt.Printf("decoding error: %s\n", err)
		os.Exit(1)
	}
	a := as[rand.Intn(len(as))]
	fmt.Printf("%s - %s", a.Title, a.Text)
}
