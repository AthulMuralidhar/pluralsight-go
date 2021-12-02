// package == the dir name for this file
package models

import (
	"fmt"
)

type User struct {
	Id        int
	Firstname string
	Lastname  string
}

var (
	users  []*User
	nextid = 1
)

func GetAllUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	// check if there an id field already in the passed struct
	// if u.Id == 0 {
	// 	return User{}, errors.New("has an exisitng id already")
	// }

	u.Id = nextid
	nextid += 1
	users = append(users, &u)

	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("user with id: %v not found", id)
}

func UpdateUser(u User) (User, error) {
	for _, user_item := range users {
		if user_item.Id == u.Id {
			user_item = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user with id: %v not found", u.Id)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with id: %v not found", id)
}
