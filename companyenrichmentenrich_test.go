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

func TestCompanyEnrichmentEnrichNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CompanyEnrichment.Enrich.New(context.TODO(), findai.CompanyEnrichmentEnrichNewParams{
		Domain: findai.F("domain"),
	})
	if err != nil {
		var apierr *findai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCompanyEnrichmentEnrichGet(t *testing.T) {
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
	_, err := client.CompanyEnrichment.Enrich.Get(context.TODO(), "token")
	if err != nil {
		var apierr *findai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
