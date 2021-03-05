package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Lat float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func main() {
	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Printf("%T\n", bytes)
	fmt.Println(bytes)
	fmt.Println(string(bytes))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
