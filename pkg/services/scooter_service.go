package services

import (
	"crudProj/model"
	"crudProj/pkg/repository"
)

type ScooterServiceI interface {
	CreateScooter(user *model.Scooter) (int, error)
	GetScooters() (*[]model.Scooter, error)
	GetScooterByID(userID int) (*model.Scooter, error)
	GetScootersByBrand(brand string) (*model.Scooter, error)
	EditScooter(user *model.Scooter) (int, error)
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
	users, err := s.scooterRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s ScooterService) CreateScooter(user *model.Scooter) (int, error) {
	lastID, err := s.scooterRepository.Create(user)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (s ScooterService) GetScooterByID(userID int) (*model.Scooter, error) {
	user, err := s.scooterRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s ScooterService) GetScootersByBrand(email string) (*model.Scooter, error) {
	user, err := s.scooterRepository.GetByBrand(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s ScooterService) EditScooter(user *model.Scooter) (int, error) {
	rowsAffected, err := s.scooterRepository.Update(user)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (s ScooterService) DeleteScooter(userID int) (int, error) {
	rowsAffected, err := s.scooterRepository.Delete(userID)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}