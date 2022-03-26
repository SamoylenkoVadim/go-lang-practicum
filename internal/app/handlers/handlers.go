package handlers

import (
	cfg "github.com/SamoylenkoVadim/golang-practicum/internal/app/configs"
	"github.com/SamoylenkoVadim/golang-practicum/internal/app/storage"
	utils "github.com/SamoylenkoVadim/golang-practicum/internal/app/utils"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"net/url"
)

type Handler struct {
	Storage *storage.Storage
}

func New(s *storage.Storage) *Handler {
	return &Handler{s}
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	defer utils.PanicCatcher(w, r)

	urlID := chi.URLParam(r, "id")
	shortLink, err := h.Storage.GetValueFromStorageByKey(urlID)
	if err != nil {
		http.Error(w, "Bad request: unknown ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", string(shortLink))
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	defer utils.PanicCatcher(w, r)

	if r.Body == http.NoBody {
		http.Error(w, "Bad request: no body in request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request: unreadable body", http.StatusBadRequest)
		return
	}

	_, err = url.ParseRequestURI(string(body))
	if err != nil {
		http.Error(w, "Bad request: invalid URL in body", http.StatusBadRequest)
		return
	}

	linkID := utils.RandStringBytes()
	err = h.Storage.SaveToStorage(linkID, string(body))
	if err != nil {
		http.Error(w, "Bad request: storage problem", http.StatusBadRequest)
		return
	}
	responseBody := cfg.BaseURL + "/" + linkID

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(responseBody))
}
