package http

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"encoding/json"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"multiplica/cmd/config"
)

func NewHTTPHandler(configuration config.Configurations) {
	log.Info("Api server started at ", configuration.Server.Port)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(middleware.Logger)
	r.Get("/api/v1/multiply/{numberA}/{numberB}", multiplyTwoNumbers)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Server.Port), r))
}


func WriteResponse(w http.ResponseWriter, resp interface{}){
	bts, _ := json.Marshal(resp)
	w.Write(bts)
	return
}
