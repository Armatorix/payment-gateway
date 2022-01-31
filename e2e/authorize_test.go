package e2e

import (
	"net/http"
	"testing"
	"time"

	apiclient "github.com/Armatorix/payment-gateway/api-client"
	"github.com/stretchr/testify/require"
)

var (
	invalidLuhnCardNumber = pointerString("7312731273127312")
	expiredCardDates      = &apiclient.AuthorizeRequestCardExpiry{ // expired
		Month: pointerInt32(12),
		Year:  pointerInt32(2021),
	}
	tooShortCVV = pointerInt32(23)
	zeroAmount  = pointerInt64(0)

	okAuthorizationReq = apiclient.AuthorizeRequest{
		CardNumber: pointerString("001230147647009683210024"),
		CardHolder: pointerString("dummy holder"),
		CardExpiry: &apiclient.AuthorizeRequestCardExpiry{
			Month: pointerInt32(12),
			Year:  pointerInt32(int32(time.Now().Year()) + 1),
		},
		CardCvv:  pointerInt32(123),
		Amount:   pointerInt64(1273),
		Currency: pointerString("PLN"),
	}
)

func testAuthorize(t *testing.T) {
	t.Run("proper request", testAuthorizeProper)
	t.Run("unauthorized", testAuthorizeUnauthoriezed)
	t.Run("bad luhn for card number", testAuthorizeBadLuhn)
	t.Run("expired card", testAuthorizeExpiredCard)
	t.Run("too short cvv", testAuthorizeTooShortCVV)
	t.Run("zero amount", testAuthorizeZeroAmount)
}

func testAuthorizeUnauthoriezed(t *testing.T) {
	_, resp, err := api.AuthorizePost(unauthedCtx).
		AuthorizeRequest(okAuthorizationReq).
		Execute()
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	require.Error(t, err)
	require.Equal(t, resp.StatusCode, http.StatusUnauthorized)
}
func testAuthorizeBadLuhn(t *testing.T) {
	_, resp, err := api.AuthorizePost(authedCtx).
		AuthorizeRequest(apiclient.AuthorizeRequest{
			CardNumber: invalidLuhnCardNumber,
			CardHolder: okAuthorizationReq.CardHolder,
			CardExpiry: okAuthorizationReq.CardExpiry,
			CardCvv:    okAuthorizationReq.CardCvv,
			Amount:     okAuthorizationReq.Amount,
			Currency:   okAuthorizationReq.Currency,
		}).Execute()

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	require.Error(t, err)
	require.Equal(t, resp.StatusCode, http.StatusBadRequest)
}

func testAuthorizeExpiredCard(t *testing.T) {
	_, resp, err := api.AuthorizePost(authedCtx).
		AuthorizeRequest(apiclient.AuthorizeRequest{
			CardNumber: okAuthorizationReq.CardNumber,
			CardHolder: okAuthorizationReq.CardHolder,
			CardExpiry: expiredCardDates,
			CardCvv:    okAuthorizationReq.CardCvv,
			Amount:     okAuthorizationReq.Amount,
			Currency:   okAuthorizationReq.Currency,
		}).Execute()

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	require.Error(t, err)
	require.Equal(t, resp.StatusCode, http.StatusBadRequest)
}

func testAuthorizeTooShortCVV(t *testing.T) {
	_, resp, err := api.AuthorizePost(authedCtx).
		AuthorizeRequest(apiclient.AuthorizeRequest{
			CardNumber: okAuthorizationReq.CardNumber,
			CardHolder: okAuthorizationReq.CardHolder,
			CardExpiry: okAuthorizationReq.CardExpiry,
			CardCvv:    tooShortCVV,
			Amount:     okAuthorizationReq.Amount,
			Currency:   okAuthorizationReq.Currency,
		}).Execute()

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	require.Error(t, err)
	require.Equal(t, resp.StatusCode, http.StatusBadRequest)
}

func testAuthorizeProper(t *testing.T) {
	authorize, resp, err := api.AuthorizePost(authedCtx).
		AuthorizeRequest(apiclient.AuthorizeRequest{
			CardNumber: okAuthorizationReq.CardNumber,
			CardHolder: okAuthorizationReq.CardHolder,
			CardExpiry: okAuthorizationReq.CardExpiry,
			CardCvv:    okAuthorizationReq.CardCvv,
			Amount:     okAuthorizationReq.Amount,
			Currency:   okAuthorizationReq.Currency,
		}).Execute()

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, http.StatusOK)
	require.Equal(t, &apiclient.AuthorizeResponse{
		AuthorizationId: authorize.AuthorizationId,
		AmountAvailable: okAuthorizationReq.Amount,
		Currency:        okAuthorizationReq.Currency,
		Status:          pointerString("success"),
	}, authorize)
}

func testAuthorizeZeroAmount(t *testing.T) {
	_, resp, err := api.AuthorizePost(authedCtx).
		AuthorizeRequest(apiclient.AuthorizeRequest{
			CardNumber: okAuthorizationReq.CardNumber,
			CardHolder: okAuthorizationReq.CardHolder,
			CardExpiry: okAuthorizationReq.CardExpiry,
			CardCvv:    okAuthorizationReq.CardCvv,
			Amount:     zeroAmount,
			Currency:   okAuthorizationReq.Currency,
		}).Execute()

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	require.Error(t, err)
	require.Equal(t, resp.StatusCode, http.StatusBadRequest)
}
