package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"pixelBattle/internal/entity"
	"pixelBattle/internal/storage"
	"time"

	"golang.org/x/net/websocket"
)

func NewClearHandler(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.ClearField()
		if err != nil {
			log.Println("Failed clear field: ", err)
			http.Error(w, "Failed clear field", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

func GetReceiveMessagesChan(ws *websocket.Conn) chan []byte {
	var message []byte
	result := make(chan []byte, 2)

	go func() {
		for {
			err := websocket.Message.Receive(ws, &message)
			if err != nil {
				ws.Close()
				return
			}
			result <- message
		}
	}()

	return result

}

func NewDrawHandler(s storage.Storage) websocket.Handler {
	return func(ws *websocket.Conn) {
		cfg := s.GetConfig()
		ticker := time.NewTicker(cfg.UpdateTime)

		messages := GetReceiveMessagesChan(ws)
		defer close(messages)

		for {
			select {
			case <-ticker.C:
				field, err := s.GetField()
				if err != nil {
					log.Println("Failed get field: ", err)
					ws.Write([]byte("Failed get field"))
					ws.Close()
					return
				}

				b, err := json.Marshal(&field)
				if err != nil {
					log.Println("Failed marshal field: ", err)
					ws.Write([]byte("Failed marshal field"))
					ws.Close()
					return
				}

				_, err = ws.Write(b)

				if err != nil {
					log.Println("Failed write to websocket: ", err)
					return
				}
			case message := <-messages:
				dot := entity.Dot{}
				err := json.Unmarshal(message, &dot)

				if err != nil {
					ws.Write([]byte("400"))
					continue
				}

				err = s.DrawDot(&dot)
				if err != nil {
					ws.Write([]byte("Error: " + err.Error()))
					continue
				}
			}
		}
	}
}
