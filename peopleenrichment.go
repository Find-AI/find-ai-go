// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findai

import (
	"github.com/Find-AI/find-ai-go/option"
)

// PeopleEnrichmentService contains methods and other services that help with
// interacting with the find-ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPeopleEnrichmentService] method instead.
type PeopleEnrichmentService struct {
	Options []option.RequestOption
	Enrich  *PeopleEnrichmentEnrichService
}

// NewPeopleEnrichmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPeopleEnrichmentService(opts ...option.RequestOption) (r *PeopleEnrichmentService) {
	r = &PeopleEnrichmentService{}
	r.Options = opts
	r.Enrich = NewPeopleEnrichmentEnrichService(opts...)
	return
}
