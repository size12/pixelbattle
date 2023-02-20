package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pixelBattle/internal/entity"
	"pixelBattle/internal/storage"
)

func NewDrawDotHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resBody, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, "wrong body: "+err.Error(), http.StatusBadRequest)
			return
		}

		dot := entity.Dot{}
		err = json.Unmarshal(resBody, &dot)

		if err != nil {
			http.Error(w, "Failed unmarshal your request.", http.StatusBadRequest)
			return
		}

		err = s.DrawDot(&dot)

		if errors.Is(err, storage.ErrOutFieldBorder) {
			http.Error(w, "Bad coordinates. Failed draw dot", http.StatusUnprocessableEntity)
			return
		}

		if err != nil {
			http.Error(w, "Failed draw dot.", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		return
	}
}

func NewGetFieldHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		field, err := s.GetField()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(field)))
		return
	}
}
