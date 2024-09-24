// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findai_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/find-ai-go"
	"github.com/stainless-sdks/find-ai-go/internal/testutil"
	"github.com/stainless-sdks/find-ai-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := findai.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Searches.Get(context.TODO(), "id")
	if err != nil {
		t.Error(err)
	}
}
