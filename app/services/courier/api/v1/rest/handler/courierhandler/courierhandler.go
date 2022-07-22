package courierhandler

import (
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	courierapi2 "github.com/nndergunov/deliveryApp/app/services/courier/api/v1/rest/courierapi"
	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/courier/pkg/service/courierservice"
)

type Params struct {
	Logger         *logger.Logger
	CourierService courierservice.CourierService
}

// courierHandler is the entrypoint into our application
type courierHandler struct {
	serveMux       *mux.Router
	log            *logger.Logger
	courierService courierservice.CourierService
}

// NewCourierHandler returns new http multiplexer with configured endpoints.
func NewCourierHandler(p Params) *mux.Router {
	serveMux := mux.NewRouter()

	handler := courierHandler{
		serveMux:       serveMux,
		log:            p.Logger,
		courierService: p.CourierService,
	}

	handler.handlerInit()

	return handler.serveMux
}

const courierIDKey = "courier_id"

// NewCourierHandler creates an courierHandler value that handle a set of routes for the application.
func (c *courierHandler) handlerInit() {
	c.serveMux.HandleFunc("/status", c.statusHandler).Methods(http.MethodPost)

	c.serveMux.HandleFunc("/v1/couriers", c.insertNewCourier).Methods(http.MethodPost)
	c.serveMux.HandleFunc("/v1/couriers", c.getCourierList).Methods(http.MethodGet)
	c.serveMux.HandleFunc("/v1/couriers/{"+courierIDKey+"}", c.deleteCourier).Methods(http.MethodDelete)
	c.serveMux.HandleFunc("/v1/couriers/{"+courierIDKey+"}", c.updateCourier).Methods(http.MethodPut)
	c.serveMux.HandleFunc("/v1/couriers/{"+courierIDKey+"}", c.getCourier).Methods(http.MethodGet)

	c.serveMux.HandleFunc("/v1/couriers-available/{"+courierIDKey+"}", c.updateCourierAvailable).Methods(http.MethodPut)

	c.serveMux.HandleFunc("/v1/locations", c.getLocationList).Methods(http.MethodGet)
	c.serveMux.HandleFunc("/v1/locations/{"+courierIDKey+"}", c.insertNewLocation).Methods(http.MethodPost)
	c.serveMux.HandleFunc("/v1/locations/{"+courierIDKey+"}", c.updateLocation).Methods(http.MethodPut)
	c.serveMux.HandleFunc("/v1/locations/{"+courierIDKey+"}", c.getLocation).Methods(http.MethodGet)
}

func (c *courierHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "courier",
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

// swagger:operation POST /couriers insertNewCourier
//
// Returns created courier
//
// ---
// produces:
// - application/json
// parameters:
// - name: Body
//   in: body
//   description: courier data
//   schema:
//     $ref: "#/definitions/NewCourierRequest"
//   required: true
// responses:
//   '200':
//     description: created courier
//     schema:
//       $ref: "#/definitions/CourierResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *courierHandler) insertNewCourier(rw http.ResponseWriter, r *http.Request) {
	var courierRequest courierapi2.NewCourierRequest

	if err := courierapi2.BindJson(r, &courierRequest); err != nil {
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	courier := requestToNewCourier(&courierRequest)

	data, err := c.courierService.InsertCourier(courier)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}
	response := courierToResponse(*data)

	if err := courierapi2.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation DELETE /couriers/{courier_id} DeleteAccount
//
// Returns "courier deleted"
//
// ---
// produces:
// - application/json
// parameters:
// - name: courier_id
//   in: path
//   description: courier_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: courier deleted
//     type: string
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *courierHandler) deleteCourier(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.courierService.DeleteCourier(id)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	if err := courierapi2.Respond(rw, http.StatusOK, data); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation PUT /couriers{courier_id} updateCourier
//
// Returns update courier
//
// ---
// produces:
// - application/json
// parameters:
// - name: Body
//   in: body
//   description: courier data
//   schema:
//     $ref: "#/definitions/UpdateCourierRequest"
//   required: true
// - name: courier_id
//   in: path
//   description: courier_id
//   schema:
//     type: integer
//   required: true

// responses:
//   '200':
//     description: courier update
//     schema:
//       $ref: "#/definitions/CourierResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *courierHandler) updateCourier(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var updateCourierRequest courierapi2.UpdateCourierRequest

	if err := courierapi2.BindJson(r, &updateCourierRequest); err != nil {
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	courier := requestToUpdateCourier(&updateCourierRequest)

	data, err := c.courierService.UpdateCourier(courier, id)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierToResponse(*data)

	if err := courierapi2.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation PUT /couriers-available/{courier_id} updateCourierAvailable
//
// Returns update courier
//
// ---
// produces:
// - application/json
// parameters:
// - name: available
//   in: query
//   description: courier data
//   schema:
//     type: bool
//   required: true
// - name: courier_id
//   in: path
//   description: courier_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: courier update
//     schema:
//       $ref: "#/definitions/CourierResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *courierHandler) updateCourierAvailable(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	queryParams := r.URL.Query()
	queryParamsList := queryParams["available"]

	if queryParamsList == nil {
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errNoAvailableParam.Error()); err != nil {
			c.log.Println(err)
		}
	}
	available := queryParamsList[0]

	data, err := c.courierService.UpdateCourierAvailable(id, available)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierToResponse(*data)
	if err := courierapi2.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation GET /couriers getCourierList
//
// Returns get all couriers
//
// ---
// produces:
// - application/json
// parameters:
// - name: Body
//   in: body
//   description: courier list data
//   schema:
//     $ref: "#/definitions/UpdateCourierRequest"
//   required: true
// - name: available
//   in: query
//   description: courier data
//   schema:
//     type: bool
//   required: true
// responses:
//   '200':
//     description: courier update
//     schema:
//       $ref: "#/definitions/CourierResponseList"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *courierHandler) getCourierList(rw http.ResponseWriter, r *http.Request) {
	param := domain.SearchParam{}

	queryParams := r.URL.Query()

	availableList := queryParams["available"]
	if availableList != nil {
		available := availableList[0]
		if available != "" {
			param["available"] = available
		}
	}

	data, err := c.courierService.GetCourierList(param)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierListToResponse(data)

	if err := courierapi2.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation GET /couriers/{courier_id} getCourier
//
// Returns "courier"
//
// ---
// produces:
// - application/json
// parameters:
// - name: courier_id
//   in: path
//   description: courier_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: courier
//     type: string
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *courierHandler) getCourier(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errNoCourierIDParam); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.courierService.GetCourier(id)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := courierToResponse(*data)

	if err := courierapi2.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation POST /locations/{courier_id} insertNewCourierLocation
//
// Returns courier location
//
// ---
// produces:
// - application/json
// parameters:
// - name: courier_id
//   in: path
//   description: courier_id
//   schema:
//     type: integer
//   required: true
// - name: Body
//   in: body
//   description: location data
//   schema:
//     $ref: "#/definitions/NewLocationRequest"
//   required: true
// responses:
//   '200':
//     description: created courier
//     schema:
//       $ref: "#/definitions/LocationResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *courierHandler) insertNewLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courierID, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var locationRequest courierapi2.NewLocationRequest

	if err := courierapi2.BindJson(r, &locationRequest); err != nil {
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	location := requestToNewLocation(&locationRequest)

	data, err := c.courierService.InsertLocation(location, courierID)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := locationToResponse(*data)

	if err := courierapi2.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation PUT /locations/{courier_id} updateLocation
//
// Returns courier location
//
// ---
// produces:
// - application/json
// parameters:
// - name: courier_id
//   in: path
//   description: courier_id
//   schema:
//     type: integer
//   required: true
// - name: Body
//   in: body
//   description: location data
//   schema:
//     $ref: "#/definitions/UpdateLocationRequest"
//   required: true
// responses:
//   '200':
//     description: location updated
//     schema:
//       $ref: "#/definitions/LocationResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
//       type: string
func (c *courierHandler) updateLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courierID, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	var updateLocationRequest courierapi2.UpdateLocationRequest

	if err := courierapi2.BindJson(r, &updateLocationRequest); err != nil {
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	location := requestToUpdateLocation(&updateLocationRequest)

	data, err := c.courierService.UpdateLocation(location, courierID)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := locationToResponse(*data)

	if err := courierapi2.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation GET /locations/{courier_id} getLocation
//
// Returns courier location
//
// ---
// produces:
// - application/json
// parameters:
// - name: courier_id
//   in: path
//   description: courier_id
//   schema:
//     type: integer
//   required: true
// responses:
//   '200':
//     description: created location
//     schema:
//       $ref: "#/definitions/LocationResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
func (c *courierHandler) getLocation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, ok := vars[courierIDKey]
	if !ok {
		if err := courierapi2.Respond(rw, http.StatusBadRequest, errNoCourierIDParam.Error()); err != nil {
			c.log.Println(err)
		}
	}

	data, err := c.courierService.GetLocation(userID)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := locationToResponse(*data)

	if err := courierapi2.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}

// swagger:operation GET /locations/{courier_id} getLocationList
//
// Returns courier location
//
// ---
// produces:
// - application/json
// parameters:
// - name: city
//   in: query
//   description: city
//   schema:
//     type: string
//   required: false
// responses:
//   '200':
//     description: created location
//     schema:
//       $ref: "#/definitions/LocationResponse"
//   '500':
//     description: internal server error
//     schema:
//       type: string
//   '400':
//     description: bad request
//     schema:
func (c *courierHandler) getLocationList(rw http.ResponseWriter, r *http.Request) {
	param := domain.SearchParam{}
	queryParams := r.URL.Query()

	cityList := queryParams["city"]
	if cityList != nil {
		city := cityList[0]
		if city != "" {
			param["city"] = city
		}
	}

	data, err := c.courierService.GetLocationList(param)
	if err != nil {
		if errors.Is(err, systemErr) {
			c.log.Println(err)
			if err := courierapi2.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				c.log.Println(err)
			}
			return
		}
		c.log.Println(err)
		if err := courierapi2.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			c.log.Println(err)
		}
		return
	}

	response := locationListToResponse(data)

	if err := courierapi2.Respond(rw, http.StatusOK, response); err != nil {
		c.log.Println(err)
		return
	}
}