package user

import (
	"fmt"
	"strings"
)

const (
	DefaultProfilePic = "static/img/default-profile-pic.jpg"
	DefaultColor      = "#000000"
)

type User struct {
	Name         string
	Surname      string
	ProfilePic   string
	ProfileColor string
	Experience   []string
	Summary      []string
	Education    []string
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

type Users map[string]User

func (u User) GetExperience() string {
	fullExperience := strings.Join(u.Experience, ",")
	return fullExperience
}

func (u User) GetSummary() string {
	fullSummary := strings.Join(u.Summary, ",")
	return fullSummary
}

func (u User) GetEducation() string {
	fullEducation := strings.Join(u.Education, ",")
	return fullEducation
}
