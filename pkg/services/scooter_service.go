package services

import (
	"crudProj/model"
	"crudProj/pkg/repository"
)

type ScooterServiceI interface {
	CreateScooter(scooter *model.Scooter) (int, error)
	GetScooters() (*[]model.Scooter, error)
	GetScooterByID(scooterID int) (*model.Scooter, error)
	GetScootersByBrand(brand string) (*model.Scooter, error)
	EditScooter(scooter *model.Scooter) (int, error)
	DeleteScooter(id int) (int, error)
}

func NewScooterService(scooterRepository repository.ScooterRepositoryI) *ScooterService {
	return &ScooterService{
		scooterRepository,
	}
}

type ScooterService struct {
	scooterRepository repository.ScooterRepositoryI
}

func (s ScooterService) GetScooters() (*[]model.Scooter, error) {
	scooters, err := s.scooterRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return scooters, nil
}

func (s ScooterService) CreateScooter(scooter *model.Scooter) (int, error) {
	lastID, err := s.scooterRepository.Create(scooter)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (s ScooterService) GetScooterByID(scooterID int) (*model.Scooter, error) {
	scooter, err := s.scooterRepository.GetByID(scooterID)
	if err != nil {
		return nil, err
	}
	return scooter, nil
}

func (s ScooterService) GetScootersByBrand(brand string) (*model.Scooter, error) {
	scooter, err := s.scooterRepository.GetByBrand(brand)
	if err != nil {
		return nil, err
	}
	return scooter, nil
}

func (s ScooterService) EditScooter(scooter *model.Scooter) (int, error) {
	rowsAffected, err := s.scooterRepository.Update(scooter)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (s ScooterService) DeleteScooter(scooterID int) (int, error) {
	rowsAffected, err := s.scooterRepository.Delete(scooterID)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}