package user

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	DefaultProfilePic = "default-profile-pic.jpg"
	DefaultColor      = "#000000"
)

type User struct {
	Name         string
	Surname      string
	ProfilePic   string
	ProfileColor string
	Experience   []string
}

func (u User) GetFullName() string {
	fullName := fmt.Sprintf("%s %s", u.Name, u.Surname)
	fullName = strings.TrimSpace(fullName)

	return fullName
}

func (u User) GetProfilePic() string {
	if u.ProfilePic != "" {
		return u.ProfilePic
	}

	return DefaultProfilePic
}

func (u User) GetProfileColor() string {
	if u.ProfileColor != "" {
		return u.ProfileColor
	}

	return DefaultColor
}

type Users []User

func (u Users) GetRandomUser() User {
	if len(u) == 0 {
		return User{}
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	randIndex := r.Intn(len(u))

	return u[randIndex]
}

func (u User) GetExperience() string {
	fullSummary := strings.Join(u.Experience, ",")
	return fullSummary
}
