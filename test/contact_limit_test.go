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

func TestCreateContactLimit(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)

	requestBody := model.CreateContactLimitRequest{
		Tenor:       1,
		LimitAmount: 100000,
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/contacts/"+contact.ID+"/limits", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactLimitResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Tenor, responseBody.Data.Tenor)
	assert.Equal(t, requestBody.LimitAmount, responseBody.Data.LimitAmount)
	assert.Equal(t, int64(0), responseBody.Data.UsedAmount)
	assert.Equal(t, requestBody.LimitAmount, responseBody.Data.RemainingAmount)
	assert.NotNil(t, responseBody.Data.ID)
	assert.NotEmpty(t, responseBody.Data.CreatedAt)
	assert.NotEmpty(t, responseBody.Data.UpdatedAt)
}

func TestCreateContactLimitFailed(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)

	requestBody := model.CreateContactLimitRequest{
		Tenor:       5,
		LimitAmount: 100000,
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/contacts/"+contact.ID+"/limits", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactLimitResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestListContactLimits(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)

	CreateContactLimit(t, contact, 1, 100000)
	CreateContactLimit(t, contact, 2, 200000)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+contact.ID+"/limits", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.ContactLimitResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 2, len(responseBody.Data))
}

func TestGetContactLimit(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	limit := CreateContactLimit(t, contact, 3, 300000)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+contact.ID+"/limits/"+limit.ID, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactLimitResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, limit.ID, responseBody.Data.ID)
	assert.Equal(t, limit.Tenor, responseBody.Data.Tenor)
	assert.Equal(t, limit.LimitAmount, responseBody.Data.LimitAmount)
}

func TestUpdateContactLimit(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	limit := CreateContactLimit(t, contact, 3, 300000)

	requestBody := model.UpdateContactLimitRequest{
		Tenor:       6,
		LimitAmount: 600000,
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/contacts/"+contact.ID+"/limits/"+limit.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactLimitResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Tenor, responseBody.Data.Tenor)
	assert.Equal(t, requestBody.LimitAmount, responseBody.Data.LimitAmount)
}

func TestDeleteContactLimit(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	limit := CreateContactLimit(t, contact, 3, 300000)

	request := httptest.NewRequest(http.MethodDelete, "/api/contacts/"+contact.ID+"/limits/"+limit.ID, nil)
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
