package services

import (
	"crudProj/model"
	"crudProj/pkg/repository"
	"encoding/csv"
	"fmt"
	"github.com/jszwec/csvutil"
	"io"
	"os"
	"strconv"
)

type FileServiceI interface {
	UploadFromFile(file *model.ScooterParse)
}

func NewFileService(fileRepository repository.FileRepositoryI) *FileService {
	return &FileService{
		fileRepository,
	}
}

type FileService struct {
	fileRepository repository.FileRepositoryI
}


func (f FileService) UploadFromFile(scooters []model.ScooterParse) (int, error) {

	for i := 0; i < len(scooters); i++ {
		fmt.Println("Widget Id: " + scooters[i].Brand)
		fmt.Println("Widget Name: " + scooters[i].Model)
		fmt.Println("Widget Weight: " + strconv.Itoa(scooters[i].MaxDistance))
		Insert(& scooters[i], conn)
	}

	lastID, err := f.fileRepository.Insert(scooters)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}


func Pars(path string)[]model.ScooterParse{
	csv_file, _ := os.Open(path)
	reader := csv.NewReader(csv_file)
	reader.Comma = ';'

	scooterHeader, _ := csvutil.Header(model.ScooterParse{}, "csv")
	dec, _ := csvutil.NewDecoder(reader, scooterHeader...)

	var scooters []model.ScooterParse
	for {
		var s model.ScooterParse
		if err := dec.Decode(&s); err == io.EOF {
			break
		}
		scooters = append(scooters, s)
	}

	return scooters
}
