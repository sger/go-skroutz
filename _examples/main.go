package main

import (
	"fmt"

	skroutz "github.com/sger/go-skroutz"
)

func main() {
	/*resp, err := skroutz.Authorize("rr0I3iyGFbPJtU9uX0WiAQ==", "eWyBtPqOBDTlbdKHLhyIsVh74RH3UMY0s7lltUclZS5QASIb09kdPjdE6uHmpP48YKnYrKTaTm3DBiR6q0C2A==")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.AccessToken)*/

	s := skroutz.New(skroutz.NewConfig("kXtTu5K6padrRqf7kWt_w0aocamvtvaaFSngWF2UrAgsLg3p_UJdJcBKljoGHxXO5hemWTTWbTAUZbmMAcJCog=="))

	/*out, err := s.Categories.GetCategories()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", out.Categories)*/

	/*test, err := s.Categories.GetCategories()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", test)*/

	/*test, err := s.Search.Autocomplete("iph")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", test)*/

	/*filter := map[string]string{
		"q": "iphone",
	}*/

	/*sq := &skroutz.SearchQuery{
		//Q:           "iphone",
		IncludeMeta:     []string{"applied_filters"},
		FilterIDS:       []string{"6282"},
		ManufacturerIDS: []string{"28"},
	}*/

	/*
		test, err := s.SKUS.GetCategorySKUS(40, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v\n", test)*/

	test, err := s.Products.SearchProducts(452, "95F")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", test)

	/*u, err := s.User.GetUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", u)*/
}
