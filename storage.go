package storage

import (
	"fmt"
	"time"
)

const layoutISO = "2006-01-02"

// User struct
type User struct {
	id        int
	firstName string
	lastName  string
	year      time.Time
}

func newUser(id int, firstName, lastName, year string) User {
	y, err := time.Parse(layoutISO, year)
	if err != nil {
		fmt.Println(err)
	}
	user := User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		year:      y,
	}
	return user
}

// MapUserStorage ...
type MapUserStorage struct {
	users map[int]User
}

// NewMapUserStorage ...
func NewMapUserStorage() *MapUserStorage {
	return &MapUserStorage{
		users: make(map[int]User),
	}
}

// GetUser ...
func (u *MapUserStorage) GetUser(id int) User {
	findUser, exists := u.users[id]

	if !exists {
		return User{}
	}

	return findUser
}

// Add ...
func (u *MapUserStorage) Add(firstName, lastName, year string) {
	id := len(u.users)

	user := newUser(id, firstName, lastName, year)

	u.users[user.id] = user
}

// FindFirstName ...
func (u *MapUserStorage) FindFirstName(firstName string) map[int]User {
	var users = make(map[int]User)

	for _, value := range u.users {
		if value.firstName == firstName {
			users[value.id] = value
		}
	}

	return users
}
