package storage

import (
	"pixelBattle/internal/config"
	"pixelBattle/internal/entity"
)

type Storage interface {
	DrawDot(dot *entity.Dot) error
	GetField() (entity.Field, error)
	ClearField() error
	GetConfig() *config.Config
}

func NewStorage(cfg *config.Config) (Storage, error) {
	s, err := NewArrayStorage(cfg)
	return s, err
}
