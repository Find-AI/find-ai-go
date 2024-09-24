// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findai

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/find-ai-go/internal/apijson"
	"github.com/stainless-sdks/find-ai-go/internal/param"
	"github.com/stainless-sdks/find-ai-go/internal/requestconfig"
	"github.com/stainless-sdks/find-ai-go/option"
)

// PeopleEnrichmentEnrichService contains methods and other services that help with
// interacting with the find-ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPeopleEnrichmentEnrichService] method instead.
type PeopleEnrichmentEnrichService struct {
	Options []option.RequestOption
}

// NewPeopleEnrichmentEnrichService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPeopleEnrichmentEnrichService(opts ...option.RequestOption) (r *PeopleEnrichmentEnrichService) {
	r = &PeopleEnrichmentEnrichService{}
	r.Options = opts
	return
}

// Returns structured data about a person based on their business email address.
func (r *PeopleEnrichmentEnrichService) New(ctx context.Context, body PeopleEnrichmentEnrichNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/people/enrich"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The endpoint to poll to check the latest results when data about a person isn't
// immediately available.
func (r *PeopleEnrichmentEnrichService) Get(ctx context.Context, token string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("v1/people/enrich/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PeopleEnrichmentEnrichNewParams struct {
	// The person's business email address. We won't accept personal email address such
	// as Gmail, Yahoo etc.
	Email param.Field[string] `json:"email"`
}

func (r PeopleEnrichmentEnrichNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
