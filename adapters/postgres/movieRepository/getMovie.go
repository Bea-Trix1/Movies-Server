package movierepository

import (
	"context"

	"github.com/Bea-Trix1/Movies-Server/core/domain"
	"github.com/Bea-Trix1/Movies-Server/core/dto"
)

func (repository repository) GetAllMovies(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{}

	err := repository.db.QueryRow(
		ctx,
		"SELECT * FROM movie",
		movie.ID,
	).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Gender,
		&movie.Year,
		&movie.Director.FirstName,
		&movie.Director.LastName,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repository repository) GetMovieById(req *dto.MovieRequest) (*domain.Movie, error) {
	ctx := context.Background()
	movie := domain.Movie{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO movie (title, year, gender, year, fistName, lastName)  VALUES ($1, $2, $3, $4, $5, $6) returning *",
		movie.Title,
		movie.Gender,
		movie.Year,
		movie.Director.FirstName,
		movie.Director.LastName,
	).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Gender,
		&movie.Year,
		&movie.Director.FirstName,
		&movie.Director.LastName,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}
