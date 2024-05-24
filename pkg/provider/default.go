// Copyright (c) 2024 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package provider

import (
	"context"
	"profanity-filter-grpc-plugin-server-go/pkg/utils"

	goaway "github.com/TwiN/go-away"
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
