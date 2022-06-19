package main

import "datatypes/organization"

func main() {
	p := organization.NewPerson("Collin", "Milner")
	println(p.ID())
	println(p.FullName())
}
