package user

import "fmt"

func Handle(inputs map[string]string) (filename string, placeholders map[string]string) {
	userRepository := NewRepository()

	userId, ok := inputs["id"]
	if !ok {
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

func HandleChange(inputs map[string]string) (filename string, placeholders map[string]string) {
	userRepository := NewRepository()
	id, ok := inputs["id"]

	message := ""

	if ok {
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

func handleCreate(repo Repository, inputs map[string]string) (id string, err error) {
	createdUser := createUserFromInput(inputs)

	return repo.CreateUser(createdUser)
}

func handleUpdate(repo Repository, id string, inputs map[string]string) (err error) {
	updatedUser := createUserFromInput(inputs)

	return repo.UpdateUser(id, updatedUser)
}

func createUserFromInput(inputs map[string]string) User {
	usr := User{}
	surname, ok := inputs["surname"]
	if ok && surname != "" {
		usr.Surname = surname
	}

	name, ok := inputs["name"]
	if ok {
		usr.Name = name
	}

	return usr
}
