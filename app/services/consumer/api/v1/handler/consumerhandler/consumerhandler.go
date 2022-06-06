// Package handler contains a small handler framework extension.
package consumerhandler

import (
	"consumer/api/v1/consumerapi"
	"consumer/pkg/service/consumerservice"
	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"io"
	"net/http"
)

type Params struct {
	Logger          *logger.Logger
	ConsumerService consumerservice.ConsumerService
}

// consumerHandler is the entrypoint into our application
type consumerHandler struct {
	serveMux        *mux.Router
	log             *logger.Logger
	consumerService consumerservice.ConsumerService
}

// NewConsumerHandler returns new http multiplexer with configured endpoints.
func NewConsumerHandler(p Params) *mux.Router {
	serveMux := mux.NewRouter()

	handler := consumerHandler{
		serveMux:        serveMux,
		log:             p.Logger,
		consumerService: p.ConsumerService,
	}

	handler.handlerInit()

	return handler.serveMux
}

const consumerIDKey = "consumer_id"

// NewConsumerHandler creates an consumerHandler value that handle a set of routes for the application.
func (c *consumerHandler) handlerInit() {

	c.serveMux.HandleFunc("/status", c.insertNewConsumer).Methods(http.MethodPost)

	c.serveMux.HandleFunc("/v1/consumer", c.insertNewConsumer).Methods(http.MethodPost)
	c.serveMux.HandleFunc("/v1/consumer/all", c.getAllConsumer).Methods(http.MethodGet)
	c.serveMux.HandleFunc("/v1/consumer/{"+consumerIDKey+"}", c.deleteConsumer).Methods(http.MethodDelete)
	c.serveMux.HandleFunc("/v1/consumer/{"+consumerIDKey+"}", c.updateConsumer).Methods(http.MethodPut)
	c.serveMux.HandleFunc("/v1/consumer/{"+consumerIDKey+"}", c.getConsumer).Methods(http.MethodGet)

	c.serveMux.HandleFunc("/v1/consumer/location/{"+consumerIDKey+"}", c.insertNewConsumerLocation).Methods(http.MethodPost)
	c.serveMux.HandleFunc("/v1/consumer/location/{"+consumerIDKey+"}", c.updateConsumerLocation).Methods(http.MethodPut)
	c.serveMux.HandleFunc("/v1/consumer/location/{"+consumerIDKey+"}", c.getConsumerLocation).Methods(http.MethodGet)
}

func (c *consumerHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "consumer",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		c.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		c.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	c.log.Printf("gave status %s", data.IsUp)
}

func (c *consumerHandler) insertNewConsumer(rw http.ResponseWriter, r *http.Request) {
	var consumerRequest consumerapi.NewConsumerRequest

	if err := consumerapi.BindJson(r, &consumerRequest); err != nil {
		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	consumer := requestToNewConsumer(&consumerRequest)

	data, err := c.consumerService.InsertConsumer(consumer)
	if err != nil && err == systemErr {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}

}

func (c *consumerHandler) deleteConsumer(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["consumer_id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.consumerService.DeleteConsumer(id)
	if err != nil && err == systemErr {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err := consumerapi.Respond(rw, http.StatusOK, data); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *consumerHandler) updateConsumer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["consumer_id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var updateConsumerRequest consumerapi.UpdateConsumerRequest

	if err := consumerapi.BindJson(r, &updateConsumerRequest); err != nil {
		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	consumer := requestToUpdateConsumer(&updateConsumerRequest)

	data, err := c.consumerService.UpdateConsumer(consumer, id)

	if err != nil && err == systemErr {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *consumerHandler) getAllConsumer(rw http.ResponseWriter, r *http.Request) {
	data, err := c.consumerService.GetAllConsumer()
	if err != nil && err == systemErr {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerListToResponse(data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *consumerHandler) getConsumer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["consumer_id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.consumerService.GetConsumer(id)

	if err != nil && err == systemErr {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *consumerHandler) insertNewConsumerLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumerID, ok := vars["consumer_id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var consumerLocationRequest consumerapi.NewConsumerLocationRequest

	if err := consumerapi.BindJson(r, &consumerLocationRequest); err != nil {
		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, "incorrect input data"); err != nil {
			c.log.Println(err)
		}
		return
	}

	consumerLocation := requestToNewConsumerLocation(&consumerLocationRequest)

	data, err := c.consumerService.InsertConsumerLocation(consumerLocation, consumerID)
	if err != nil && err == systemErr {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerLocationToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}

}

func (c *consumerHandler) updateConsumerLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumerID, ok := vars["consumer_id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var updateConsumerLocationRequest consumerapi.UpdateConsumerLocationRequest

	if err := consumerapi.BindJson(r, &updateConsumerLocationRequest); err != nil {
		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	consumerLocation := requestToUpdateConsumerLocation(&updateConsumerLocationRequest)

	data, err := c.consumerService.UpdateConsumerLocation(consumerLocation, consumerID)

	if err != nil && err == systemErr {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerLocationToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

func (c *consumerHandler) getConsumerLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["consumer_id"]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.consumerService.GetConsumerLocation(id)

	if err != nil && err == systemErr {
		if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := consumerLocationToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}