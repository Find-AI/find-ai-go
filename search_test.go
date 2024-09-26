// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Find-AI/find-ai-go"
	"github.com/Find-AI/find-ai-go/internal/testutil"
	"github.com/Find-AI/find-ai-go/option"
)

func TestSearchNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := findai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Searches.New(context.TODO(), findai.SearchNewParams{
		MaxMatches: findai.F(0.000000),
		Query:      findai.F("query"),
		ResultMode: findai.F("result_mode"),
		Scope:      findai.F("scope"),
	})
	if err != nil {
		var apierr *findai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := findai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Searches.Get(context.TODO(), "id")
	if err != nil {
		var apierr *findai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
