package test

import (
	"encoding/json"
	"go-finansia-multi-finance-technical-test/internal/entity"
	"go-finansia-multi-finance-technical-test/internal/model"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateContact(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	requestBody := model.CreateContactRequest{
		NIK:         "3200000000000001",
		FullName:    "Nono Cahyono",
		LegalName:   "Nono",
		BirthPlace:  "Jakarta",
		BirthDate:   "1990-01-01",
		Salary:      10000000,
		KtpPhoto:    "https://example.com/ktp-eko.jpg",
		SelfiePhoto: "https://example.com/selfie-eko.jpg",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/contacts", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.NIK, responseBody.Data.NIK)
	assert.Equal(t, requestBody.FullName, responseBody.Data.FullName)
	assert.Equal(t, requestBody.LegalName, responseBody.Data.LegalName)
	assert.Equal(t, requestBody.BirthPlace, responseBody.Data.BirthPlace)
	assert.Equal(t, requestBody.BirthDate, responseBody.Data.BirthDate)
	assert.Equal(t, requestBody.Salary, responseBody.Data.Salary)
	assert.Equal(t, requestBody.KtpPhoto, responseBody.Data.KtpPhoto)
	assert.Equal(t, requestBody.SelfiePhoto, responseBody.Data.SelfiePhoto)
	assert.NotNil(t, responseBody.Data.ID)
	assert.NotEmpty(t, responseBody.Data.CreatedAt)
	assert.NotEmpty(t, responseBody.Data.UpdatedAt)
}

func TestCreateContactFailed(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	requestBody := model.CreateContactRequest{
		NIK:         "",
		FullName:    "",
		LegalName:   "",
		BirthPlace:  "",
		BirthDate:   "",
		Salary:      0,
		KtpPhoto:    "",
		SelfiePhoto: "",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/contacts", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestGetConnect(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	contact := new(entity.Contact)
	err = db.Where("user_id = ?", user.ID).First(contact).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+contact.ID, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, contact.ID, responseBody.Data.ID)
	assert.Equal(t, contact.NIK, responseBody.Data.NIK)
	assert.Equal(t, contact.FullName, responseBody.Data.FullName)
	assert.Equal(t, contact.LegalName, responseBody.Data.LegalName)
	assert.Equal(t, contact.BirthPlace, responseBody.Data.BirthPlace)
	assert.Equal(t, contact.BirthDate, responseBody.Data.BirthDate)
	assert.Equal(t, contact.Salary, responseBody.Data.Salary)
	assert.Equal(t, contact.KtpPhoto, responseBody.Data.KtpPhoto)
	assert.Equal(t, contact.SelfiePhoto, responseBody.Data.SelfiePhoto)
	assert.Equal(t, contact.CreatedAt.Format(time.RFC3339), responseBody.Data.CreatedAt)
	assert.Equal(t, contact.UpdatedAt.Format(time.RFC3339), responseBody.Data.UpdatedAt)
}

func TestGetContactFailed(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+uuid.NewString(), nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestUpdateContact(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	contact := new(entity.Contact)
	err = db.Where("user_id = ?", user.ID).First(contact).Error
	assert.Nil(t, err)

	requestBody := model.UpdateContactRequest{
		NIK:         "3200000000000002",
		FullName:    "Eko Budiman",
		LegalName:   "Eko Budiman",
		BirthPlace:  "Bandung",
		BirthDate:   "1991-02-02",
		Salary:      12000000,
		KtpPhoto:    "https://example.com/ktp-budiman.jpg",
		SelfiePhoto: "https://example.com/selfie-budiman.jpg",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/contacts/"+contact.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.NIK, responseBody.Data.NIK)
	assert.Equal(t, requestBody.FullName, responseBody.Data.FullName)
	assert.Equal(t, requestBody.LegalName, responseBody.Data.LegalName)
	assert.Equal(t, requestBody.BirthPlace, responseBody.Data.BirthPlace)
	assert.Equal(t, requestBody.BirthDate, responseBody.Data.BirthDate)
	assert.Equal(t, requestBody.Salary, responseBody.Data.Salary)
	assert.Equal(t, requestBody.KtpPhoto, responseBody.Data.KtpPhoto)
	assert.Equal(t, requestBody.SelfiePhoto, responseBody.Data.SelfiePhoto)
	assert.NotNil(t, responseBody.Data.ID)
	assert.NotEmpty(t, responseBody.Data.CreatedAt)
	assert.NotEmpty(t, responseBody.Data.UpdatedAt)
}

func TestUpdateContactFailed(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	contact := new(entity.Contact)
	err = db.Where("user_id = ?", user.ID).First(contact).Error
	assert.Nil(t, err)

	requestBody := model.UpdateContactRequest{
		NIK:         "",
		FullName:    "",
		LegalName:   "",
		BirthPlace:  "",
		BirthDate:   "",
		Salary:      0,
		KtpPhoto:    "",
		SelfiePhoto: "",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/contacts/"+contact.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestUpdateContactNotFound(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	requestBody := model.UpdateContactRequest{
		NIK:         "",
		FullName:    "",
		LegalName:   "",
		BirthPlace:  "",
		BirthDate:   "",
		Salary:      0,
		KtpPhoto:    "",
		SelfiePhoto: "",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/contacts/"+uuid.NewString(), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestDeleteContact(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	contact := new(entity.Contact)
	err = db.Where("user_id = ?", user.ID).First(contact).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodDelete, "/api/contacts/"+contact.ID, nil)
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

func TestDeleteContactFailed(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodDelete, "/api/contacts/"+uuid.NewString(), nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestSearchContact(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	CreateContacts(user, 20)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 10, len(responseBody.Data))
	assert.Equal(t, int64(20), responseBody.Paging.TotalItem)
	assert.Equal(t, int64(2), responseBody.Paging.TotalPage)
	assert.Equal(t, 1, responseBody.Paging.Page)
	assert.Equal(t, 10, responseBody.Paging.Size)
}

func TestSearchContactWithPagination(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	CreateContacts(user, 20)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts?page=2&size=5", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 5, len(responseBody.Data))
	assert.Equal(t, int64(20), responseBody.Paging.TotalItem)
	assert.Equal(t, int64(4), responseBody.Paging.TotalPage)
	assert.Equal(t, 2, responseBody.Paging.Page)
	assert.Equal(t, 5, responseBody.Paging.Size)
}

func TestSearchContactWithFilter(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("username = ?", "khannedy").First(user).Error
	assert.Nil(t, err)

	CreateContacts(user, 20)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts?full_name=Contact&legal_name=Contact", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.ContactResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 10, len(responseBody.Data))
	assert.Equal(t, int64(20), responseBody.Paging.TotalItem)
	assert.Equal(t, int64(2), responseBody.Paging.TotalPage)
	assert.Equal(t, 1, responseBody.Paging.Page)
	assert.Equal(t, 10, responseBody.Paging.Size)
}
