package handlers

import (
	"io"
	"net/http"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
)

type endpointHandler struct {
	mux *http.ServeMux
	log *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(log *logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	handler := endpointHandler{
		mux: mux,
		log: log,
	}

	handler.handlerInit()

	return handler.mux
}

func (e *endpointHandler) handlerInit() {
	e.mux.HandleFunc("/v1/status", e.statusHandler)
	e.mux.HandleFunc("/v1/payment", e.paymentHandler)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "accounting",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		e.log.Println(err)
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		e.log.Printf("\nstatus write: %v", err)

		return
	}

	e.log.Printf("\ngave status %s", data.IsUp)
}

func (e endpointHandler) paymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}