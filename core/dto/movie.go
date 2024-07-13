package dto

import (
	"encoding/json"
	"io"
)

type MovieRequest struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Gender   string `json:"gender"`
	Year     string `json:"year"`
	Director string `json:"director"`
}

func JSONMovieRequest(body io.Reader) (*MovieRequest, error) {
	req := MovieRequest{}
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
