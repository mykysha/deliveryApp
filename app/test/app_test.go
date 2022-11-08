package test

import (
	"github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/services/accounting/api/v1/rest/accountingapi"
	"github.com/nndergunov/deliveryApp/app/services/consumer/api/v1/rest/consumerapi"
	"github.com/nndergunov/deliveryApp/app/services/courier/api/v1/rest/courierapi"
	"github.com/nndergunov/deliveryApp/app/services/delivery/api/v1/rest/deliveryapi"
	"github.com/nndergunov/deliveryApp/app/services/order/api/v1/orderapi"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/api/v1/restaurantapi"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"bytes"
	"net/http"
	"os"
	"strconv"
	"testing"
)

const configFile = "/config.yaml"

func TestAppSuccess(t *testing.T) {
	confPath, err := os.Getwd()
	require.NoError(t, err)

	err = configreader.SetConfigFile(confPath + configFile)
	require.NoError(t, err)
	consumerURL := configreader.GetString("service.local.consumer")
	courierURL := configreader.GetString("service.local.courier")
	accountingURL := configreader.GetString("service.local.accounting")
	deliveryURL := configreader.GetString("service.local.delivery")
	//kitchenURL := configreader.GetString("service.local.kitchen")
	orderURL := configreader.GetString("service.local.order")
	restaurantURL := configreader.GetString("service.local.restaurant")

	//consumer
	consumerID := 0
	{
		consumer := consumerapi.NewConsumerRequest{
			Firstname: "consumer1FirstName",
			Lastname:  "consumer1lastName",
			Email:     "consumer1@email.com",
			Phone:     "123456789",
		}
		reqBody, err := v1.Encode(consumer)
		require.NoError(t, err)

		resp, err := http.Post(consumerURL+"/v1/consumers", "application/json", bytes.NewBuffer(reqBody))
		require.NoError(t, err)

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := consumerapi.ConsumerResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)

		consumerID = respData.ID
		consumerResp := consumerapi.ConsumerResponse{
			ID:        consumerID,
			Firstname: consumer.Firstname,
			Lastname:  consumer.Lastname,
			Email:     consumer.Email,
			Phone:     consumer.Phone,
		}
		assert.Equal(t, consumerResp, respData)
	}
	defer func() {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodDelete, consumerURL+"/v1/consumers/"+strconv.Itoa(consumerID), nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}
	}()

	//consumer location
	{
		consumerLoc := consumerapi.NewLocationRequest{
			Latitude:  "41.035100",
			Longitude: "28.862448",
			Country:   "Turkey",
			City:      "Istanbul",
		}
		consumerLocResp := consumerapi.LocationResponse{
			UserID:    consumerID,
			Latitude:  consumerLoc.Latitude,
			Longitude: consumerLoc.Longitude,
			Country:   consumerLoc.Country,
			City:      consumerLoc.City,
		}

		reqBody, err := v1.Encode(consumerLoc)
		require.NoError(t, err)

		resp, err := http.Post(consumerURL+"/v1/locations/"+strconv.Itoa(consumerID), "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := consumerapi.LocationResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		assert.Equal(t, consumerLocResp, respData)
	}

	//courier
	courierID := 0
	{
		//courier
		courier := courierapi.NewCourierRequest{
			Firstname: "courier1FirstName",
			Lastname:  "courier1lastName",
			Username:  "courier1Username",
			Password:  "courier1Password",
			Email:     "courier1@email.com",
			Phone:     "123456789",
		}

		reqBody, err := v1.Encode(courier)
		require.NoError(t, err)

		resp, err := http.Post(courierURL+"/v1/couriers", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := courierapi.CourierResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		courierID = respData.ID
		courierResp := courierapi.CourierResponse{
			ID:        courierID,
			Username:  courier.Username,
			Firstname: courier.Firstname,
			Lastname:  courier.Lastname,
			Email:     courier.Email,
			Phone:     courier.Phone,
			Available: true,
		}
		assert.Equal(t, courierResp, respData)
	}

	defer func() {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodDelete, courierURL+"/v1/couriers/"+strconv.Itoa(courierID), nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}
	}()

	//courier location
	{
		courierLoc := courierapi.NewLocationRequest{
			Latitude:  "41.015148",
			Longitude: "29.177157",
			Country:   "Turkey",
			City:      "Istanbul",
		}
		courierLocResp := courierapi.LocationResponse{
			UserID:    courierID,
			Latitude:  courierLoc.Latitude,
			Longitude: courierLoc.Longitude,
			Country:   courierLoc.Country,
			City:      courierLoc.City,
		}
		reqBody, err := v1.Encode(courierLoc)
		require.NoError(t, err)

		resp, err := http.Post(courierURL+"/v1/locations/"+strconv.Itoa(courierID), "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := courierapi.LocationResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		assert.Equal(t, courierLocResp, respData)
	}

	//restaurant
	restaurantID := 0
	{
		//restaurant
		restaurant := restaurantapi.RestaurantData{
			Name:            "TestRestaurant1",
			City:            "Istanbul",
			AcceptingOrders: true,
			Address:         "Gebze",
			Longitude:       40.797099,
			Latitude:        29.438401,
		}

		reqBody, err := v1.Encode(restaurant)
		require.NoError(t, err)

		resp, err := http.Post(restaurantURL+"/v1/admin/restaurants", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := restaurantapi.ReturnRestaurant{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		restaurantID = respData.ID
		restaurantResp := restaurantapi.ReturnRestaurant{
			ID:              restaurantID,
			Name:            restaurant.Name,
			AcceptingOrders: restaurant.AcceptingOrders,
			City:            restaurant.City,
			Address:         restaurant.Address,
			Longitude:       restaurant.Longitude,
			Latitude:        restaurant.Latitude,
		}
		assert.Equal(t, restaurantResp, respData)

	}

	defer func() {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodDelete, restaurantURL+"/v1/admin/restaurants/"+strconv.Itoa(courierID), nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}
	}()

	//restaurant menu
	{
		restaurantMenu := restaurantapi.MenuData{MenuItems: []restaurantapi.MenuItemData{
			restaurantapi.MenuItemData{
				ID:     1,
				Name:   "testMenu1",
				Price:  10,
				Course: "test1",
			}},
		}
		//restaurantMenuResp := restaurantapi.ReturnMenu{
		//	RestaurantID: restaurantID,
		//	MenuItems: []restaurantapi.ReturnMenuItem{restaurantapi.ReturnMenuItem{
		//		ID:     1,
		//		Name:   "testMenu1",
		//		Price:  10,
		//		Course: "test1",
		//	}},
		//}
		reqBody, err := v1.Encode(restaurantMenu)
		require.NoError(t, err)

		resp, err := http.Post(restaurantURL+"/v1/admin/restaurants/"+strconv.Itoa(restaurantID)+"/menu", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := restaurantapi.ReturnMenu{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
	}

	//account consumer
	consumerAccountID := 0
	{
		accountConsumer := accountingapi.NewAccountRequest{
			UserID:   consumerID,
			UserType: "consumer",
		}
		reqBody, err := v1.Encode(accountConsumer)
		require.NoError(t, err)

		resp, err := http.Post(accountingURL+"/v1/accounts", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := accountingapi.AccountResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		consumerAccountID = respData.ID
		accountConsumerResp := accountingapi.AccountResponse{
			ID:       consumerAccountID,
			UserID:   consumerID,
			UserType: accountConsumer.UserType,
			Balance:  0,
		}
		equalAccount(t, accountConsumerResp, respData)
	}

	defer func() {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodDelete, accountingURL+"/v1/accounts/"+strconv.Itoa(consumerAccountID), nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}
	}()

	//consumer add balance
	consumerTrID := 0
	{
		accountConsumerTr := accountingapi.TransactionRequest{
			ToAccountID: consumerAccountID,
			Amount:      100,
		}

		accountConsumerTr.ToAccountID = consumerAccountID
		reqBody, err := v1.Encode(accountConsumerTr)
		require.NoError(t, err)

		resp, err := http.Post(accountingURL+"/v1/transactions", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := accountingapi.TransactionResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		consumerTrID = respData.ID
		accountConsumerTrResp := accountingapi.TransactionResponse{
			ID:            consumerTrID,
			FromAccountID: 0,
			ToAccountID:   consumerID,
			Amount:        accountConsumerTr.Amount,
			Valid:         true,
		}
		equalTr(t, respData, accountConsumerTrResp)
	}

	//account courier
	courierAccountID := 0
	{
		accountCourier := accountingapi.NewAccountRequest{
			UserID:   courierID,
			UserType: "courier",
		}
		reqBody, err := v1.Encode(accountCourier)
		require.NoError(t, err)

		resp, err := http.Post(accountingURL+"/v1/accounts", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := accountingapi.AccountResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		courierAccountID = respData.ID
		accountCourierResp := accountingapi.AccountResponse{
			ID:       courierAccountID,
			UserID:   courierID,
			UserType: accountCourier.UserType,
			Balance:  0,
		}
		equalAccount(t, accountCourierResp, respData)
	}

	defer func() {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodDelete, accountingURL+"/v1/accounts/"+strconv.Itoa(consumerAccountID), nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}
	}()

	//courier add balance
	courierTrID := 0
	{
		accountCourierTr := accountingapi.TransactionRequest{
			ToAccountID: courierAccountID,
			Amount:      100,
		}

		reqBody, err := v1.Encode(accountCourierTr)
		require.NoError(t, err)

		resp, err := http.Post(accountingURL+"/v1/transactions", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := accountingapi.TransactionResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		courierTrID = respData.ID
		accountCourierTrResp := accountingapi.TransactionResponse{
			ID:            courierTrID,
			FromAccountID: 0,
			ToAccountID:   courierID,
			Amount:        accountCourierTr.Amount,
			Valid:         true,
		}
		equalTr(t, respData, accountCourierTrResp)
	}

	//account restaurant
	restaurantAccountID := 0
	{
		accountRestaurant := accountingapi.NewAccountRequest{
			UserID:   restaurantID,
			UserType: "restaurant",
		}
		reqBody, err := v1.Encode(accountRestaurant)
		require.NoError(t, err)

		resp, err := http.Post(accountingURL+"/v1/accounts", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := accountingapi.AccountResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		restaurantAccountID = respData.ID
		accountRestaurantResp := accountingapi.AccountResponse{
			ID:       restaurantAccountID,
			UserID:   restaurantID,
			UserType: accountRestaurant.UserType,
			Balance:  0,
		}
		equalAccount(t, accountRestaurantResp, respData)
	}

	defer func() {
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodDelete, accountingURL+"/v1/accounts/"+strconv.Itoa(restaurantAccountID), nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}
	}()

	restaurantTrID := 0
	//restaurant add balance
	{
		accountRestaurantTr := accountingapi.TransactionRequest{
			ToAccountID: restaurantAccountID,
			Amount:      100,
		}

		reqBody, err := v1.Encode(accountRestaurantTr)
		require.NoError(t, err)

		resp, err := http.Post(accountingURL+"/v1/transactions/"+strconv.Itoa(restaurantAccountID), "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := accountingapi.TransactionResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		restaurantTrID = respData.ID
		accountRestaurantTrResp := accountingapi.TransactionResponse{
			ID:            restaurantTrID,
			FromAccountID: 0,
			ToAccountID:   restaurantAccountID,
			Amount:        accountRestaurantTr.Amount,
			Valid:         true,
		}
		equalTr(t, respData, accountRestaurantTrResp)
	}

	//delivery estimate
	{
		resp, err := http.Get(deliveryURL + "/v1//estimate" + "?consumer_id=" + strconv.Itoa(consumerID) + "&restaurant_id=" + strconv.Itoa(restaurantID))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := deliveryapi.EstimateDeliveryResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
	}

	//order
	orderID := 0
	{
		orderData := orderapi.OrderData{
			FromUserID:   consumerID,
			RestaurantID: restaurantID,
			OrderItems:   []int{1},
		}

		reqBody, err := v1.Encode(orderData)
		require.NoError(t, err)

		resp, err := http.Post(orderURL+"/v1/orders", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := orderapi.ReturnOrder{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		orderID = respData.OrderID

		orderDataResp := orderapi.ReturnOrder{
			OrderID:      orderID,
			FromUserID:   consumerID,
			RestaurantID: restaurantTrID,
			OrderItems:   []int{1},
			Status:       "",
		}
		assert.Equal(t, orderDataResp, respData)
	}

	//delivery
	{

		AssignCourier := deliveryapi.AssignOrderRequest{
			FromUserID:   consumerID,
			RestaurantID: restaurantID,
		}
		AssignCourierResp := deliveryapi.AssignOrderResponse{
			OrderID:   orderID,
			CourierID: courierID,
		}

		reqBody, err := v1.Encode(AssignCourier)
		require.NoError(t, err)

		resp, err := http.Post(deliveryURL+"/v1/orders/"+strconv.Itoa(orderID)+"/assign", "application/json", bytes.NewBuffer(reqBody))

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("StatusCode: %d", resp.StatusCode)
		}

		respData := deliveryapi.AssignOrderResponse{}
		err = consumerapi.DecodeJSON(resp.Body, &respData)
		require.NoError(t, err)
		assert.Equal(t, AssignCourierResp, respData)
	}

}

// equalAccount - function to compare selected fields from struct. Fields which needed
func equalAccount(t *testing.T, get accountingapi.AccountResponse, want accountingapi.AccountResponse) {
	if get.UserID != want.UserID {
		t.Errorf("UserID: Expected: %v, Got: %v", want.UserID, get.UserID)
	}

	if get.UserType != want.UserType {
		t.Errorf("UserType: Expected: %s, Got: %s", want.UserType, get.UserType)
	}

	if get.Balance != want.Balance {
		t.Errorf("Balance: Expected: %v, Got: %v", want.Balance, get.Balance)
	}
}

// equalTr - function to compare selected fields from struct. Fields which needed
func equalTr(t *testing.T, get accountingapi.TransactionResponse, want accountingapi.TransactionResponse) {
	if get.FromAccountID != want.FromAccountID {
		t.Errorf("FromAccountID: Expected: %v, Got: %v", want.FromAccountID, get.FromAccountID)
	}

	if get.ToAccountID != want.ToAccountID {
		t.Errorf("ToAccountID: Expected: %v, Got: %v", want.ToAccountID, get.ToAccountID)
	}

	if get.Amount != want.Amount {
		t.Errorf("Amount: Expected: %v, Got: %v", want.Amount, get.Amount)
	}
}
