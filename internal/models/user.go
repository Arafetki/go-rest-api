package models

type User struct {
	ID       string
	FullName string
	Email    string
}

var AnonymousUser = &User{}

// Check if a User instance is the AnonymousUser.
func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}
