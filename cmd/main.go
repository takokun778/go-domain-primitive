package main

import (
	"fmt"
	"go-domain-primitive/internal/sample"
	"log"
)

func main() {
	name, err := sample.Name.New("new")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name.String():\t\t %s\n", name.String())
	fmt.Printf("name.Value():\t\t %s\n", name.Value())

	name = sample.Name.Reconstruct("reconstruct")

	fmt.Printf("name.String():\t\t %s\n", name.String())
	fmt.Printf("name.Value():\t\t %s\n", name.Value())

	fmt.Printf("sample.Name.Value():\t\t %s\n", sample.Name.Value())
	fmt.Printf("sample.Name.String():\t %s\n", sample.Name.String())

	if _, err := sample.Name.New(""); err != nil {
		fmt.Printf("failed to new name error: %v\n", err)
	}
}
