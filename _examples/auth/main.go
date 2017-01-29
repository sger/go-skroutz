package main

import (
	"fmt"

	skroutz "github.com/sger/go-skroutz"
)

func main() {
	resp, err := skroutz.Authorization("<clientID>", "<clientSecret>")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.AccessToken)
}
