package main

import (
	"fmt"

	skroutz "github.com/sger/go-skroutz"
)

func main() {
	resp, err := skroutz.Authorization("", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.AccessToken)

	s := skroutz.New(skroutz.NewConfig(""))

	out, err := s.Categories.GetCategories()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", out.Categories)
}
