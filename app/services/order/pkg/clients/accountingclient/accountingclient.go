// Package accountingclient is responsible for communication with accounting service.
package accountingclient

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// AccountingClient communicates with accounting service.
type AccountingClient struct {
	accountingURL string
}

// NewAccountingClient returns new AccountingClient instance.
func NewAccountingClient(url string) *AccountingClient {
	return &AccountingClient{accountingURL: url}
}

// CreateTransaction sends transaction data to the accounting service.
func (a AccountingClient) CreateTransaction(accountID, restaurantID int, orderPrice float64) (bool, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(a.accountingURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, fmt.Errorf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	c := pb.NewAccountingClient(conn)

	ctx := context.TODO()

	_, err = c.InsertTransaction(ctx, &pb.TransactionRequest{
		FromAccountID: int64(accountID),
		ToAccountID:   int64(restaurantID),
		Amount:        orderPrice,
	})

	if err != nil {
		return false, fmt.Errorf("got error : %w", err)
	}

	return true, nil
}
