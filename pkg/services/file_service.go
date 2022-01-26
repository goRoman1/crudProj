package services

import (
	"crudProj/pkg/repository"
	"mime/multipart"
)

type FileServiceI interface {
	InsertScootersToDb(file multipart.File) string
}

func NewFileService(fileRepository repository.FileRepositoryI) *FileService {
	return &FileService{
		fileRepository,
	}
}

type FileService struct {
	fileRepository repository.FileRepositoryI
}


func (f FileService)InsertScootersToDb(file multipart.File)string{
	tempFilePath := f.fileRepository.CreateTempFile(file)
	f.fileRepository.ConvertToStruct(tempFilePath)

/*
	err := f.fileRepository.InsertToDb(convertedStruct)
	if err != nil {
		return "Cant insert to DB"
	}
 */
	return tempFilePath
}