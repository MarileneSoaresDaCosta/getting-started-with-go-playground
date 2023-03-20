package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // we use pointer here to be able to manipulate them without copying / moving
	nextID = 1     // the compiler will imply type int
)
