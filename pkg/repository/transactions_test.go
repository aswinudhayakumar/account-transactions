package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type testTransactionsTableSuite struct {
	suite.Suite

	db   *sqlx.DB
	mock sqlmock.Sqlmock
	repo DataRepo
}

func (s *testTransactionsTableSuite) SetupTest() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	s.Require().NoError(err)

	sqlxDB := sqlx.NewDb(db, "postgres")

	s.db = sqlxDB
	s.mock = mock
	s.repo = NewDataRepo(sqlxDB)
}

func (s *testTransactionsTableSuite) TearDownTest() {
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			return
		}
	}
}

func TestTransactionsTableSuite(t *testing.T) {
	suite.Run(t, new(testTransactionsTableSuite))
}

func (s *testTransactionsTableSuite) TestCreateTransactionSuccess() {
	req := CreateTransactionReqParams{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100.12,
	}

	s.mock.ExpectBegin()

	validateResponse := sqlmock.NewRows([]string{"is_account_exists", "is_operation_type_id_exists"}).
		AddRow(true, true)
	s.mock.ExpectQuery(validateCreateTrxQuery).
		WithArgs(
			req.AccountID,
			req.OperationTypeID,
		).WillReturnRows(validateResponse)

	s.mock.ExpectExec(createTransactionQuery).
		WithArgs(
			req.AccountID,
			req.OperationTypeID,
			req.Amount,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectCommit()

	err := s.repo.CreateTransaction(context.Background(), req)
	s.Require().NoError(err)
}

func (s *testTransactionsTableSuite) TestCreateTransactionInvalidAccountID() {
	req := CreateTransactionReqParams{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100.12,
	}

	s.mock.ExpectBegin()

	validateResponse := sqlmock.NewRows([]string{"is_account_exists", "is_operation_type_id_exists"}).
		AddRow(false, true)
	s.mock.ExpectQuery(validateCreateTrxQuery).
		WithArgs(
			req.AccountID,
			req.OperationTypeID,
		).WillReturnRows(validateResponse)

	s.mock.ExpectRollback()

	err := s.repo.CreateTransaction(context.Background(), req)
	s.Require().Error(err)
	s.Require().ErrorIs(err, ErrAccountIDNotExists)
}

func (s *testTransactionsTableSuite) TestCreateTransactionInvalidOperationTypeID() {
	req := CreateTransactionReqParams{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100.12,
	}

	s.mock.ExpectBegin()

	validateResponse := sqlmock.NewRows([]string{"is_account_exists", "is_operation_type_id_exists"}).
		AddRow(true, false)
	s.mock.ExpectQuery(validateCreateTrxQuery).
		WithArgs(
			req.AccountID,
			req.OperationTypeID,
		).WillReturnRows(validateResponse)

	s.mock.ExpectRollback()

	err := s.repo.CreateTransaction(context.Background(), req)
	s.Require().Error(err)
	s.Require().ErrorIs(err, ErrOperationTypeIDNotExists)
}

func (s *testTransactionsTableSuite) TestCreateTransactionError() {
	req := CreateTransactionReqParams{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100.12,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(createTransactionQuery).
		WithArgs(
			req.AccountID,
			req.OperationTypeID,
			req.Amount,
		).WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectRollback()

	err := s.repo.CreateTransaction(context.Background(), req)
	s.Require().Error(err)
}
