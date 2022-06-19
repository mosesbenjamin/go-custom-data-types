package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Collin", "Milner")
	err := p.SetTwitterHandler("@morrison")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s", err.Error())
	}
	println(p.ID())
	println(p.FullName())
}
