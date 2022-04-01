package api

import (
	"net/http"

	"github.com/nndergunov/deliveryApp/app/libs/logger"
)

type API struct {
	mux *http.ServeMux
	log *logger.Logger
}

func NewAPI(log *logger.Logger) *API {
	mux := http.NewServeMux()

	api := &API{
		mux: mux,
		log: log,
	}

	api.handlerInit()

	return api
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
