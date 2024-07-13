package movierepository

import (
	"context"

	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (repository repository) UpdateMovie(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{}

	err := repository.db.QueryRow(
		ctx,
		"UPDATE movie SET title = ?, gender = ?, year = ?, director = ?",
		movie.Title,
		movie.Gender,
		movie.Year,
		movie.Director,
	).Scan(
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
