package handlers

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/nndergunov/deliveryApp/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/pkg/logger"
	"github.com/nndergunov/deliveryApp/services/restaurant/api/v1/restaurantapi"
	"github.com/nndergunov/deliveryApp/services/restaurant/pkg/service"
)

const (
	restaurantIDKey = "restaurantID"
	menuIDKey       = "menuID"
)

type endpointHandler struct {
	service  service.AppService
	serveMux *mux.Router
	log      *logger.Logger
}

// NewEndpointHandler returns new http multiplexer with configured endpoints.
func NewEndpointHandler(serviceInstance service.AppService, log *logger.Logger) *mux.Router {
	serveMux := mux.NewRouter()

	handler := endpointHandler{
		service:  serviceInstance,
		serveMux: serveMux,
		log:      log,
	}

	handler.handlerInit()

	return handler.serveMux
}

func (e *endpointHandler) handlerInit() {
	e.serveMux.HandleFunc("/status", e.statusHandler)

	e.serveMux.HandleFunc("/v1/restaurants", e.returnRestaurantList).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/restaurants/{"+restaurantIDKey+"}", e.returnRestaurant).Methods(http.MethodGet)
	e.serveMux.HandleFunc("/v1/restaurants/{"+restaurantIDKey+"}/menu", e.returnMenu).Methods(http.MethodGet)

	e.serveMux.HandleFunc("/v1/admin/restaurants", e.createRestaurant).Methods(http.MethodPost)
	e.serveMux.HandleFunc("/v1/admin/restaurants/{"+restaurantIDKey+"}",
		e.updateRestaurant).Methods(http.MethodPut)
	e.serveMux.HandleFunc("/v1/admin/restaurants/{"+restaurantIDKey+"}",
		e.deleteRestaurant).Methods(http.MethodDelete)

	e.serveMux.HandleFunc(
		"/v1/admin/restaurants/{"+restaurantIDKey+"}/menu", e.createMenu).Methods(http.MethodPost)

	e.serveMux.HandleFunc(
		"/v1/admin/restaurants/{"+restaurantIDKey+"}/menu", e.addMenuItem).Methods(http.MethodPut)
	e.serveMux.HandleFunc(
		"/v1/admin/restaurants/{"+restaurantIDKey+"}/menu/{"+menuIDKey+"}",
		e.updateMenuItem).Methods(http.MethodPatch)
	e.serveMux.HandleFunc(
		"/v1/admin/restaurants/{"+restaurantIDKey+"}/menu/{"+menuIDKey+"}",
		e.deleteMenuItem).Methods(http.MethodDelete)
}

func (e endpointHandler) statusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "restaurant",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		e.handleError(err, responseWriter)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	e.log.Printf("gave status %s", data.IsUp)
}

// swagger:operation GET /restaurants returnRestaurantList
//
// # Returns restaurant list
//
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	  description: restaurant list
//	  schema:
//	    $ref: "#/definitions/ReturnRestaurantList"
func (e endpointHandler) returnRestaurantList(w http.ResponseWriter, _ *http.Request) {
	restaurants, err := e.service.ReturnAllRestaurants()
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := restaurantListToResponse(restaurants)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

// swagger:operation GET /restaurants/{id} returnRestaurant
//
// # Returns requested restaurant
//
// ---
// parameters:
//   - name: id
//     in: path
//     type: integer
//     required: true
//
// produces:
// - application/json
// responses:
//
//	'200':
//	  description: requested restaurtant
//	  schema:
//	    $ref: "#/definitions/ReturnRestaurant"
func (e endpointHandler) returnRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	restaurant, err := e.service.ReturnRestaurant(restaurantID)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := restaurantToResponse(*restaurant)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

// swagger:operation GET /restaurants/{id}/menu returnMenu
//
// # Returns menu of the requested restaurant
//
// ---
// parameters:
//   - name: id
//     in: path
//     type: integer
//     required: true
//
// produces:
// - application/json
// responses:
//
//	'200':
//	  description: menu of the requested restaurant
//	  schema:
//	    $ref: "#/definitions/ReturnMenu"
func (e endpointHandler) returnMenu(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menu, err := e.service.ReturnMenu(restaurantID)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := menuToResponse(*menu)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

// swagger:operation POST /admin/restaurants createRestaurant
//
// # Returns menu of the requested restaurant
//
// ---
// produces:
// - application/json
// parameters:
//   - name: Body
//     in: body
//     description: restaurant data
//     schema:
//     $ref: "#/definitions/RestaurantData"
//     required: true
//
// responses:
//
//	'200':
//	  description: created restaurant
//	  schema:
//	    $ref: "#/definitions/ReturnRestaurant"
func (e *endpointHandler) createRestaurant(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	restaurantData := new(restaurantapi.RestaurantData)

	err = v1.Decode(req, restaurantData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	rest := requestToRestaurant(0, restaurantData)

	createdRest, err := e.service.CreateNewRestaurant(rest)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := restaurantToResponse(*createdRest)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

// swagger:operation PUT /admin/restaurants/{id} updateRestaurant
//
// # Updates restaurant data
//
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     type: integer
//     required: true
//   - name: Body
//     in: body
//     description: updated restaurant data
//     schema:
//     $ref: "#/definitions/RestaurantData"
//     required: true
//
// responses:
//
//	'200':
//	  description: updated restaurant data
//	  schema:
//	    $ref: "#/definitions/ReturnRestaurant"
func (e *endpointHandler) updateRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	restaurantData := new(restaurantapi.RestaurantData)

	err = v1.Decode(req, restaurantData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	rest := requestToRestaurant(restaurantID, restaurantData)

	updatedRestaurant, err := e.service.UpdateRestaurant(rest)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := restaurantToResponse(*updatedRestaurant)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

// swagger:operation DELETE /admin/restaurants/{id} deleteRestaurant
//
// # Deletes restaurant data
//
// ---
// parameters:
//   - name: id
//     in: path
//     type: integer
//     required: true
//
// responses:
//
//	'200':
func (e *endpointHandler) deleteRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	err = e.service.DeleteRestaurant(restaurantID)
	if err != nil {
		e.handleError(err, w)

		return
	}

	err = v1.Respond(nil, w)
	if err != nil {
		e.log.Println(err)
	}
}

// swagger:operation POST /admin/restaurants/{id}/menu createMenu
//
// # Creates menu in the restaurant
//
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     type: integer
//     required: true
//   - name: Body
//     in: body
//     description: menu data
//     schema:
//     $ref: "#/definitions/MenuData"
//     required: true
//
// responses:
//
//	'200':
//	  description: created menu
//	  schema:
//	    $ref: "#/definitions/ReturnMenu"
func (e *endpointHandler) createMenu(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuData := new(restaurantapi.MenuData)

	err = v1.Decode(req, menuData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menu := requestToMenu(restaurantID, menuData)

	createdMenu, err := e.service.CreateMenu(menu)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := menuToResponse(*createdMenu)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

// swagger:operation PUT /admin/restaurants/{id}/menu addMenuItem
//
// # Adds new menu item in the restaurant
//
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     type: integer
//     required: true
//   - name: Body
//     in: body
//     description: menu item data
//     schema:
//     $ref: "#/definitions/MenuItemData"
//     required: true
//
// responses:
//
//	'200':
//	  description: created menu item
//	  schema:
//	    $ref: "#/definitions/ReturnMenuItem"
func (e *endpointHandler) addMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItemData := new(restaurantapi.MenuItemData)

	err = v1.Decode(req, menuItemData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItem := requestToMenuItem(0, menuItemData)

	addedMenuItem, err := e.service.AddMenuItem(restaurantID, menuItem)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := menuItemToResponse(*addedMenuItem)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

// swagger:operation PATCH /admin/restaurants/{id}/menu/{itemid} updateMenuItem
//
// # Updates menu item in the restaurant
//
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     type: integer
//     required: true
//   - name: itemid
//     in: path
//     type: integer
//     required: true
//   - name: Body
//     in: body
//     description: updated menu item data
//     schema:
//     $ref: "#/definitions/MenuItemData"
//     required: true
//
// responses:
//
//	'200':
//	  description: updated menu item
//	  schema:
//	    $ref: "#/definitions/ReturnMenuItem"
func (e *endpointHandler) updateMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItemID, err := getIDFromEndpoint(menuIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItemData := new(restaurantapi.MenuItemData)

	err = v1.Decode(req, menuItemData)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItem := requestToMenuItem(menuItemID, menuItemData)

	updatedMenuItem, err := e.service.UpdateMenuItem(restaurantID, menuItem)
	if err != nil {
		e.handleError(err, w)

		return
	}

	response := menuItemToResponse(*updatedMenuItem)

	err = v1.Respond(response, w)
	if err != nil {
		e.log.Println(err)
	}
}

// swagger:operation DELETE /admin/restaurants/{id}/menu/{itemid} deleteMenuItem
//
// # Deletes menu item in the restaurant
//
// ---
// parameters:
//   - name: id
//     in: path
//     type: integer
//     required: true
//   - name: itemid
//     in: path
//     type: integer
//     required: true
//
// responses:
//
//	'200':
func (e *endpointHandler) deleteMenuItem(w http.ResponseWriter, r *http.Request) {
	restaurantID, err := getIDFromEndpoint(restaurantIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	menuItemID, err := getIDFromEndpoint(menuIDKey, r)
	if err != nil {
		e.handleError(err, w)

		return
	}

	err = e.service.DeleteMenuItem(restaurantID, menuItemID)
	if err != nil {
		e.handleError(err, w)

		return
	}

	w.WriteHeader(http.StatusOK)
}
