package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"pismo-back-teste/internal"
	"pismo-back-teste/internal/api/dto"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type fakeAccountRepository struct {
	GetAccountMethod    func(ctx context.Context, id int) (dto.Account, error)
	CreateAccountMethod func(ctx context.Context, account dto.Account) error
}

func (r *fakeAccountRepository) GetAccount(ctx context.Context, id int) (dto.Account, error) {
	return r.GetAccountMethod(ctx, id)
}

func (r *fakeAccountRepository) CreateAccount(ctx context.Context, account dto.Account) error {
	return r.CreateAccountMethod(ctx, account)
}

type fakeRepoProvider struct {
	accountResponse dto.Account
	err             error
}

func (r *fakeRepoProvider) AccountRepository(ctx context.Context) internal.IAccountRepository {
	accountRepo := &fakeAccountRepository{}
	accountRepo.GetAccountMethod = func(ctx context.Context, id int) (dto.Account, error) {
		return r.accountResponse, r.err
	}

	accountRepo.CreateAccountMethod = func(ctx context.Context, account dto.Account) error {
		return r.err
	}

	return accountRepo
}

func (r *fakeRepoProvider) TransactionRepository(ctx context.Context) internal.ITransactionRepository {
	return nil
}

func fakeServiceProvider(accountResponse dto.Account, err error) *internal.ServiceProvider {
	fakeRepoProvider := fakeRepoProvider{
		accountResponse: accountResponse,
		err:             err,
	}
	return internal.NewServiceProvider(&fakeRepoProvider)
}

var accountJSON = `{"id":1,"document":"123456789"}`
var emptyAccountJSON = `{"id":0,"document":""}`

func setupGetAccount() (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/accounts/:accountId")
	c.SetParamNames("accountId")
	c.SetParamValues("1")
	return rec, c
}

func TestGet(t *testing.T) {
	t.Run("returns an account and status ok", func(t *testing.T) {
		rec, c := setupGetAccount()
		accountResponse := dto.Account{ID: 1, Document: "123456789"}
		h := NewAccountHandler(fakeServiceProvider(accountResponse, nil))

		// Assertions
		if assert.NoError(t, h.get(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, accountJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
		}
		rec.Body.Reset()
	})

	t.Run("returns an empty account object and status not found", func(t *testing.T) {
		rec, c := setupGetAccount()
		accountResponse := dto.Account{ID: 0, Document: ""}
		h := NewAccountHandler(fakeServiceProvider(accountResponse, nil))

		// Assertions
		if assert.NoError(t, h.get(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, emptyAccountJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
		}
	})
}

var reqAccountJSON = `{"id":0,"document":"123456789"}`

func setupCreateAccount() (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqAccountJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}

func TestPost(t *testing.T) {
	t.Run("saves an account and returns status created", func(t *testing.T) {
		rec, c := setupCreateAccount()
		h := NewAccountHandler(fakeServiceProvider(dto.Account{}, nil))
		// Assertions
		if assert.NoError(t, h.post(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, reqAccountJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
		}
	})
}
