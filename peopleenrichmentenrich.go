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

// PeopleEnrichmentEnrichService contains methods and other services that help with
// interacting with the Find AI API.
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
func (r *PeopleEnrichmentEnrichService) New(ctx context.Context, body PeopleEnrichmentEnrichNewParams, opts ...option.RequestOption) (res *PeopleEnrichmentEnrichNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/people/enrich"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The endpoint to poll to check the latest results when data about a person isn't
// immediately available.
func (r *PeopleEnrichmentEnrichService) Get(ctx context.Context, token string, opts ...option.RequestOption) (res *PeopleEnrichmentEnrichGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("v1/people/enrich/%s", token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PeopleEnrichmentEnrichNewResponse struct {
	Designation string `json:"designation,required"`
	Email       string `json:"email,required"`
	// A list of facts we have on record about the person.
	Facts     []string `json:"facts,required"`
	FirstName string   `json:"first_name,required"`
	LastName  string   `json:"last_name,required"`
	Name      string   `json:"name,required"`
	// A summary of information about the person.
	ShortDescription string `json:"short_description,required"`
	// The URL to the person's LinkedIn profile if one exists.
	LinkedinURL string `json:"linkedin_url"`
	// A URL to the person's profile image obtained from LinkedIn. Valid for 10
	// minutes.
	PhotoURL string                                `json:"photo_url"`
	JSON     peopleEnrichmentEnrichNewResponseJSON `json:"-"`
}

// peopleEnrichmentEnrichNewResponseJSON contains the JSON metadata for the struct
// [PeopleEnrichmentEnrichNewResponse]
type peopleEnrichmentEnrichNewResponseJSON struct {
	Designation      apijson.Field
	Email            apijson.Field
	Facts            apijson.Field
	FirstName        apijson.Field
	LastName         apijson.Field
	Name             apijson.Field
	ShortDescription apijson.Field
	LinkedinURL      apijson.Field
	PhotoURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PeopleEnrichmentEnrichNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peopleEnrichmentEnrichNewResponseJSON) RawJSON() string {
	return r.raw
}

type PeopleEnrichmentEnrichGetResponse struct {
	Designation string `json:"designation,required"`
	Email       string `json:"email,required"`
	// A list of facts we have on record about the person.
	Facts     []string `json:"facts,required"`
	FirstName string   `json:"first_name,required"`
	LastName  string   `json:"last_name,required"`
	Name      string   `json:"name,required"`
	// A summary of information about the person.
	ShortDescription string `json:"short_description,required"`
	// The URL to the person's LinkedIn profile if one exists.
	LinkedinURL string `json:"linkedin_url"`
	// A URL to the person's profile image obtained from LinkedIn. Valid for 10
	// minutes.
	PhotoURL string                                `json:"photo_url"`
	JSON     peopleEnrichmentEnrichGetResponseJSON `json:"-"`
}

// peopleEnrichmentEnrichGetResponseJSON contains the JSON metadata for the struct
// [PeopleEnrichmentEnrichGetResponse]
type peopleEnrichmentEnrichGetResponseJSON struct {
	Designation      apijson.Field
	Email            apijson.Field
	Facts            apijson.Field
	FirstName        apijson.Field
	LastName         apijson.Field
	Name             apijson.Field
	ShortDescription apijson.Field
	LinkedinURL      apijson.Field
	PhotoURL         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PeopleEnrichmentEnrichGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peopleEnrichmentEnrichGetResponseJSON) RawJSON() string {
	return r.raw
}

type PeopleEnrichmentEnrichNewParams struct {
	// The person's business email address. We won't accept personal email address such
	// as Gmail, Yahoo etc.
	Email param.Field[string] `json:"email"`
}

func (r PeopleEnrichmentEnrichNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
