package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlejandroAM91/athenadb/internal/app/server/adapters/web"
	"github.com/AlejandroAM91/athenadb/internal/app/server/core/model"
	"github.com/AlejandroAM91/athenadb/internal/app/server/core/services"
	"github.com/go-chi/chi/v5"
)

type storagemock struct{}

func (s storagemock) GetAll() ([]model.DB, error) {
	dblist := []model.DB{
		{Name: "El nombre del viento"},
	}
	return dblist, nil
}

func main() {
	dbsrv := services.NewDBSrv(&storagemock{})

	r := chi.NewRouter()
	r.Route("/admin", func(r chi.Router) {
		r.Route("/database", func(r chi.Router) {
			dbctrl := web.NewDBCtrl(dbsrv)
			r.Get("/", dbctrl.GetAll)
		})
	})

	s := http.Server{Addr: ":8080", Handler: r}
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Print("Http Server stopped unexpected")
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := s.Shutdown(context.Background()); err != nil {
		log.Println("Failed to shutdown http server")
	}
}
