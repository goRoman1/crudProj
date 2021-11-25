package repository

import (
	"context"
	"crudProj/model"
	"encoding/csv"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jszwec/csvutil"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"sync"
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
	Test(scooter *model.Test)error
	Insert(scooter *model.Scooter) error
	InsertToDb(scooters []model.Scooter) error
	CreateTempFile(file multipart.File)string
	ConvertToStruct(path string)[]model.Scooter
}


func (f FileRepository) CreateTempFile(file multipart.File)string{
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	//./../internal/temp_files
	tempFile, err := ioutil.TempFile("./", "upload-*.сsv")
	if err != nil {
		fmt.Println(err)
	}
//	defer tempFile.Close()
	tempFile.Write(fileBytes)
	return tempFile.Name()
}

func (f FileRepository) ConvertToStruct(path string)[]model.Scooter {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(csvFile)
	reader.Comma = ';'

	scooterHeader, _ := csvutil.Header(model.Scooter{}, "csv")
	dec, _ := csvutil.NewDecoder(reader, scooterHeader...)

	var scooters []model.Scooter
	for {
		var s model.Scooter
		if err := dec.Decode(&s); err == io.EOF {
			break
		}
		scooters = append(scooters, s)
	}
	return scooters
}

func (f FileRepository)Insert(scooter *model.Scooter) error {
	if _, err := f.db.Exec(context.Background(),
		"INSERT INTO scooters(model, brand, capacity, max_weight, max_distance, serial) VALUES($1, $2, $3, $4, $5, $6, $7)",
		scooter.Id,scooter.Model, scooter.Brand, scooter.Capacity, scooter.MaxWeight, scooter.MaxDistance, scooter.Serial)
		err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return err
	}
	fmt.Println("Insertion Successfull")
	return nil
}

func (f FileRepository) InsertToDb(scooters []model.Scooter) error {
	var wg sync.WaitGroup
	wg.Add(len(scooters))
	defer wg.Done()
	for i := 0; i < len(scooters); i++ {
		fmt.Println("Widget Brand: " + scooters[i].Brand)
		fmt.Println("Widget Model: " + scooters[i].Model)
		fmt.Println("Widget MaxDistance: " + strconv.Itoa(scooters[i].MaxDistance))
		go f.Insert(&scooters[i])
	}
	return nil
}

func (f FileRepository)Test(scooter *model.Test) error {
	if _, err := f.db.Exec(context.Background(),
		"INSERT INTO test(id, model, brand) VALUES($1, $2, $3)",
		scooter.Id,scooter.Model, scooter.Brand)
		err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return err
	}
	fmt.Println("Insertion Successfull")
	return nil
}