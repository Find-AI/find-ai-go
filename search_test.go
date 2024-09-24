// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findai_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stainless-sdks/find-ai-go"
	"github.com/stainless-sdks/find-ai-go/internal/testutil"
	"github.com/stainless-sdks/find-ai-go/option"
)

func TestSearchNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := findai.NewClient(
		option.WithBaseURL(baseURL),
	)
	resp, err := client.Searches.New(context.TODO(), findai.SearchNewParams{
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *findai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
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
	)
	err := client.Searches.Get(context.TODO(), "id")
	if err != nil {
		var apierr *findai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
