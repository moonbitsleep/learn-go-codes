package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Name string  `json:"landmark"`
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func main() {
	locations := []location {
		{Name: "Bradbury Landing", Lat: -4.567, Long: 137.9878},
		{Name: "Columbia Memorial", Lat: -14.5678, Long: 175.46729},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.8793},
	}
	for _, loc :=range locations {
		bytes, err := json.Marshal(loc)
		exitOnError(err)
		fmt.Println(string(bytes))
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}