package main

import "datatypes/organization"

func main() {
	// p := organization.Person{}
	var p organization.Identifiable = organization.Person{}
	println(p.ID())
}
