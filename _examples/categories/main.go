package main

import (
	"fmt"

	skroutz "github.com/sger/go-skroutz"
)

func main() {
	s := skroutz.New(skroutz.NewConfig("<AccessToken>"))

	out, err := s.Categories.GetCategories()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", out.Categories)
}
