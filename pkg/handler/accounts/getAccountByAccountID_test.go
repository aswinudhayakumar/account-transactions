package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aswinudhayakumar/account-transactions/internal/logger"
	"github.com/aswinudhayakumar/account-transactions/internal/mocks"
	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	getAccountByAccountIDEndpoint = "/app/v1/accounts/%d"
)

type testGetAccountByAccountIDSuite struct {
	suite.Suite

	dataRepo        *mocks.DataRepo
	router          *chi.Mux
	accountsHandler AccountsHandler
	recorder        *httptest.ResponseRecorder
}

func (s *testGetAccountByAccountIDSuite) SetupTest() {
	s.recorder = httptest.NewRecorder()
	s.dataRepo = new(mocks.DataRepo)
	s.accountsHandler = NewAccountsHandler(s.dataRepo)

	s.router = chi.NewRouter()
	s.router.Post("/app/v1/accounts/{id}", s.accountsHandler.GetAccountByAccountID)

	if err := logger.InitLogger(); err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
		return
	}
	defer logger.SyncLogger()
}

func TestGetAccountByAccountIDSuite(t *testing.T) {
	suite.Run(t, new(testGetAccountByAccountIDSuite))
}

func (s *testGetAccountByAccountIDSuite) TestGetAccountByAccountIDSuccess() {
	accountID := 1
	t := time.Now()
	expected := &repository.AccountResponse{
		AccountID:      1,
		DocumentNumber: "1234567890",
		CreatedAt:      t,
		UpdatedAt:      t,
	}

	s.dataRepo.Mock.On("GetAccountByAccountID", mock.Anything, accountID).
		Return(expected, nil)

	endpoint := fmt.Sprintf(getAccountByAccountIDEndpoint, accountID)
	req := httptest.NewRequest(http.MethodPost, endpoint, nil)

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusOK, s.recorder.Code)
}

func (s *testGetAccountByAccountIDSuite) TestGetAccountByAccountIDInvalidAccountID() {
	req := httptest.NewRequest(http.MethodPost, "/app/v1/accounts/abc", nil)

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
}

func (s *testGetAccountByAccountIDSuite) TestGetAccountByAccountIDDataNotFound() {
	accountID := 1

	s.dataRepo.Mock.On("GetAccountByAccountID", mock.Anything, accountID).
		Return(nil, sql.ErrNoRows)

	endpoint := fmt.Sprintf(getAccountByAccountIDEndpoint, accountID)
	req := httptest.NewRequest(http.MethodPost, endpoint, nil)

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusNotFound, s.recorder.Code)
}

func (s *testGetAccountByAccountIDSuite) TestGetAccountByAccountIDInternalServerError() {
	accountID := 1

	s.dataRepo.Mock.On("GetAccountByAccountID", mock.Anything, accountID).
		Return(nil, errors.New("something went wrong"))

	endpoint := fmt.Sprintf(getAccountByAccountIDEndpoint, accountID)
	req := httptest.NewRequest(http.MethodPost, endpoint, nil)

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusInternalServerError, s.recorder.Code)
}
