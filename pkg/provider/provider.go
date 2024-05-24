// Copyright (c) 2024 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package provider

import "context"

type ProfanityFilterProvider interface {
	Validate(ctx context.Context, req *ProfanityValidationRequest) (*ProfanityValidationResponse, error)
}
