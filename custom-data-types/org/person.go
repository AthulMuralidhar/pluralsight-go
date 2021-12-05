package org

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// type alias
// this one copies the field and the methods of string
// type TwitterHandler =  string

// type declaration
// only fields are copied
type TwitterHandler string

func (th TwitterHandler) RedirectURL() string {
	handler := strings.TrimPrefix(string(th), "@")

	return fmt.Sprintf("https://www.twitter.com/%s", handler)
}

// interface
type Identify interface {
	// only funcs
	GetID() string
}
type Citizen interface {
	Identify
	Country() string
}

// types
type Name struct {
	first string
	last  string
}
type Person struct {
	Name    //embeded struct
	twitter TwitterHandler
	Citizen // embedded interface
}
type Employee struct {
	Name //embeded struct
}

// constructor
func NewPerson(first, last string, identy Citizen) Person {
	return Person{
		Name:    Name{first: first, last: last},
		Citizen: identy,
	}
}

type socialSecurity string

func NewSocialSecurity(v string) Citizen {
	return socialSecurity(v)
}
func (ssn socialSecurity) GetID() string {
	return string(ssn)
}
func (ssn socialSecurity) Country() string {
	return "US"
}

type euSecurity struct {
	id      string
	country string
}

func NewEUSecurity(id interface{}, country string) Citizen {

	switch v := id.(type) {
	case string:
		return euSecurity{
			id:      v,
			country: country,
		}
	case int:
		return euSecurity{
			id:      strconv.Itoa(v),
			country: country,
		}
	default:
		panic("invalid type")
	}
}

func (eui euSecurity) GetID() string {
	return eui.id
}
func (eui euSecurity) Country() string {
	return eui.country
}

// getters

//  read only ones can have value based returns
//  or they can just have a non ref based struct
// but i am adding the * here for consistency
func (p *Person) GetFirst() string {
	return p.first
}
func (p *Person) GetLast() string {
	return p.last
}
func (p *Person) GetTwitter() TwitterHandler {
	return p.twitter
}
func (p *Person) GetID() string {
	// implicit inherits interfaces
	// based on the signature of the method
	return p.Citizen.GetID()
}

// setters

// when the setters are called, a copy is implicitely made
// so, when setting, use a pointer ref or make a new object implicetly
func (p *Person) SetTwitter(handle TwitterHandler) error {
	if !strings.HasPrefix(string(handle), "@") {
		return errors.New("twitter handle must start with @")
	}

	p.twitter = handle
	return nil
}
