// Copyright (c) 2024 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package provider

type ProfanityValidationRequest struct {
	Value  string `json:"value"`
	UserID string `json:"userId,omitempty"`
}

type ProfanityValidationResponse struct {
	IsProfane bool   `json:"isProfane"`
	Message   string `json:"message,omitempty"`
}
