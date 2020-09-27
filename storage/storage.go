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
	Users map[int]User
}

// NewMapUserStorage ...
func NewMapUserStorage() *MapUserStorage {
	return &MapUserStorage{
		Users: make(map[int]User),
	}
}

// GetUser ...
func (u *MapUserStorage) GetUser(Id int) User {
	findUser, exists := u.Users[Id]

	if !exists {
		return User{}
	}

	return findUser
}

// Add ...
func (u *MapUserStorage) Add(FirstName, LastName, Year string) {
	Id := len(u.Users)

	user := NewUser(Id, FirstName, LastName, Year)

	u.Users[user.Id] = user
}

// FindFirstName ...
func (u *MapUserStorage) FindFirstName(FirstName string) map[int]User {
	var users = make(map[int]User)

	for _, value := range u.Users {
		if value.FirstName == FirstName {
			users[value.Id] = value
		}
	}

	return users
}
