package main

import (
	"fmt"
	"go-readfile-variable/data"
)

func main() {
	karnData := data.GetDataFromFile("data/variables.yml")
	fmt.Printf("Hello %s , Your's weight is %s",karnData.FullName ,karnData.Body.Weight )
}
