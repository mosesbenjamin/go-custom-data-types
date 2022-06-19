package main

import "datatypes/organization"

func main() {
	// p := organization.Person{}
	p := organization.Person{FirstName: "Ben", LastName: "Murray"}
	p.FirstName = "Chris"
	println(p.ID())
}
