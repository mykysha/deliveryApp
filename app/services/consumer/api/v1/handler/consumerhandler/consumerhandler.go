package consumerhandler

import (
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/consumerapi"

	"github.com/nndergunov/deliveryApp/app/services/consumer/pkg/service/consumerservice"
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

	c.serveMux.HandleFunc("/v1/consumers", c.insertNewConsumer).Methods(http.MethodPost)
	c.serveMux.HandleFunc("/v1/consumers", c.getAllConsumer).Methods(http.MethodGet)
	c.serveMux.HandleFunc("/v1/consumers/{"+consumerIDKey+"}", c.deleteConsumer).Methods(http.MethodDelete)
	c.serveMux.HandleFunc("/v1/consumers/{"+consumerIDKey+"}", c.updateConsumer).Methods(http.MethodPut)
	c.serveMux.HandleFunc("/v1/consumers/{"+consumerIDKey+"}", c.getConsumer).Methods(http.MethodGet)

	c.serveMux.HandleFunc("/v1/locations/{"+consumerIDKey+"}", c.insertNewConsumerLocation).Methods(http.MethodPost)
	c.serveMux.HandleFunc("/v1/locations/{"+consumerIDKey+"}", c.updateConsumerLocation).Methods(http.MethodPut)
	c.serveMux.HandleFunc("/v1/locations/{"+consumerIDKey+"}", c.getConsumerLocation).Methods(http.MethodGet)
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

// swagger:operation POST /consumers insertNewConsumer
//
// Returns created consumer
//
// ---
// produces:
// - application/json
// parameters:
// - name: Body
//   in: body
//   description: cosnumer data
//   schema:
//     $ref: "#/definitions/NewConsumerRequest"
//   required: true
// responses:
//   '200':
//     description: created consumer
//     schema:
//       $ref: "#/definitions/ConsumerResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
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
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

		c.log.Println(err)
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

// swagger:operation DELETE /consumers/{consumer_id} DeleteAccount
//
// Returns "consumer deleted"
//
// ---
// produces:
// - application/json
// parameters:
// - name: consumer_id
//   in: path
//   description: consumer_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: consumer deleted
//     type: string
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *consumerHandler) deleteConsumer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[consumerIDKey]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.consumerService.DeleteConsumer(id)
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

		c.log.Println(err)
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

// swagger:operation PUT /consumers updateConsumer
//
// Returns update consumer
//
// ---
// produces:
// - application/json
// parameters:
// - name: Body
//   in: body
//   description: cosnumer data
//   schema:
//     $ref: "#/definitions/UpdateConsumerRequest"
//   required: true
// responses:
//   '200':
//     description: consumer update
//     schema:
//       $ref: "#/definitions/ConsumerResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *consumerHandler) updateConsumer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[consumerIDKey]
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
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

		c.log.Println(err)
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

// swagger:operation GET /consumers getAllConsumer
//
// Returns get all consumer
//
// ---
// produces:
// - application/json
// parameters:
// - name: Body
//   in: body
//   description: costumer list data
//   schema:
//     $ref: "#/definitions/UpdateConsumerRequest"
//   required: true
// responses:
//   '200':
//     description: consumer update
//     schema:
//       $ref: "#/definitions/ReturnConsumerResponseList"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *consumerHandler) getAllConsumer(rw http.ResponseWriter, r *http.Request) {
	data, err := c.consumerService.GetAllConsumer()
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

		c.log.Println(err)
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

// swagger:operation GET /consumers/{consumer_id} getConsumer
//
// Returns "consumer"
//
// ---
// produces:
// - application/json
// parameters:
// - name: consumer_id
//   in: path
//   description: consumer_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: consumer
//     type: string
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *consumerHandler) getConsumer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[consumerIDKey]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.consumerService.GetConsumer(id)
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

		c.log.Println(err)
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

// swagger:operation POST /locations/{consumer_id} insertNewConsumerLocation
//
// Returns consumer location
//
// ---
// produces:
// - application/json
// parameters:
// - name: consumer_id
//   in: path
//   description: consumer_id
//   schema:
//     type: integer
//   required: true
// - name: Body
//   in: body
//   description: cosnumer data
//   schema:
//     $ref: "#/definitions/NewConsumerRequest"
//   required: true
// responses:
//   '200':
//     description: created consumer
//     schema:
//       $ref: "#/definitions/ConsumerResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *consumerHandler) insertNewConsumerLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumerID, ok := vars[consumerIDKey]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var locationRequest consumerapi.NewLocationRequest

	if err := consumerapi.BindJson(r, &locationRequest); err != nil {
		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData); err != nil {
			c.log.Println(err)
		}
		return
	}

	location := requestToNewLocation(&locationRequest)

	data, err := c.consumerService.InsertLocation(location, consumerID)
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}
	response := locationToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation PUT /locations/{consumer_id} updateConsumerLocation
//
// Returns consumer location
//
// ---
// produces:
// - application/json
// parameters:
// - name: consumer_id
//   in: path
//   description: consumer_id
//   schema:
//     type: integer
//   required: true
// - name: Body
//   in: body
//   description: cosnumer data
//   schema:
//     $ref: "#/definitions/UpdateLocationRequest"
//   required: true
// responses:
//   '200':
//     description: created consumer
//     schema:
//       $ref: "#/definitions/ConsumerResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *consumerHandler) updateConsumerLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consumerID, ok := vars[consumerIDKey]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var updateLocationRequest consumerapi.UpdateLocationRequest

	if err := consumerapi.BindJson(r, &updateLocationRequest); err != nil {
		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	location := requestToUpdateLocation(&updateLocationRequest)

	data, err := c.consumerService.UpdateLocation(location, consumerID)
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := locationToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation GET /locations/{consumer_id} getConsumerLocation
//
// Returns consumer location
//
// ---
// produces:
// - application/json
// parameters:
// - name: consumer_id
//   in: path
//   description: consumer_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: created consumer
//     schema:
//       $ref: "#/definitions/ConsumerResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *consumerHandler) getConsumerLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[consumerIDKey]
	if !ok {
		if err := consumerapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.consumerService.GetLocation(id)
	if err != nil {

		if errors.Is(err, systemErr) {
			if err := consumerapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}

		c.log.Println(err)
		if err := consumerapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}
	response := locationToResponse(*data)

	if err := consumerapi.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}
