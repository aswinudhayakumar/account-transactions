package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aswinudhayakumar/account-transactions/internal/logger"
	"github.com/aswinudhayakumar/account-transactions/internal/mocks"
	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	createTransactionEndpoint = "/app/v1/transactions"
)

// testCreateTransactionSuite is a test suite object to test CreateTransaction API handler.
type testCreateTransactionSuite struct {
	suite.Suite

	dataRepo   *mocks.DataRepo
	router     *chi.Mux
	trxHandler TransactionsHandler
	recorder   *httptest.ResponseRecorder
}

// SetupTest setups and initializes the testCreateTransactionSuite.
func (s *testCreateTransactionSuite) SetupTest() {
	s.recorder = httptest.NewRecorder()
	s.dataRepo = new(mocks.DataRepo)
	s.trxHandler = NewTransactionsHandler(s.dataRepo)

	s.router = chi.NewRouter()
	s.router.Post("/app/v1/transactions", s.trxHandler.CreateTransaction)

	if err := logger.InitLogger(); err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
		return
	}
	defer logger.SyncLogger()
}

// TestCreateTransactionSuite is the custom test suite runner for CreateTransaction API handler.
func TestCreateTransactionSuite(t *testing.T) {
	suite.Run(t, new(testCreateTransactionSuite))
}

// @Success testcase - statusCode (201)
func (s *testCreateTransactionSuite) TestCreateTransactionSuccess() {
	accountID := 1
	operationTypeID := 1
	amount := 100.50

	s.dataRepo.Mock.On("CreateTransaction", mock.Anything, repository.CreateTransactionReqParams{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}).Return(nil)

	reqBody := fmt.Sprintf(`{
		"account_id": %d,
		"operation_type_id": %d,
		"amount": %f
	}`, accountID, operationTypeID, amount)
	req := httptest.NewRequest(http.MethodPost, createTransactionEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusCreated, s.recorder.Code)
}

// @Failed testcase - statusCode (400)
func (s *testCreateTransactionSuite) TestCreateTransactionInvalidRequest() {
	accountID := 1
	operationTypeID := 1
	amount := 100.50

	reqBody := fmt.Sprintf(`{
		"account_id": %d,
		"operation_type_id": %d,
		"amount": %f,
	}`, accountID, operationTypeID, amount)
	req := httptest.NewRequest(http.MethodPost, createTransactionEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
}

// @Failed testcase - statusCode (400)
func (s *testCreateTransactionSuite) TestCreateTransactionInvalidAccountID() {
	operationTypeID := 1
	amount := 100.50

	reqBody := fmt.Sprintf(`{
		"account_id": "test",
		"operation_type_id": %d,
		"amount": %f
	}`, operationTypeID, amount)
	req := httptest.NewRequest(http.MethodPost, createTransactionEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
}

// @Failed testcase - statusCode (400)
func (s *testCreateTransactionSuite) TestCreateTransactionInvalidOperationTypeID() {
	accountID := 1
	amount := 100.50

	reqBody := fmt.Sprintf(`{
		"account_id": %d,
		"operation_type_id": "test",
		"amount": %f,
	}`, accountID, amount)
	req := httptest.NewRequest(http.MethodPost, createTransactionEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
}

// @Failed testcase - statusCode (400)
func (s *testCreateTransactionSuite) TestCreateTransactionInvalidAmount() {
	accountID := 1
	operationTypeID := 1

	reqBody := fmt.Sprintf(`{
		"account_id": %d,
		"operation_type_id": %d,
		"amount": "test"
	}`, accountID, operationTypeID)
	req := httptest.NewRequest(http.MethodPost, createTransactionEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
}

// @Failed testcase - statusCode (400)
func (s *testCreateTransactionSuite) TestCreateTransactionAccountIDNotFound() {
	accountID := 1
	operationTypeID := 1
	amount := 100.50

	s.dataRepo.Mock.On("CreateTransaction", mock.Anything, repository.CreateTransactionReqParams{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}).Return(repository.ErrAccountIDNotExists)

	reqBody := fmt.Sprintf(`{
		"account_id": %d,
		"operation_type_id": %d,
		"amount": %f
	}`, accountID, operationTypeID, amount)
	req := httptest.NewRequest(http.MethodPost, createTransactionEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
}

// @Failed testcase - statusCode (400)
func (s *testCreateTransactionSuite) TestCreateTransactionOperationTypeIDNotFound() {
	accountID := 1
	operationTypeID := 1
	amount := 100.50

	s.dataRepo.Mock.On("CreateTransaction", mock.Anything, repository.CreateTransactionReqParams{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}).Return(repository.ErrOperationTypeIDNotExists)

	reqBody := fmt.Sprintf(`{
		"account_id": %d,
		"operation_type_id": %d,
		"amount": %f
	}`, accountID, operationTypeID, amount)
	req := httptest.NewRequest(http.MethodPost, createTransactionEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
}

// @Failed testcase - statusCode (500)
func (s *testCreateTransactionSuite) TestCreateTransactionInternalServerError() {
	accountID := 1
	operationTypeID := 1
	amount := 100.50

	s.dataRepo.Mock.On("CreateTransaction", mock.Anything, repository.CreateTransactionReqParams{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}).Return(errors.New("something went wrong"))

	reqBody := fmt.Sprintf(`{
		"account_id": %d,
		"operation_type_id": %d,
		"amount": %f
	}`, accountID, operationTypeID, amount)
	req := httptest.NewRequest(http.MethodPost, createTransactionEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusInternalServerError, s.recorder.Code)
}
