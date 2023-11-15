package usecases

import (
	"os"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/domain"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/repository"
)

type EjolUseCase struct {
	repository *repository.EjolRepository
}

func NewEjolUseCase(er *repository.EjolRepository) *EjolUseCase {
	return &EjolUseCase{er}
}

func (eu *EjolUseCase) GetEjlogNFS(fn string) domain.Ejol {
	var out domain.Ejol
	content, err := os.ReadFile(fn)
	if err != nil {
		return out
	}

	out = domain.Ejol{
		Ejlog: string(content),
	}
	return out
}

func (eu *EjolUseCase) GetEjlogDB(fn string) (domain.Ejol, error) {
	var out domain.Ejol
	content, err := os.ReadFile(fn)
	if err != nil {
		return out, err
	}

	out = domain.Ejol{
		Ejlog: string(content),
	}
	return out, nil
}
