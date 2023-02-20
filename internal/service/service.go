package service

import (
	"log"
	"net/http"
	"pixelBattle/internal/config"
	"pixelBattle/internal/handlers"
	"pixelBattle/internal/middleware"
	"pixelBattle/internal/storage"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	Cfg     *config.Config
	Storage storage.Storage
}

func NewService(cfg *config.Config) *Service {
	s, err := storage.NewStorage(cfg)

	if err != nil {
		log.Fatalln("Failed get storage: ", err)
	}

	return &Service{
		Cfg:     cfg,
		Storage: s,
	}
}

func (service *Service) Run() error {
	r := chi.NewRouter()
	s, err := storage.NewStorage(service.Cfg)

	if err != nil {
		log.Fatalln("Failed open storage:", err)
	}

	server := http.Server{Addr: service.Cfg.RunAddress, Handler: r}

	r.Group(func(r chi.Router) {
		r.Use(middleware.OnlyJSONContent)
		r.Post("/draw", handlers.NewDrawDotHandler(s))
	})

	r.Get("/draw", handlers.NewGetFieldHandler(s))
	r.MethodNotAllowedHandler()

	return server.ListenAndServe()
}
