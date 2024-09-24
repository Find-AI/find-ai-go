// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findai

import (
	"github.com/stainless-sdks/find-ai-go/option"
)

// CompanyEnrichmentService contains methods and other services that help with
// interacting with the find-ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompanyEnrichmentService] method instead.
type CompanyEnrichmentService struct {
	Options []option.RequestOption
	Enrich  *CompanyEnrichmentEnrichService
}

// NewCompanyEnrichmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCompanyEnrichmentService(opts ...option.RequestOption) (r *CompanyEnrichmentService) {
	r = &CompanyEnrichmentService{}
	r.Options = opts
	r.Enrich = NewCompanyEnrichmentEnrichService(opts...)
	return
}
