package main

import (
	cfg "github.com/SamoylenkoVadim/golang-practicum/internal/app/configs"
	"github.com/SamoylenkoVadim/golang-practicum/internal/app/handlers"
	"github.com/SamoylenkoVadim/golang-practicum/internal/app/routers"
	"github.com/SamoylenkoVadim/golang-practicum/internal/app/storage"
	"net/http"
)

func main() {
	storage := storage.New()
	handlers, _ := handlers.New(storage)
	router, _ := routers.NewRouter(handlers)
	http.ListenAndServe(cfg.AddressPort, router)
}
