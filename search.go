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

// SearchService contains methods and other services that help with interacting
// with the find-ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r *SearchService) {
	r = &SearchService{}
	r.Options = opts
	return
}

// Starts a search.
func (r *SearchService) New(ctx context.Context, body SearchNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/searches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The endpoint to poll to check the latest results of a search.
func (r *SearchService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/searches/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type SearchNewParams struct {
	// The maximum number of results to return. optional for result_mode exact
	MaxMatches param.Field[float64] `json:"max_matches"`
	// Search query.
	Query param.Field[string] `json:"query"`
	// The mode of the search. Valid values are 'exact' or 'best'.
	ResultMode param.Field[string] `json:"result_mode"`
	// The scope of the search. Valid values are 'people' or 'companies'.
	Scope param.Field[string] `json:"scope"`
}

func (r SearchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
