package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	city "main/internal/city"
	create "main/internal/handlers/create"
	delete "main/internal/handlers/delete"
	getByDistrict "main/internal/handlers/get_by_district"
	getByFoundation "main/internal/handlers/get_by_foundation_diapazon"
	getById "main/internal/handlers/get_by_id"
	getByPopulation "main/internal/handlers/get_by_population_diapazon"
	getByRegion "main/internal/handlers/get_by_region"
	population "main/internal/handlers/patch"
	"main/internal/storage"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

type App struct {
	router *chi.Mux
	done   chan os.Signal
	store  map[int]city.City
}

func NewApp() *App {
	ret := &App{
		router: chi.NewRouter(),
		done:   make(chan os.Signal, 1),
	}
	signal.Notify(ret.done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	return ret
}

func (a *App) run() {
	bd := storage.LoadFileRAM(filepath.Join("internal", "data", "cities.csv"), storage.NewMainStorage())
	//a.router.Post("/user", createUser.Create)
	a.router.Route("/city", func(r chi.Router) {
		r.Post("/", create.New(bd))
		r.Get("/population", getByPopulation.New(bd))
		r.Get("/foundation", getByFoundation.New(bd))
		r.Route("/{cityId}", func(r chi.Router) {
			r.Get("/", getById.New(bd))
			r.Delete("/", delete.New(bd))
			r.Patch("/population", population.New(bd))

		})
		r.Route("/region", func(r chi.Router) {
			r.Get("/{region}", getByRegion.New(bd))
		})
		r.Route("/district", func(r chi.Router) {
			r.Get("/{district}", getByDistrict.New(bd))
		})
	})

	go func() {
		fmt.Println("Starting worker")
		log.Fatal(http.ListenAndServe(":8001", a.router))

	}()
	<-a.done
	fmt.Println("Exiting")
	storage.SaveFile(filepath.Join("internal", "data", "cities.csv"), bd)
	fmt.Println("Succesful writing")

}

func main() {
	var app = NewApp()
	app.run()

}
