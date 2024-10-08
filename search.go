// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findai

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Find-AI/find-ai-go/internal/apijson"
	"github.com/Find-AI/find-ai-go/internal/param"
	"github.com/Find-AI/find-ai-go/internal/requestconfig"
	"github.com/Find-AI/find-ai-go/option"
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
func (r *SearchService) New(ctx context.Context, body SearchNewParams, opts ...option.RequestOption) (res *SearchNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/searches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The endpoint to poll to check the latest results of a search.
func (r *SearchService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *[]SearchGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/searches/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SearchNewResponse struct {
	Poll SearchNewResponsePoll `json:"poll,required"`
	JSON searchNewResponseJSON `json:"-"`
}

// searchNewResponseJSON contains the JSON metadata for the struct
// [SearchNewResponse]
type searchNewResponseJSON struct {
	Poll        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchNewResponseJSON) RawJSON() string {
	return r.raw
}

type SearchNewResponsePoll struct {
	Token string                    `json:"token,required"`
	Path  string                    `json:"path,required"`
	JSON  searchNewResponsePollJSON `json:"-"`
}

// searchNewResponsePollJSON contains the JSON metadata for the struct
// [SearchNewResponsePoll]
type searchNewResponsePollJSON struct {
	Token       apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchNewResponsePoll) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchNewResponsePollJSON) RawJSON() string {
	return r.raw
}

type SearchGetResponse struct {
	LinkedinURL string `json:"linkedin_url,required"`
	Name        string `json:"name,required"`
	// Returned only for a person.
	Company            string                               `json:"company"`
	CriteriaAndReasons []SearchGetResponseCriteriaAndReason `json:"criteria_and_reasons"`
	// Returned only for a company.
	Domain string `json:"domain"`
	// Returned only for a person.
	Title string                `json:"title"`
	JSON  searchGetResponseJSON `json:"-"`
}

// searchGetResponseJSON contains the JSON metadata for the struct
// [SearchGetResponse]
type searchGetResponseJSON struct {
	LinkedinURL        apijson.Field
	Name               apijson.Field
	Company            apijson.Field
	CriteriaAndReasons apijson.Field
	Domain             apijson.Field
	Title              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SearchGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGetResponseJSON) RawJSON() string {
	return r.raw
}

type SearchGetResponseCriteriaAndReason struct {
	// Match criteria
	Criteria string `json:"criteria"`
	// Whether it's a match
	Match bool `json:"match"`
	// Reason for the match
	Reason string                                 `json:"reason"`
	JSON   searchGetResponseCriteriaAndReasonJSON `json:"-"`
}

// searchGetResponseCriteriaAndReasonJSON contains the JSON metadata for the struct
// [SearchGetResponseCriteriaAndReason]
type searchGetResponseCriteriaAndReasonJSON struct {
	Criteria    apijson.Field
	Match       apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchGetResponseCriteriaAndReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGetResponseCriteriaAndReasonJSON) RawJSON() string {
	return r.raw
}

type SearchNewParams struct {
	// The maximum number of results to return. optional for result_mode exact
	MaxMatches param.Field[float64] `json:"max_matches"`
	// Search query.
	Query param.Field[string] `json:"query"`
	// The mode of the search. Valid values are 'exact' or 'best'.
	ResultMode param.Field[SearchNewParamsResultMode] `json:"result_mode"`
	// The scope of the search. Valid values are 'person' or 'company'.
	Scope param.Field[SearchNewParamsScope] `json:"scope"`
}

func (r SearchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode of the search. Valid values are 'exact' or 'best'.
type SearchNewParamsResultMode string

const (
	SearchNewParamsResultModeExact SearchNewParamsResultMode = "exact"
	SearchNewParamsResultModeBest  SearchNewParamsResultMode = "best"
)

func (r SearchNewParamsResultMode) IsKnown() bool {
	switch r {
	case SearchNewParamsResultModeExact, SearchNewParamsResultModeBest:
		return true
	}
	return false
}

// The scope of the search. Valid values are 'person' or 'company'.
type SearchNewParamsScope string

const (
	SearchNewParamsScopePerson  SearchNewParamsScope = "person"
	SearchNewParamsScopeCompany SearchNewParamsScope = "company"
)

func (r SearchNewParamsScope) IsKnown() bool {
	switch r {
	case SearchNewParamsScopePerson, SearchNewParamsScopeCompany:
		return true
	}
	return false
}
