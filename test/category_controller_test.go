package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/model/domain"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@/learn-golang-restful-api-test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)                  // Set maximum idle connections
	db.SetMaxOpenConns(20)                 // Set maximum open connections
	db.SetConnMaxLifetime(60 * time.Minute) // Set maximum connection lifetime to 60 minutes
	db.SetConnMaxIdleTime(10 * time.Minute) // Set maximum idle time for connections to 10 minutes

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

// mengosongkan data kategori tiap menjalankan test
func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
	db.Exec("ALTER TABLE category AUTO_INCREMENT = 1")
}


// *SKENARIO 1: CREATE CATEGORY SUKSES
func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Food"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RahasiaPabrik")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Food", responseBody["data"].(map[string]interface{})["name"])
}

// !SKENARIO 2: CREATE CATEGORY GAGAL
func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RahasiaPabrik")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

// *SKENARIO 3: UPDATE CATEGORY SUKSES
func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Voucher",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Voucher"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RahasiaPabrik")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Voucher", responseBody["data"].(map[string]interface{})["name"])
}

// !SKENARIO 4: UPDATE CATEGORY GAGAL
func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Voucher",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RahasiaPabrik")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

// *SKENARIO 5: DELETE CATEGORY SUKSES
func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Sports",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RahasiaPabrik")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

// !SKENARIO 6: DELETE CATEGORY GAGAL
func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RahasiaPabrik")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

// *SKENARIO 7: GET CATEGORY BY ID SUKSES
func TestGetCategoryByIdSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Laundry",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-Key", "RahasiaPabrik")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}

// !SKENARIO 8: GET CATEGORY BY ID GAGAL
func TestGetCategoryByIdFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("X-API-Key", "RahasiaPabrik")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

// *SKENARIO 9: GET ALL CATEGORY SUKSES
func TestGetAllCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Sports",
	})
	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Foods",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "RahasiaPabrik")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	var categories = responseBody["data"].([]interface{})

	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])
	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])
}

// *SKENARIO 10: UNAUTHORIZED
func TestUnauthorizedSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "Salah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}