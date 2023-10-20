package video

import (
		"github.com/kubestaff/web-helper/server"
)

func HandleList(inputs server.Input) (o server.Output) {
	provider := Provider{}
	videos, err := provider.FetchAllVideos()
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
		Data: videos,
		Code: 200,
	}
}
