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

// CompanyEnrichmentEnrichService contains methods and other services that help
// with interacting with the Find AI API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompanyEnrichmentEnrichService] method instead.
type CompanyEnrichmentEnrichService struct {
	Options []option.RequestOption
}

// NewCompanyEnrichmentEnrichService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCompanyEnrichmentEnrichService(opts ...option.RequestOption) (r *CompanyEnrichmentEnrichService) {
	r = &CompanyEnrichmentEnrichService{}
	r.Options = opts
	return
}

// Returns structured data about a company based on a domain input.
func (r *CompanyEnrichmentEnrichService) New(ctx context.Context, body CompanyEnrichmentEnrichNewParams, opts ...option.RequestOption) (res *CompanyEnrichmentEnrichNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/companies/enrich"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The endpoint to poll to check the latest results when data about a company isn't
// immediately available.
func (r *CompanyEnrichmentEnrichService) Get(ctx context.Context, token string, opts ...option.RequestOption) (res *CompanyEnrichmentEnrichGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("v1/companies/enrich/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CompanyEnrichmentEnrichNewResponse struct {
	// A list of facts we have on record about the company.
	Facts []string `json:"facts,required"`
	Name  string   `json:"name,required"`
	// A summary of information about the company.
	ShortDescription string `json:"short_description,required"`
	Website          string `json:"website,required"`
	// The URL to the company's LinkedIn profile if one exists.
	LinkedinURL string `json:"linkedin_url"`
	// A URL to an image of the company's logo. Valid for 10 minutes.
	PhotoURL string                                 `json:"photo_url"`
	JSON     companyEnrichmentEnrichNewResponseJSON `json:"-"`
}

// companyEnrichmentEnrichNewResponseJSON contains the JSON metadata for the struct
// [CompanyEnrichmentEnrichNewResponse]
type companyEnrichmentEnrichNewResponseJSON struct {
	Facts            apijson.Field
	Name             apijson.Field
	ShortDescription apijson.Field
	Website          apijson.Field
	LinkedinURL      apijson.Field
	PhotoURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CompanyEnrichmentEnrichNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r companyEnrichmentEnrichNewResponseJSON) RawJSON() string {
	return r.raw
}

type CompanyEnrichmentEnrichGetResponse struct {
	// A list of facts we have on record about the company.
	Facts []string `json:"facts,required"`
	Name  string   `json:"name,required"`
	// A summary of information about the company.
	ShortDescription string `json:"short_description,required"`
	Website          string `json:"website,required"`
	// The URL to the company's LinkedIn profile if one exists.
	LinkedinURL string `json:"linkedin_url"`
	// A URL to an image of the company's logo. Valid for 10 minutes.
	PhotoURL string                                 `json:"photo_url"`
	JSON     companyEnrichmentEnrichGetResponseJSON `json:"-"`
}

// companyEnrichmentEnrichGetResponseJSON contains the JSON metadata for the struct
// [CompanyEnrichmentEnrichGetResponse]
type companyEnrichmentEnrichGetResponseJSON struct {
	Facts            apijson.Field
	Name             apijson.Field
	ShortDescription apijson.Field
	Website          apijson.Field
	LinkedinURL      apijson.Field
	PhotoURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CompanyEnrichmentEnrichGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r companyEnrichmentEnrichGetResponseJSON) RawJSON() string {
	return r.raw
}

type CompanyEnrichmentEnrichNewParams struct {
	// The domain of the company.
	Domain param.Field[string] `json:"domain"`
}

func (r CompanyEnrichmentEnrichNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
