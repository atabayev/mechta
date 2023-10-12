package main

import (
	"flag"
	"fmt"
	"log"
	"mechta/internal/pkg/mechta"
	"os"
	"time"
)

//easyjson:json
type Data []struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	gc := flag.Int("gc", 1, "Goroutine count. Must be: 1 <= gc <= 9223372036854775807")
	filename := flag.String("file", "data.json", "The path to file containing JSON data.")
	flag.Parse()
	if *gc < 1 && *gc <= 9223372036854775807 {
		println("The number of goroutines must be greater than 0")
		os.Exit(1)
	}
	startTime := time.Now()

	a := mechta.New(*filename, *gc)
	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}

	println(fmt.Sprintf("Bench time: %.2f seconds", float64(time.Since(startTime).Milliseconds())*0.001))
}
