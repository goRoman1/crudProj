package services

import (
	"crudProj/entities"
	"crudProj/pkg/repository"
)

type ScooterServiceI interface {
	CreateScooter(scooter *entities.Scooter) (int, error)
	GetScooters() (*[]entities.Scooter, error)
	GetScooterByID(scooterID int) (*entities.Scooter, error)
	GetScootersByBrand(brand string) (*entities.Scooter, error)
	EditScooter(scooter *entities.Scooter) (int, error)
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

func (s ScooterService) GetScooters() (*[]entities.Scooter, error) {
	scooters, err := s.scooterRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return scooters, nil
}

func (s ScooterService) CreateScooter(scooter *entities.Scooter) (int, error) {
	lastID, err := s.scooterRepository.Create(scooter)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (s ScooterService) GetScooterByID(scooterID int) (*entities.Scooter, error) {
	scooter, err := s.scooterRepository.GetByID(scooterID)
	if err != nil {
		return nil, err
	}
	return scooter, nil
}

func (s ScooterService) GetScootersByBrand(brand string) (*entities.Scooter, error) {
	scooter, err := s.scooterRepository.GetByBrand(brand)
	if err != nil {
		return nil, err
	}
	return scooter, nil
}

func (s ScooterService) EditScooter(scooter *entities.Scooter) (int, error) {
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