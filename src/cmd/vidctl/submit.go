package vidctl

import (
	"fmt"
)

type Video struct {
	ID          string
	Title       string
	Description string
	Location    string
	Destination string
	EnableLLM   bool
}

func NewVideo(id, title, description string) *Video {
	return &Video{
		ID:          id,
		Title:       title,
		Description: description,
	}
}

func SubmitVideo(videoID string) error {
	return fmt.Errorf("not implemented")
}
