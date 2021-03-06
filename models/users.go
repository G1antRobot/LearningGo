package models

type User struct {
	// Name is capital because it is exportable.
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var (
	users  []*User // stores pointers
	nextID int     = 1
	// notice that these dont have to be used in program, go will not complain about this.
)
