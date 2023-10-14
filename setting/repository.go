package setting

import (
	"encoding/json"
	"os"
)

const settingsDatabaseFileName = "settingsDatabase.json"

type Repository struct {
}

func NewRepository() Repository {
	repo := Repository{}

	return repo
}

func (r Repository) GetOne() (Settings, error) {
	settings, err := r.readFromFile()
	if err != nil {
		return Settings{}, err
	}

	return settings, nil
}

func (r Repository) Persist(setting Settings) (err error) {
	err = r.writeToFile(setting)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) readFromFile() (Settings, error) {
	dat, err := os.ReadFile(settingsDatabaseFileName)
	if err != nil {
		return Settings{}, err
	}

	setting := Settings{}

	err = json.Unmarshal(dat, &setting)
	if err != nil {
		return Settings{}, err
	}

	return setting, nil
}

func (r Repository) writeToFile(setting Settings) error {
	dat, err := json.Marshal(setting)
	if err != nil {
		return err
	}

	err = os.WriteFile(settingsDatabaseFileName, dat, 0640)
	if err != nil {
		return err
	}

	return nil
}
