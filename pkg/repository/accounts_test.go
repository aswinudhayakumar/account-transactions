package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

// testAccountsTableSuite is a test suite object to test database operations from Accounts table.
type testAccountsTableSuite struct {
	suite.Suite

	db   *sqlx.DB
	mock sqlmock.Sqlmock
	repo DataRepo
}

// SetupTest setups and initializes the testAccountsTableSuite.
func (s *testAccountsTableSuite) SetupTest() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	s.Require().NoError(err)

	sqlxDB := sqlx.NewDb(db, "postgres")

	s.db = sqlxDB
	s.mock = mock
	s.repo = NewDataRepo(sqlxDB)
}

// TearDownTest gracefully closes the test suite, by closing the db connection.
func (s *testAccountsTableSuite) TearDownTest() {
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			return
		}
	}
}

// TestAccountsTableSuite is the custom test suite to test database operations from Accounts table.
func TestAccountsTableSuite(t *testing.T) {
	suite.Run(t, new(testAccountsTableSuite))
}

// @Success testcase
func (s *testAccountsTableSuite) TestCreateAccountSuccess() {
	req := CreateAccountReqParams{
		DocumentNumber: "1234567",
	}

	s.mock.ExpectExec(createAccountQuery).
		WithArgs(req.DocumentNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repo.CreateAccount(context.Background(), req)
	s.Require().NoError(err)
}

// @Failed testcase
func (s *testAccountsTableSuite) TestCreateAccountError() {
	req := CreateAccountReqParams{
		DocumentNumber: "1234567",
	}

	s.mock.ExpectExec(createAccountQuery).
		WithArgs(req.DocumentNumber).
		WillReturnError(errors.New("something went wrong"))

	err := s.repo.CreateAccount(context.Background(), req)
	s.Require().Error(err)
}

// @Success testcase
func (s *testAccountsTableSuite) TestGetAccountByAccountIDSuccess() {
	t := time.Now()
	expected := &AccountResponse{
		AccountID:      1,
		DocumentNumber: "1234567",
		CreatedAt:      t,
		UpdatedAt:      t,
	}

	sqlResponse := sqlmock.NewRows([]string{"account_id", "document_number", "created_at", "updated_at"}).
		AddRow(
			expected.AccountID,
			expected.DocumentNumber,
			expected.CreatedAt,
			expected.UpdatedAt,
		)

	s.mock.ExpectQuery(getAccountByAccountIDQuery).
		WithArgs(expected.AccountID).
		WillReturnRows(sqlResponse)

	actual, err := s.repo.GetAccountByAccountID(context.Background(), expected.AccountID)
	s.Require().NoError(err)

	s.Equal(expected, actual)
}

// @Failed testcase
func (s *testAccountsTableSuite) TestGetAccountByAccountIDNoRowsError() {
	account_id := 1

	s.mock.ExpectQuery(getAccountByAccountIDQuery).
		WithArgs(account_id).
		WillReturnError(sql.ErrNoRows)

	actual, err := s.repo.GetAccountByAccountID(context.Background(), account_id)
	s.Require().Error(err)

	s.Require().Nil(actual)
}

// @Failed testcase
func (s *testAccountsTableSuite) TestGetAccountByAccountIDInternalServerError() {
	account_id := 1

	s.mock.ExpectQuery(getAccountByAccountIDQuery).
		WithArgs(account_id).
		WillReturnError(errors.New("something went wrong"))

	actual, err := s.repo.GetAccountByAccountID(context.Background(), account_id)
	s.Require().Error(err)

	s.Require().Nil(actual)
}
