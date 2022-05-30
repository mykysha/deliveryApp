package consumerservice

import (
	"accounting/pkg/domain"
	"database/sql"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"strconv"
)

// ConsumerService is the interface for the accounting service.
type ConsumerService interface {
	InsertNewConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	GetConsumerAccount(consumerID string) (*domain.ConsumerAccount, error)
	DeleteConsumerAccount(consumerID string) (string, error)

	AddToBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
	SubFromBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	Storage ConsumerStorage
	Logger  *logger.Logger
}

type Service struct {
	Storage ConsumerStorage
	logger  *logger.Logger
}

// NewService constructs a new NewService.
func NewService(p Params) *Service {
	ServiceItem := &Service{
		Storage: p.Storage,
		logger:  p.Logger,
	}

	return ServiceItem
}

func (c Service) InsertNewConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {

	if account.ConsumerID < 1 {
		return nil, errWrongConsumerIDType
	}
	//check duplicate
	consumerAccount, err := c.Storage.GetConsumerAccountByID(account.ConsumerID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if consumerAccount != nil {
		return nil, errConsumerAccountExist
	}

	//insertNewConsumerAccount
	newConsumerAccount, err := c.Storage.InsertNewConsumerAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return newConsumerAccount, nil
}

func (c Service) GetConsumerAccount(consumerID string) (*domain.ConsumerAccount, error) {
	idUint, err := strconv.ParseUint(consumerID, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return nil, errWrongConsumerIDType
	}

	consumerAccount, err := c.Storage.GetConsumerAccountByID(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if consumerAccount == nil {
		return nil, errConsumerAccountNotFound
	}

	return consumerAccount, nil
}

func (c Service) DeleteConsumerAccount(consumerID string) (string, error) {
	idUint, err := strconv.ParseUint(consumerID, 10, 64)
	if err != nil {
		c.logger.Println(err)
		return "", errWrongConsumerIDType
	}

	consumerAccount, err := c.Storage.GetConsumerAccountByID(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}
	if consumerAccount != nil {
		return "", errConsumerAccountExist
	}

	err = c.Storage.DeleteConsumerAccount(idUint)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return "", systemErr
	}

	return "consumer account deleted", nil
}

func (c Service) AddToBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {

	if account.ConsumerID < 1 {
		return nil, errWrongConsumerIDType
	}

	if account.Balance < 1 {
		return nil, errWrongAmount
	}

	consumerAccount, err := c.Storage.GetConsumerAccountByID(account.ConsumerID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if consumerAccount == nil {
		return nil, errConsumerAccountNotFound
	}

	consumerAccountUpdated, err := c.Storage.AddToBalanceConsumerAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return consumerAccountUpdated, nil
}

func (c Service) SubFromBalanceConsumerAccount(account domain.ConsumerAccount) (*domain.ConsumerAccount, error) {

	if account.ConsumerID < 1 {
		return nil, errWrongConsumerIDType
	}

	if account.Balance < 1 {
		return nil, errWrongAmount
	}

	consumerAccount, err := c.Storage.GetConsumerAccountByID(account.ConsumerID)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}
	if consumerAccount == nil {
		return nil, errConsumerAccountNotFound
	}

	if consumerAccount.Balance < account.Balance {
		return nil, errNotEnoughBalance
	}

	consumerAccountUpdated, err := c.Storage.SubFromBalanceConsumerAccount(account)
	if err != nil && err != sql.ErrNoRows {
		c.logger.Println(err)
		return nil, systemErr
	}

	return consumerAccountUpdated, nil
}
