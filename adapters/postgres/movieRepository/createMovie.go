package movierepository

import (
	"context"

	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (repository repository) CreateMovie(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO movie (title, year, gender, year, fistName, lastName)  VALUES ($1, $2, $3, $4, $5, $6) returning *",
		movie.Title,
		movie.Gender,
		movie.Year,
		movie.Director,
	).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Gender,
		&movie.Year,
		&movie.Director,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}
