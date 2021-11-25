package server

import (
	"context"
	"crudProj/config"
	"crudProj/pgdb"
	"crudProj/pkg/handlers"
	"crudProj/pkg/repository"
	"crudProj/pkg/services"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func Run()error{
	cfg := config.Get()
	pgDB :=  pgdb.Dial(cfg)

	fileRepository := repository.NewFileRepository(pgDB)
	fileService := services.NewFileService(fileRepository)
	fileHandler := handlers.NewFileHandler(fileService)

	scooterRepository := repository.NewScooterRepository(pgDB)
	scooterService := services.NewScooterService(scooterRepository)
	scooterHandler := handlers.NewScooterHandler(scooterService)

	r := mux.NewRouter()
	r.HandleFunc("/upload", fileHandler.UploadFile).Methods("POST")
	r.HandleFunc("/test", fileHandler.Test).Methods("POST")

	r.HandleFunc("/create", scooterHandler.Create).Methods("POST")
	r.HandleFunc("/scooters", scooterHandler.GetAll).Methods("GET")
	r.HandleFunc("/scooter/{id}", scooterHandler.GetById).Methods("GET")
	r.HandleFunc("/scooter/{brand}", scooterHandler.GetByModel).Methods("GET")
	r.HandleFunc("/edit", scooterHandler.EditInfo).Methods("PUT")
	r.HandleFunc("/delete/{id}", scooterHandler.Delete).Methods("DELETE")

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	http.Handle("/", r)
	srv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shutting down")
	os.Exit(0)
	return nil

}

