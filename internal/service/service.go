package service

import (
	"log"
	"net/http"
	"net/url"
	"pixelBattle/internal/config"
	"pixelBattle/internal/handlers"
	"pixelBattle/internal/storage"

	"github.com/go-chi/chi/v5"
	"golang.org/x/net/websocket"
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

	URL, err := url.Parse("http://localhost:8080")

	if err != nil {
		log.Fatalln("Failed parse url: ", err)
	}

	config := websocket.Config{Origin: URL}

	r.Handle("/draw", websocket.Server{Handler: handlers.NewDrawHandler(s), Config: config})
	r.Get("/clear", handlers.NewClearHandler(s))
	r.MethodNotAllowedHandler()

	return server.ListenAndServe()
}
