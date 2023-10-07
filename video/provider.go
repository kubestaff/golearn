package video

import (
	"encoding/json"
	"os"
)

const videosDatabaseFileName = "videosDatabase.json"

type Provider struct {
}

func (p Provider) FetchAllVideos() ([]Video, error) {
	videos, err := p.readFromFile()
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (p Provider) readFromFile() ([]Video, error) {
	dat, err := os.ReadFile(videosDatabaseFileName)
	if err != nil {
		return nil, err
	}

	videos := []Video{}

	err = json.Unmarshal(dat, &videos)
	if err != nil {
		return nil, err
	}

	return videos, nil
}
