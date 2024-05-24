// Copyright (c) 2024 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package server

import (
	"context"
	"os"
	registered_v1 "profanity-filter-grpc-plugin-server-go/pkg/pb"
	"profanity-filter-grpc-plugin-server-go/pkg/provider"
)

type ProfanityFilterServer struct {
	registered_v1.UnimplementedProfanityFilterServiceServer

	profanityFilterProvider provider.ProfanityFilterProvider
}

func NewProfanityFilterServer() *ProfanityFilterServer {
	providerName := provider.DefaultProviderName
	if str := os.Getenv("PROFANITY_FILTER_PROVIDER"); str != "" {
		providerName = str
	}
	server := ProfanityFilterServer{}
	switch providerName {
	case provider.DefaultProviderName:
		server.profanityFilterProvider = provider.NewDefaultProvider()
	default:
		server.profanityFilterProvider = provider.NewDefaultProvider()
	}

	return &server
}

func (server *ProfanityFilterServer) Validate(ctx context.Context, request *registered_v1.ExtendProfanityValidationRequest) (*registered_v1.ExtendProfanityValidationResponse, error) {
	profanityReq := provider.ProfanityValidationRequest{Value: request.Value}
	profanityResp, err := server.profanityFilterProvider.Validate(ctx, &profanityReq)
	if err != nil {
		return nil, err
	}

	return &registered_v1.ExtendProfanityValidationResponse{
		IsProfane: profanityResp.IsProfane,
		Message:   profanityResp.Message,
	}, nil
}
