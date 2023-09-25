package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const userDatabaseFileName = "usersDatabase.json"

type Repository struct {
}

func NewRepository() Repository {
	repo := Repository{}

	return repo
}

// crud - create a user, read a user, update a user and delete a user
// crud - create a video, read a video/videos, update a video, delete a video
func (r Repository) FindUser(id string) (User, error) {
	users, err := r.readFromFile()
	if err != nil {
		return User{}, err
	}

	user, ok := users[id]
	if !ok {
		return User{}, errors.New("no user found")
	}

	return user, nil
}

func (r Repository) CreateUser(usr User) (id string, err error) {
	if usr.Surname == "" {
		return "", errors.New("can't create a user with an empty surname")
	}

	users := Users{}

	id = strings.ToLower(usr.Surname)
	usr, ok := users[id]
	if ok {
		return "", fmt.Errorf("we have another user with same surname %q", usr.Surname)
	}

	users[id] = usr

	err = r.writeToFile(users)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r Repository) UpdateUser(id string, inputUser User) (err error) {
	users, err := r.readFromFile()
	if err != nil {
		return err
	}

	_, ok := users[id]
	if !ok {
		return fmt.Errorf("no user exists for id %q", id)
	}

	users[id] = inputUser
	err = r.writeToFile(users)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) DeleteUser(id string) (err error) {
	users, err := r.readFromFile()
	if err != nil {
		return err
	}

	_, ok := users[id]
	if !ok {
		return fmt.Errorf("no user exists for id %q", id)
	}

	delete(users, id)

	err = r.writeToFile(users)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) readFromFile() (Users, error) {
	dat, err := os.ReadFile(userDatabaseFileName)
	if err != nil {
		return Users{}, err
	}

	mapOfUsers := Users{}

	err = json.Unmarshal(dat, &mapOfUsers)
	if err != nil {
		return Users{}, err
	}

	return mapOfUsers, nil
}

func (r Repository) writeToFile(usrs Users) error {
	dat, err := json.Marshal(usrs)
	if err != nil {
		return err
	}

	err = os.WriteFile(userDatabaseFileName, dat, 0644)
	if err != nil {
		return err
	}

	return nil
}
