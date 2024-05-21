package provider

import "context"

type ProfanityFilterProvider interface {
	Validate(ctx context.Context, req *ProfanityValidationRequest) (*ProfanityValidationResponse, error)
}
