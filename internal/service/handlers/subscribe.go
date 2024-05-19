package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Waramoto/hryvnia-svc/internal/data"
	pg "github.com/Waramoto/hryvnia-svc/internal/data/postgres"
	"github.com/Waramoto/hryvnia-svc/internal/service/requests"
	"github.com/Waramoto/hryvnia-svc/internal/types"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request, err := requests.NewSubscribeRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = DB(r).Subscribers().Insert(data.Subscriber{
		Email:    request.Email,
		LastSend: time.Now(),
		Status:   types.StatusNotSent,
	})
	if err != nil {
		if errors.Is(err, pg.ErrAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Log(r).WithError(err).Error("failed to insert email to DB")
		return
	}

	w.WriteHeader(http.StatusOK)
}
