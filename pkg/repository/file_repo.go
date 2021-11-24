package repository

import (
	"context"
	"crudProj/model"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type FileRepository struct {
	db *pgx.Conn
}

func NewFileRepository(db *pgx.Conn) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

type FileRepositoryI interface {
	Insert(user *model.Scooter) (int, error)
}

func (u FileRepository)Insert(scooter *model.ScooterParse, conn *pgx.Conn)  {
	if _, err := conn.Exec(context.Background(), "INSERT INTO scooters(id, name, weight) VALUES($1, $2, $3)",
		scooter.Id, scooter.Model, scooter.Brand, scooter.BatteryCapacity, scooter.MaxWeight, scooter.BatteryCapacity)
		err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return
	}
	fmt.Println("Insertion Successfull")
}
