package storage

import (
	"fmt"
	"time"
)

const layoutISO = "2006-01-02"

// User struct
type User struct {
	Id        int
	FirstName string
	LastName  string
	Year      time.Time
}

// NewUser ...
func NewUser(Id int, FirstName, LastName, Year string) User {
	y, err := time.Parse(layoutISO, Year)
	if err != nil {
		fmt.Println(err)
	}
	user := User{
		Id:        Id,
		FirstName: FirstName,
		LastName:  LastName,
		Year:      y,
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
func (u *MapUserStorage) GetUser(Id int) User {
	findUser, exists := u.users[Id]

	if !exists {
		return User{}
	}

	return findUser
}

// Add ...
func (u *MapUserStorage) Add(FirstName, LastName, Year string) {
	Id := len(u.users)

	user := newUser(Id, FirstName, LastName, Year)

	u.users[user.Id] = user
}

// FindFirstName ...
func (u *MapUserStorage) FindFirstName(FirstName string) map[int]User {
	var users = make(map[int]User)

	for _, value := range u.users {
		if value.FirstName == FirstName {
			users[value.Id] = value
		}
	}

	return users
}
