package setting

import (
	"fmt"
	"strconv"

	"github.com/kubestaff/web-helper/server"
	"gorm.io/gorm"
)

const MaxVideosCountOnMainPage = 100

type SavingSuccess struct {
	Message string
}

type Handler struct {
	DbConn *gorm.DB
}

func (h Handler) Read(inputs server.Input) (o server.Output) {
	repo := NewRepository(h.DbConn)
	settings, err := repo.GetOne()
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}

	if settings.ID == 0 {
		return server.Output{
			Data: server.JsonError{
				Error: "Settings not found",
				Code:  404,
			},
			Code: 404,
		}
	}

	return server.Output{
		Data: settings,
		Code: 200,
	}
}

type SettingsInput struct {
	AboutTitle            string
	AboutText             string
	VideosCountOnMainPage string
}

type ValidationError struct {
	Error   string
	Code    int
	FieldId string
}

func (h Handler) Persist(inputs server.Input) (o server.Output) {
	setting := SettingsInput{}

	err := inputs.Scan(&setting)
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  400,
			},
			Code: 400,
		}
	}

	videosCountInt, err := strconv.Atoi(setting.VideosCountOnMainPage)
	if err != nil {
		return server.Output{
			Data: ValidationError{
				Error:   fmt.Sprintf("invalid number provided for videos count: %s", setting.VideosCountOnMainPage),
				Code:    400,
				FieldId: "videosCountOnMainPageInput",
			},
			Code: 400,
		}
	}

	if videosCountInt < 0 {
		videosCountInt *= -1
	}

	if videosCountInt > MaxVideosCountOnMainPage {
		return server.Output{
			Data: ValidationError{
				Error:   fmt.Sprintf("too big number for videos count: %d, max limit is %d", videosCountInt, MaxVideosCountOnMainPage),
				Code:    400,
				FieldId: "videosCountOnMainPageInput",
			},
			Code: 400,
		}
	}

	settingToSave := Settings{
		AboutTitle: setting.AboutTitle,
		AboutText:  setting.AboutText,
	}
	settingToSave.VideosCountOnMainPage = uint(videosCountInt)

	repo := NewRepository(h.DbConn)
	err = repo.Persist(settingToSave)
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}

	return server.Output{
		Data: SavingSuccess{
			Message: "Your settings were successfully saved",
		},
		Code: 200,
	}
}

func (h Handler) Delete(inputs server.Input) (o server.Output) {
	repo := NewRepository(h.DbConn)

	settings, err := repo.GetOne()
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}

	if settings.ID == 0 {
		return server.Output{
			Data: server.JsonError{
				Error: "Settings not found",
				Code:  404,
			},
			Code: 404,
		}
	}

	err = repo.Delete()
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}

	return server.Output{
		Data: SavingSuccess{
			Message: "Your settings were successfully deleted",
		},
		Code: 200,
	}
}
