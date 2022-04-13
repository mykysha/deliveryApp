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
	e.mux.HandleFunc("/v1/order", e.orderHandler)
	e.mux.HandleFunc("/v1/order/status", e.orderStatusHandler)
	e.mux.HandleFunc("/v1/courier", e.courierHandler)
	e.mux.HandleFunc("/v1/courier/info", e.courierInfoHandler)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "kitchen",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		e.log.Println(err)
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		e.log.Printf("status write: %v", err)

		return
	}

	e.log.Printf("gave status %s", data.IsUp)
}

func (e endpointHandler) orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}

func (e endpointHandler) orderStatusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		// TODO logic.
	case http.MethodGet:
		// TODO logic.
	default:
		// TODO return error "unsupported method".
	}
}

func (e endpointHandler) courierHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}

func (e endpointHandler) courierInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// TODO return error "unsupported method".
	}

	// TODO logic.
}