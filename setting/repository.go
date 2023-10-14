package setting

import (
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
	return nil
}
