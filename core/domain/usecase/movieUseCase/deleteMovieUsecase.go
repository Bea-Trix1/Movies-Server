package movieusecase

func (usecase usecase) DeleteMovie(id uint32) error {
	err := usecase.repository.DeleteMovie(id)

	if err != nil {
		return err
	}

	return nil
}
