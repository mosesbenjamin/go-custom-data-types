package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// type Handler struct {
// 	handle string
// 	name   string
// }

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

// be careful with interface{}
func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      v,
			country: country,
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(v),
			country: country,
		}
	case europeanUnionIdentifier:
		return v
	case Person:
		euID, _ := v.Citizen.(europeanUnionIdentifier)
		return euID
	default:
		panic("using an invalid type to initialize EU identifier")
	}
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}
func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("Eu: %s", eui.country)
}

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
