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
	createAccountEndpoint = "/app/v1/accounts"
)

type testCreateAccountSuite struct {
	suite.Suite

	dataRepo        *mocks.DataRepo
	router          *chi.Mux
	accountsHandler AccountsHandler
	recorder        *httptest.ResponseRecorder
}

func (s *testCreateAccountSuite) SetupTest() {
	s.recorder = httptest.NewRecorder()
	s.dataRepo = new(mocks.DataRepo)
	s.accountsHandler = NewAccountsHandler(s.dataRepo)

	s.router = chi.NewRouter()
	s.router.Post("/app/v1/accounts", s.accountsHandler.CreateAccount)

	if err := logger.InitLogger(); err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
		return
	}
	defer logger.SyncLogger()
}

func TestCreateAccountSuite(t *testing.T) {
	suite.Run(t, new(testCreateAccountSuite))
}

func (s *testCreateAccountSuite) TestCreateAccountSuccess() {
	documentNumber := "12345678900"
	s.dataRepo.Mock.On("CreateAccount", mock.Anything, repository.CreateAccountReqParams{
		DocumentNumber: documentNumber,
	}).Return(nil)

	reqBody := fmt.Sprintf(`{
		"document_number": "%s"
	}`, documentNumber)
	req := httptest.NewRequest(http.MethodPost, createAccountEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusCreated, s.recorder.Code)
}

func (s *testCreateAccountSuite) TestCreateAccountInvalidRequest() {
	documentNumber := "1"

	// invalid payload
	reqBody := fmt.Sprintf(`{
		"document_number": "%s"
	`, documentNumber)
	req := httptest.NewRequest(http.MethodPost, createAccountEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
}

func (s *testCreateAccountSuite) TestCreateAccountInvalidRequestPayload() {
	documentNumber := "1"

	// invalid field in request payload
	reqBody := fmt.Sprintf(`{
		"document_number": "%s"
	}`, documentNumber)
	req := httptest.NewRequest(http.MethodPost, createAccountEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
}

func (s *testCreateAccountSuite) TestCreateAccountInternalServerError() {
	documentNumber := "12345678900"
	s.dataRepo.Mock.On("CreateAccount", mock.Anything, repository.CreateAccountReqParams{
		DocumentNumber: documentNumber,
	}).Return(errors.New("something went wrong"))

	reqBody := fmt.Sprintf(`{
		"document_number": "%s"
	}`, documentNumber)
	req := httptest.NewRequest(http.MethodPost, createAccountEndpoint, strings.NewReader(reqBody))

	s.router.ServeHTTP(s.recorder, req)
	s.Equal(http.StatusInternalServerError, s.recorder.Code)
}
