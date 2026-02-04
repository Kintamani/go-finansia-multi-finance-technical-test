package test

import (
	"encoding/json"
	"go-finansia-multi-finance-technical-test/internal/model"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	CreateContactLimit(t, contact, 1, 500000)

	requestBody := model.CreateTransactionRequest{
		ContractNumber:    "CN-1001",
		Tenor:             1,
		Channel:           "ecommerce",
		Otr:               100000,
		AdminFee:          5000,
		InstallmentAmount: 25000,
		InterestAmount:    2000,
		AssetName:         "Motor",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/contacts/"+contact.ID+"/transactions", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.TransactionResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.ContractNumber, responseBody.Data.ContractNumber)
	assert.Equal(t, requestBody.Tenor, responseBody.Data.Tenor)
	assert.Equal(t, requestBody.Channel, responseBody.Data.Channel)
	assert.Equal(t, requestBody.Otr, responseBody.Data.Otr)
	assert.Equal(t, requestBody.AdminFee, responseBody.Data.AdminFee)
	assert.Equal(t, requestBody.InstallmentAmount, responseBody.Data.InstallmentAmount)
	assert.Equal(t, requestBody.InterestAmount, responseBody.Data.InterestAmount)
	assert.Equal(t, requestBody.AssetName, responseBody.Data.AssetName)
	assert.NotNil(t, responseBody.Data.ID)
	assert.NotEmpty(t, responseBody.Data.CreatedAt)
	assert.NotEmpty(t, responseBody.Data.UpdatedAt)
}

func TestCreateTransactionExceeded(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	CreateContactLimit(t, contact, 1, 100000)

	requestBody := model.CreateTransactionRequest{
		ContractNumber:    "CN-1002",
		Tenor:             1,
		Channel:           "dealer",
		Otr:               200000,
		AdminFee:          5000,
		InstallmentAmount: 25000,
		InterestAmount:    2000,
		AssetName:         "Motor",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/contacts/"+contact.ID+"/transactions", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.TransactionResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestListTransactions(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	CreateContactLimit(t, contact, 1, 1000000)
	CreateTransactions(t, contact, 3, 1)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+contact.ID+"/transactions", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.TransactionResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 3, len(responseBody.Data))
}

func TestGetTransaction(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	CreateContactLimit(t, contact, 1, 1000000)
	CreateTransactions(t, contact, 1, 1)
	transaction := GetFirstTransaction(t, contact)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+contact.ID+"/transactions/"+transaction.ID, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.TransactionResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, transaction.ID, responseBody.Data.ID)
	assert.Equal(t, transaction.ContractNumber, responseBody.Data.ContractNumber)
	assert.Equal(t, transaction.Tenor, responseBody.Data.Tenor)
	assert.Equal(t, transaction.Channel, responseBody.Data.Channel)
	assert.Equal(t, transaction.Otr, responseBody.Data.Otr)
	assert.Equal(t, transaction.AdminFee, responseBody.Data.AdminFee)
	assert.Equal(t, transaction.InstallmentAmount, responseBody.Data.InstallmentAmount)
	assert.Equal(t, transaction.InterestAmount, responseBody.Data.InterestAmount)
	assert.Equal(t, transaction.AssetName, responseBody.Data.AssetName)
}

func TestUpdateTransaction(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	CreateContactLimit(t, contact, 1, 500000)

	createBody := model.CreateTransactionRequest{
		ContractNumber:    "CN-2001",
		Tenor:             1,
		Channel:           "ecommerce",
		Otr:               100000,
		AdminFee:          5000,
		InstallmentAmount: 25000,
		InterestAmount:    2000,
		AssetName:         "Motor",
	}
	createJson, err := json.Marshal(createBody)
	assert.Nil(t, err)

	createReq := httptest.NewRequest(http.MethodPost, "/api/contacts/"+contact.ID+"/transactions", strings.NewReader(string(createJson)))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Accept", "application/json")
	createReq.Header.Set("Authorization", user.Token)

	createResp, err := app.Test(createReq)
	assert.Nil(t, err)

	createBytes, err := io.ReadAll(createResp.Body)
	assert.Nil(t, err)

	createResponseBody := new(model.WebResponse[model.TransactionResponse])
	err = json.Unmarshal(createBytes, createResponseBody)
	assert.Nil(t, err)

	updateBody := model.UpdateTransactionRequest{
		ContractNumber:    "CN-2001-UPDATED",
		Tenor:             1,
		Channel:           "dealer",
		Otr:               200000,
		AdminFee:          6000,
		InstallmentAmount: 30000,
		InterestAmount:    3000,
		AssetName:         "Mobil",
	}
	updateJson, err := json.Marshal(updateBody)
	assert.Nil(t, err)

	updateReq := httptest.NewRequest(http.MethodPut, "/api/contacts/"+contact.ID+"/transactions/"+createResponseBody.Data.ID, strings.NewReader(string(updateJson)))
	updateReq.Header.Set("Content-Type", "application/json")
	updateReq.Header.Set("Accept", "application/json")
	updateReq.Header.Set("Authorization", user.Token)

	updateResp, err := app.Test(updateReq)
	assert.Nil(t, err)

	updateBytes, err := io.ReadAll(updateResp.Body)
	assert.Nil(t, err)

	updateResponseBody := new(model.WebResponse[model.TransactionResponse])
	err = json.Unmarshal(updateBytes, updateResponseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, updateResp.StatusCode)
	assert.Equal(t, updateBody.ContractNumber, updateResponseBody.Data.ContractNumber)
	assert.Equal(t, updateBody.Channel, updateResponseBody.Data.Channel)
	assert.Equal(t, updateBody.Otr, updateResponseBody.Data.Otr)
	assert.Equal(t, updateBody.AssetName, updateResponseBody.Data.AssetName)
}

func TestDeleteTransaction(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	CreateContactLimit(t, contact, 1, 500000)

	createBody := model.CreateTransactionRequest{
		ContractNumber:    "CN-3001",
		Tenor:             1,
		Channel:           "ecommerce",
		Otr:               100000,
		AdminFee:          5000,
		InstallmentAmount: 25000,
		InterestAmount:    2000,
		AssetName:         "Motor",
	}
	createJson, err := json.Marshal(createBody)
	assert.Nil(t, err)

	createReq := httptest.NewRequest(http.MethodPost, "/api/contacts/"+contact.ID+"/transactions", strings.NewReader(string(createJson)))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Accept", "application/json")
	createReq.Header.Set("Authorization", user.Token)

	createResp, err := app.Test(createReq)
	assert.Nil(t, err)

	createBytes, err := io.ReadAll(createResp.Body)
	assert.Nil(t, err)

	createResponseBody := new(model.WebResponse[model.TransactionResponse])
	err = json.Unmarshal(createBytes, createResponseBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodDelete, "/api/contacts/"+contact.ID+"/transactions/"+createResponseBody.Data.ID, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, true, responseBody.Data)
}
