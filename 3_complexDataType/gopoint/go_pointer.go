package gopoint

import "fmt"

type Group struct {
	Name string
	ID   int64
}

type User struct {
	Name   string
	Group  *Group
	Group1 Group
}

func Pointer() {
	var user User
	user.Name = "test_user"
	var group Group
	user.Group = &group
	ChangePointer(user)
	ChangeGroup(user)
	fmt.Printf("user : %+v, group: %+v, group1: %+v.", user, user.Group, user.Group1)
}

func ChangePointer(user User) {
	user.Group.Name = "test_name"
	user.Group.ID = int64(1)
}

func ChangeGroup(user User) {
	user.Group1.Name = "test_name"
	user.Group1.ID = int64(1)
}
