// package == the dir name for this file
package models

type User struct {
	Id int
	Firstname string
	Lastname string 
}

var (
	users []*User
	nextid = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.Id = nextid
	nextid +=1 
	users = append(users, &u)

	return u, nil 
}