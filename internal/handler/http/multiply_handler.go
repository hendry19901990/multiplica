package http

import (
	"net/http"
	"strconv"
	 service "multiplica/internal/service"
	 "multiplica/internal/entity"

	 "github.com/go-chi/chi"
	 log "github.com/sirupsen/logrus"
)

func multiplyTwoNumbers(w http.ResponseWriter, r *http.Request) {
	aString := chi.URLParam(r, "numberA")
	bString := chi.URLParam(r, "numberB")

	log.WithFields(log.Fields{
        "NumberA": aString,
        "NumberB": bString,
    }).Info("Numbers received from client")

	numberA, erra := strconv.Atoi(aString)
	if erra != nil {
		http.Error(w, "numberA and numberB must be a integer", http.StatusBadRequest)
    return
	}

	numberB, errb := strconv.Atoi(bString)
	if errb != nil {
		http.Error(w, "numberA and numberB must be a integer", http.StatusBadRequest)
    return
	}

	n := service.Multiply(int32(numberA), int32(numberB))
	WriteResponse(w, &entity.MultiplyResponse{n})
}
