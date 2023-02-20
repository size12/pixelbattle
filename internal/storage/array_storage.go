package storage

import (
	"pixelBattle/internal/config"
	"pixelBattle/internal/entity"
	"sync"
)

type ArrayStorage struct {
	Field entity.Field
	Cfg   *config.Config
	*sync.RWMutex
}

func NewArrayStorage(cfg *config.Config) (*ArrayStorage, error) {
	size := cfg.FieldSize
	field := make([][]entity.Color, size)
	for i := 0; i < size; i++ {
		field[i] = make([]entity.Color, size)
	}

	s := ArrayStorage{Field: field, Cfg: cfg, RWMutex: &sync.RWMutex{}}
	return &s, nil
}

func (s *ArrayStorage) DrawDot(dot *entity.Dot) error {
	s.Lock()
	defer s.Unlock()

	if dot.X > s.Cfg.FieldSize || dot.Y > s.Cfg.FieldSize || dot.X < 0 || dot.Y < 0 {
		return ErrOutFieldBorder
	}

	s.Field[dot.Y-1][dot.X-1] = dot.Color

	return nil
}

func (s *ArrayStorage) GetField() (entity.Field, error) {
	s.RLock()
	defer s.RUnlock()

	return s.Field, nil
}
