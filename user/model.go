package user

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	DefaultProfilePicture = "default.jpg"
	DefaultColor          = "#ffffff"
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
	return DefaultProfilePicture
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
	r := rand.New(s) // initialize local pseudorandom generator
	randIndex := r.Intn(len(u))

	return u[randIndex]
}

func (u []string) GetExperience() string {
	fullSummary := fmt.Sprintf("%s", u.Experience)
	fullSummary = strings.Join(fullSummary)
	return fullSummary
}
