package handlers

import (
	"crudProj/model"
	"crudProj/pkg/services"
	"encoding/json"
	"net/http"
)

type ScooterHandler struct {
	scooterService services.ScooterServiceI
}

func NewScooterHandler(scooterService services.ScooterServiceI) *ScooterHandler {
	return &ScooterHandler{
		scooterService: scooterService,
	}
}

type ScooterHandlerI interface {
	CreateScooter(w http.ResponseWriter, r *http.Request)
	GetScooterById(w http.ResponseWriter, r *http.Request)
	GetScooterByEmail(w http.ResponseWriter, r *http.Request)
	EditScooter(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (u ScooterHandler) Create(w http.ResponseWriter, r *http.Request) {
	var scooter model.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = u.scooterService.UploadFromFile(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u ScooterHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var scooters []model.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = u.scooterService.GetScooters()
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u ScooterHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var scooter model.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_, err = u.scooterService.GetScooterByID(scooter.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u ScooterHandler) GetByModel(w http.ResponseWriter, r *http.Request) {
	var scooter model.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_, err = u.scooterService.GetScootersByBrand(scooter.Model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u ScooterHandler) EditInfo(w http.ResponseWriter, r *http.Request) {
	var scooter model.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = u.scooterService.EditScooter(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u ScooterHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var scooter model.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = u.scooterService.DeleteScooter(scooter.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}
