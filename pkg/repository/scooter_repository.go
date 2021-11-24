package repository

import (
	"crudProj/model"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"time"
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
	Create(user *model.Scooter) (int, error)
	GetAll() (*[]model.Scooter, error)
	GetByBrand(brand string) (*model.Scooter, error)
	GetByID(id int) (*model.Scooter, error)
	Update(user *model.Scooter) (int, error)
	Delete(id int) (int, error)
}

func (s ScooterRepository) GetAll() (*[]model.Scooter, error) {
	var users []model.Scooter
	rows, err := u.db.Query("SELECT * FROM scooters WHERE deleted=false")

	if err != nil {
		return nil, err
	}
	scooter := model.Scooter{}
	for rows.Next() {
		err = rows.Scan(&scooter.Id, &scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.BatteryCapacity, &scooter.MaxWeight)
		if err != nil {
			return nil, err
		}
		users = append(users, scooter)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (s ScooterRepository) Create(scooter *model.Scooter) (int, error) {
	result, err := s.db.Exec("INSERT INTO scooters (id, model, brand, max_distance, battery_capacity, max_weight, created, updated) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		0, &scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.BatteryCapacity, &scooter.MaxWeight, time.Now(), time.Now())
	if err != nil {
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

func (s ScooterRepository) GetByBrand(brand string) (*model.Scooter, error) {
	scooter := model.Scooter{}
	rows, err := s.db.Query("SELECT * FROM scooters WHERE brand=$1", brand)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.BatteryCapacity, &scooter.MaxWeight, &scooter.Created, &scooter.Updated)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &scooter, nil
}

func (s ScooterRepository) GetByID(id int) (*model.Scooter, error) {
	scooter := model.Scooter{}
	rows, err := s.db.Query("SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&scooter.Id, &scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.BatteryCapacity, &scooter.MaxWeight, &scooter.Created, &scooter.Updated)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return &scooter, nil
}

func (s ScooterRepository) Update(scooter *model.Scooter) (int, error) {
	result, err := s.db.Exec("UPDATE scooters SET model=$1, barnd=$2,max_distance=$3 battery_capacity=$4,max_weight=$5, updated=$4 WHERE id=$6",
		&scooter.Model, &scooter.Brand, &scooter.MaxDistance, &scooter.BatteryCapacity, &scooter.MaxWeight, time.Now(), &scooter.Id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (s ScooterRepository) Delete(id int) (int, error) {
	result, err := s.db.Exec("UPDATE scooters SET deleted=true, updated=current_timestamp WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
