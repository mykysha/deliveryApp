package accounthandler

import (
	"github.com/gorilla/mux"
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"io"
	"net/http"

	"accounting/api/v1/accountingapi"
	"accounting/pkg/service/accountservice"
)

type Params struct {
	Logger         *logger.Logger
	AccountService accountservice.AccountService
}

// accountHandler is the entrypoint into our application
type accountHandler struct {
	serveMux *mux.Router
	log      *logger.Logger
	service  accountservice.AccountService
}

// NewAccountHandler returns new http multiplexer with configured endpoints.
func NewAccountHandler(p Params) *mux.Router {
	serveMux := mux.NewRouter()

	handler := accountHandler{
		serveMux: serveMux,
		log:      p.Logger,
		service:  p.AccountService,
	}

	handler.handlerInit()

	return handler.serveMux
}

// NewAccountHandler creates an accountHandler value that handle a set of routes for the application.
func (a *accountHandler) handlerInit() {

	version := "/v1"

	a.serveMux.HandleFunc(version+"/status", a.StatusHandler).Methods(http.MethodPost)

	a.serveMux.HandleFunc(version+"/accounts", a.InsertNewAccount).Methods(http.MethodPost)
	a.serveMux.HandleFunc(version+"/accounts/{"+accountIDKey+"}", a.GetAccount).Methods(http.MethodGet)
	a.serveMux.HandleFunc(version+"/accounts/{"+accountIDKey+"}", a.DeleteAccount).Methods(http.MethodDelete)
	a.serveMux.HandleFunc(version+"/transaction", a.DoTransaction).Methods(http.MethodPut)
}

const accountIDKey = "account_id"

func (a accountHandler) StatusHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	data := v1.Status{
		ServiceName: "accounting",
		IsUp:        "up",
	}

	status, err := v1.EncodeIndent(data, "", " ")
	if err != nil {
		a.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = io.WriteString(responseWriter, string(status))
	if err != nil {
		a.log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	a.log.Printf("gave status %s", data.IsUp)
}

func (a accountHandler) InsertNewAccount(rw http.ResponseWriter, r *http.Request) {
	var newAccountRequest accountingapi.NewAccountRequest

	if err := accountingapi.BindJson(r, &newAccountRequest); err != nil {
		a.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	account := requestToNewAccount(&newAccountRequest)

	data, err := a.service.InsertNewAccount(account)

	if err != nil && err == systemErr {
		a.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			a.log.Println(err)
		}
		return
	}

	if err != nil {
		a.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}
	response := accountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		a.log.Println(err)
		return
	}
}

func (a accountHandler) GetAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars[accountIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.service.GetAccount(id)

	if err != nil {
		if err == systemErr {
			a.log.Println(err)
			if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
				a.log.Println(err)
			}
			return
		}

		if err == errAccountNotFound {
			if err := accountingapi.Respond(rw, http.StatusOK, err.Error()); err != nil {
				a.log.Println(err)
			}
		}

		a.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	response := accountToResponse(*data)

	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		a.log.Println(err)
		return
	}
}

func (a accountHandler) DeleteAccount(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars[accountIDKey]
	if !ok {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errNoConsumerIDParam.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	data, err := a.service.DeleteAccount(id)

	if err != nil && err == systemErr {
		a.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			a.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	if err := accountingapi.Respond(rw, http.StatusOK, data); err != nil {
		a.log.Println(err)
		return
	}
}

func (a accountHandler) DoTransaction(rw http.ResponseWriter, r *http.Request) {

	var transactionRequest accountingapi.TransactionRequest

	if err := accountingapi.BindJson(r, &transactionRequest); err != nil {
		a.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusBadRequest, errIncorrectInputData.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	transaction := requestToTransaction(&transactionRequest)

	data, err := a.service.Transact(transaction)

	if err != nil && err == systemErr {
		a.log.Println(err)
		if err := accountingapi.Respond(rw, http.StatusInternalServerError, ""); err != nil {
			a.log.Println(err)
		}
		return
	}

	if err != nil {
		if err := accountingapi.Respond(rw, http.StatusBadRequest, err.Error()); err != nil {
			a.log.Println(err)
		}
		return
	}

	response := transactionToResponse(data)
	if err := accountingapi.Respond(rw, http.StatusOK, response); err != nil {
		a.log.Println(err)
		return
	}
}
