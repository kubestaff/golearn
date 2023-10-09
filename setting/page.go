package setting

import (
	"fmt"
	"strconv"

	"github.com/kubestaff/web-helper/server"
)

const MaxVideosCountOnMainPage = 100

type SavingSuccess struct {
	Message string
}

func HandleRead(inputs server.Input) (o server.Output) {
	repo := NewRepository()
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

	return server.Output{
		Data: settings,
		Code: 200,
	}
}

func HandlePersist(inputs server.Input) (o server.Output) {
	setting := Settings{
		AboutTitle: inputs.Get("about-title"),
		AboutText:  inputs.Get("about-text"),
	}

	videosCount := inputs.Get("videos-count")

	videosCountInt, err := strconv.Atoi(videosCount)
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: fmt.Sprintf("invalid number provided for videos count: %s", videosCount),
				Code:  400,
			},
			Code: 400,
		}
	}

	if videosCountInt < 0 {
		videosCountInt *= -1
	}

	if videosCountInt > MaxVideosCountOnMainPage {
		return server.Output{
			Data: server.JsonError{
				Error: fmt.Sprintf("too big number for videos count: %d, max limit is %d", videosCountInt, MaxVideosCountOnMainPage),
				Code:  400,
			},
			Code: 400,
		}
	}

	setting.VideosCountOnMainPage = uint(videosCountInt)

	repo := NewRepository()
	err = repo.Persist(setting)
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
