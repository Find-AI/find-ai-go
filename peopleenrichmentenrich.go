// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findai

import (
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
