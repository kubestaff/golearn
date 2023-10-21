package setting

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	dbConn *gorm.DB
}

func NewRepository(dbConn *gorm.DB) Repository {
	repo := Repository{
		dbConn: dbConn,
	}

	return repo
}

func (r Repository) GetOne() (Settings, error) {
	settn := Settings{}
	r.dbConn.First(&settn)

	return settn, nil
}

func (r Repository) Persist(setting Settings) (err error) {
	existingSettings, err := r.GetOne()
	if err != nil {
		return err
	}

	if existingSettings.ID == 0 {
		r.dbConn.Create(&setting)
		return nil
	}

	r.dbConn.Model(&existingSettings).Updates(setting)

	return nil
}

func (r Repository) Delete() (err error) {
	existingSettings, err := r.GetOne()
	if err != nil {
		return err
	}

	if existingSettings.ID == 0 {
		return errors.New("no setting found to be deleted")
	}

	settn := Settings{}
	r.dbConn.Unscoped().Delete(&settn, existingSettings.ID)

	return nil
}
