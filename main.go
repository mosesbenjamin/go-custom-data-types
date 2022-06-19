package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Collin", "Milner", organization.NewEuropeanUnionIdentifier("123-456-789", "Germany"))
	err := p.SetTwitterHandler("@morrison")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s", err.Error())
	}

	// name1 := Name{First: "Ben", Last: "Smith"}
	// name2 := otherName{First: "Ben2", Last: "Smith"}

	ssn := organization.NewSocialSecurityNumber("123-444-222")
	eu := organization.NewEuropeanUnionIdentifier("123-444-222", "France")
	eu2 := organization.NewEuropeanUnionIdentifier("123-444-222", "France")

	fmt.Printf("%T\n", ssn)
	fmt.Printf("%T\n", eu)

	if eu == eu2 {
		fmt.Println("We match")
	}
}

type Name struct {
	First string
	Last  string
	// Middle []string
	// Middle func
	// Middle map[string]string
}

type otherName struct {
	First string
	Last  string
}
