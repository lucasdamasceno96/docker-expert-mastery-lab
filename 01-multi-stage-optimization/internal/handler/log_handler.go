package handler

import (
	"encoding/json"
	"minimalrestapi/internal/entity"
	"minimalrestapi/internal/repository"
	"net/http"
)

type LogHandler struct {
	Repo *repository.LogRepository
}

func (h *LogHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *LogHandler) CreateLog(w http.ResponseWriter, r *http.Request) {
	log := &entity.Log{
		ID:   "log-" + r.RemoteAddr,
		Data: "request_received",
	}

	if err := h.Repo.Save(log); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "saved", "id": log.ID})
}
