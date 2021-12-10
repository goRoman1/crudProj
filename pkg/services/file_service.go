package services

import (
	"crudProj/entities"
	"crudProj/pkg/repository"
	"mime/multipart"
)

type FileServiceI interface {
	InsertScootersToDb(file multipart.File) string
	TestService(scooter *entities.Test) error
}

func NewFileService(fileRepository repository.FileRepositoryI) *FileService {
	return &FileService{
		fileRepository,
	}
}

type FileService struct {
	fileRepository repository.FileRepositoryI
}

func (f FileService) TestService(scooter *entities.Test) error {
	err:= f.fileRepository.Test(scooter)
	if err != nil {
		return err
	}
	return  err
}

func (f FileService)InsertScootersToDb(file multipart.File)string{
	tempFilePath := f.fileRepository.CreateTempFile(file)
	convertedStruct := f.fileRepository.ConvertToStruct(tempFilePath)

	err := f.fileRepository.InsertToDb(convertedStruct)
	if err != nil {
		return "Cant insert to DB"
	}

	return tempFilePath
}