package repository

import (
	"context"
	"crudProj/entities"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type ScooterRepository struct {
	db *pgx.Conn
}

func NewScooterRepository(db *pgx.Conn) *ScooterRepository {
	return &ScooterRepository{
		db: db,
	}
}

type ScooterRepositoryI interface {
	Create(scooter *entities.Scooter) (int, error)
	GetAll() (*[]entities.Scooter, error)
	GetByBrand(brand string) (*entities.Scooter, error)
	GetByID(id int) (*entities.Scooter, error)
	Update(scooter *entities.Scooter) (int, error)
	Delete(id int) (int, error)
}

func (s ScooterRepository) GetAll() (*[]entities.Scooter, error) {
	var scooters []entities.Scooter
	rows, err := s.db.Query(context.Background(),"SELECT * FROM scooters")

	if err != nil {
		return nil, err
	}
	scooter := entities.Scooter{}
	for rows.Next() {
		err = rows.Scan(&scooter.Id, &scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.Capacity, &scooter.MaxWeight)
		if err != nil {
			return nil, err
		}
		scooters = append(scooters, scooter)
	}

	return &scooters, nil
}

func (s ScooterRepository) Create(scooter *entities.Scooter) (int, error) {
	res, err := s.db.Exec(context.Background(),"INSERT INTO scooters (id, entities, brand, max_distance, capacity, max_weight) VALUES ($1, $2, $3, $4, $5, $6)",
		0, &scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.Capacity, &scooter.MaxWeight)
	if err != nil {
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	lastID := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

func (s ScooterRepository) GetByBrand(brand string) (*entities.Scooter, error) {
	scooter := entities.Scooter{}
	rows, err := s.db.Query(context.Background(), "SELECT * FROM scooters WHERE brand=$1", brand)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.Capacity, &scooter.MaxWeight)
		if err != nil {
			return nil, err
		}
	}

	return &scooter, nil
}

func (s ScooterRepository) GetByID(id int) (*entities.Scooter, error) {
	scooter := entities.Scooter{}
	rows, err := s.db.Query(context.Background(), "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&scooter.Id, &scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.Capacity, &scooter.MaxWeight)
		if err != nil {
			return nil, err
		}
	}

	return &scooter, nil
}

func (s ScooterRepository) Update(scooter *entities.Scooter) (int, error) {
	res, err := s.db.Exec(context.Background(), "UPDATE scooters SET entities=$1, barnd=$2,max_distance=$3 capacity=$4,max_weight=$5 WHERE id=$6",
		&scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.Capacity, &scooter.MaxWeight, &scooter.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := res.RowsAffected()
	return int(rowsAffected), nil
}

func (s ScooterRepository) Delete(id int) (int, error) {
	res, err := s.db.Exec(context.Background(), "`DELETE FROM scooters WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsAffected := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
