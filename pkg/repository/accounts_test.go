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

type testAccountsTableSuite struct {
	suite.Suite

	db   *sqlx.DB
	mock sqlmock.Sqlmock
	repo DataRepo
}

func (s *testAccountsTableSuite) SetupTest() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	s.Require().NoError(err)

	sqlxDB := sqlx.NewDb(db, "postgres")

	s.db = sqlxDB
	s.mock = mock
	s.repo = NewDataRepo(sqlxDB)
}

func (s *testAccountsTableSuite) TearDownTest() {
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			return
		}
	}
}

func TestAccountsTableSuite(t *testing.T) {
	suite.Run(t, new(testAccountsTableSuite))
}

func (s *testAccountsTableSuite) TestCreateAccountSuccess() {
	req := CreateAccountReqParams{
		DocumentNumber: "1234567",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(createAccountQuery).
		WithArgs(req.DocumentNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repo.CreateAccount(context.Background(), req)
	s.Require().NoError(err)
}

func (s *testAccountsTableSuite) TestCreateAccountError() {
	req := CreateAccountReqParams{
		DocumentNumber: "1234567",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(createAccountQuery).
		WithArgs(req.DocumentNumber).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectRollback()

	err := s.repo.CreateAccount(context.Background(), req)
	s.Require().Error(err)
}

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

func (s *testAccountsTableSuite) TestGetAccountByAccountIDNoRowsError() {
	account_id := 1

	s.mock.ExpectQuery(getAccountByAccountIDQuery).
		WithArgs(account_id).
		WillReturnError(sql.ErrNoRows)

	actual, err := s.repo.GetAccountByAccountID(context.Background(), account_id)
	s.Require().Error(err)

	s.Require().Nil(actual)
}

func (s *testAccountsTableSuite) TestGetAccountByAccountIDInternalServerError() {
	account_id := 1

	s.mock.ExpectQuery(getAccountByAccountIDQuery).
		WithArgs(account_id).
		WillReturnError(errors.New("something went wrong"))

	actual, err := s.repo.GetAccountByAccountID(context.Background(), account_id)
	s.Require().Error(err)

	s.Require().Nil(actual)
}
