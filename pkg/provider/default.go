package provider

import (
	"context"
	goaway "github.com/TwiN/go-away"
	"profanity-filter-grpc-plugin-server-go/pkg/utils"
)

const DefaultProviderName = "default"

type DefaultProvider struct {
}

func NewDefaultProvider() *DefaultProvider {
	return &DefaultProvider{}
}

func (provider *DefaultProvider) Validate(ctx context.Context, req *ProfanityValidationRequest) (*ProfanityValidationResponse, error) {
	isProfane := goaway.IsProfane(req.Value)
	message := utils.Ternary(isProfane, "this contains banned words", "")
	return &ProfanityValidationResponse{
		IsProfane: isProfane,
		Message:   message,
	}, nil
}
