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
	"strings"
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
	valueStrings := make([]string, 0, len(scooters))
	valueArgs := make([]interface{}, 0, len(scooters) * 7)
	for i, scooter := range scooters {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*7+1, i*7+2, i*7+3, i*7+4, i*7+5, i*7+6, i*7+7))
		valueArgs = append(valueArgs, scooter.Id)
		valueArgs = append(valueArgs, scooter.Model)
		valueArgs = append(valueArgs, scooter.Brand)
		valueArgs = append(valueArgs, scooter.Capacity)
		valueArgs = append(valueArgs, scooter.MaxWeight)
		valueArgs = append(valueArgs, scooter.MaxDistance)
		valueArgs = append(valueArgs, scooter.Serial)
	}

	stmt := fmt.Sprintf("INSERT INTO scooters(id, model, brand, capacity, max_weight, max_distance, serial) VALUES %s", strings.Join(valueStrings, ","))
	if _, err := f.db.Exec(context.Background(),stmt, valueArgs...)
		err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return err
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