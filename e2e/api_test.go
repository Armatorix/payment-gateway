package e2e

import (
	"context"
	"testing"

	apiclient "github.com/Armatorix/payment-gateway/api-client"
)

var (
	api       = apiclient.NewAPIClient(apiclient.NewConfiguration()).DefaultApi
	authedCtx = context.WithValue(context.Background(),
		apiclient.ContextBasicAuth,
		apiclient.BasicAuth{
			UserName: "test_merchant",
			Password: "test_api_secret",
		})
	unauthedCtx = context.WithValue(context.Background(),
		apiclient.ContextBasicAuth,
		apiclient.BasicAuth{
			UserName: "test_merchant",
			Password: "bad_api_secret",
		})
)

func TestAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping e2e tests in short mode")
	}
	t.Run("authorize", testAuthorize)
}
