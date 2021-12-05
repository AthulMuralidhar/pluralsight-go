package org

import (
	"errors"
	"fmt"
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

type Identify interface {
	// only funcs
	GetID() string
}

type Person struct {
	first string
	last string
	twitter TwitterHandler
}

// constructor
func NewPerson(first, last string) Person {
	return Person{first: first, last: last}
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
	return "1"

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
