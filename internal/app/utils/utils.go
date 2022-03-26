package utils

import (
	cfg "github.com/SamoylenkoVadim/golang-practicum/internal/app/configs"
	"math/rand"
	"net/http"
	"time"
)

func PanicCatcher(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		http.Error(w, "Bad request: panic happened", http.StatusBadRequest)
		return
	}
}

func RandStringBytes() string {
	shortlink := make([]byte, cfg.BaseLength)
	rand.Seed(time.Now().UnixNano())
	for i := range shortlink {
		shortlink[i] = cfg.LetterBytes[rand.Intn(len(cfg.LetterBytes))]
	}
	return string(shortlink)
}
