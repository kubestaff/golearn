package user

import (
	"fmt"

	"github.com/kubestaff/web-helper/server"
)

func Handle(inputs server.Input) (filename string, placeholders map[string]string) {
	userRepository := NewRepository()

	userId := inputs.Get("id")
	if userId == "" {
		variables := map[string]string{
			"%errorMessage%": "User not found",
		}
		return "html/error.html", variables
	}

	usr, err := userRepository.FindUser(userId)
	if err != nil {
		variables := map[string]string{
			"%errorMessage%": err.Error(),
		}
		return "html/error.html", variables
	} else {
		variables := map[string]string{
			"%name%":            usr.GetFullName(),
			"%profile-picture%": usr.GetProfilePic(),
			"%profileColor%":    usr.GetProfileColor(),
			"%experience%":      usr.GetExperience(),
			"%summary%":         usr.GetSummary(),
			"%education%":       usr.GetEducation(),
		}

		// we output the contents of index.html
		return "html/user.html", variables
	}
}

func HandleChange(inputs server.Input) (filename string, placeholders map[string]string) {
	userRepository := NewRepository()
	id := inputs.Get("id")

	message := ""

	if id != "" {
		err := handleUpdate(userRepository, id, inputs)
		if err != nil {
			variables := map[string]string{
				"%errorMessage%": err.Error(),
			}
			return "html/error.html", variables
		}
		message = fmt.Sprintf("Updated a user with id %q", id)
	} else {
		id, err := handleCreate(userRepository, inputs)
		if err != nil {
			variables := map[string]string{
				"%errorMessage%": err.Error(),
			}
			return "html/error.html", variables
		}

		message = fmt.Sprintf("Created a new user with id %q", id)
	}

	variables := map[string]string{
		"%message%": message,
	}

	return "html/userChange.html", variables
}

func handleCreate(repo Repository, inputs server.Input) (id string, err error) {
	createdUser := createUserFromInput(inputs)

	return repo.CreateUser(createdUser)
}

func handleUpdate(repo Repository, id string, inputs server.Input) (err error) {
	updatedUser := createUserFromInput(inputs)

	return repo.UpdateUser(id, updatedUser)
}

func createUserFromInput(inputs server.Input) User {
	usr := User{}
	surname := inputs.Get("surname")
	if surname != "" {
		usr.Surname = surname
	}

	name := inputs.Get("name")
	if name != "" {
		usr.Name = name
	}

	return usr
}
