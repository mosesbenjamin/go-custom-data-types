package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Collin", "Milner", organization.NewEuropeanUnionIdentifier(123789, "Germany"))
	err := p.SetTwitterHandler("@morrison")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s", err.Error())
	}

	println(p.ID())
	println(p.Country())
}
