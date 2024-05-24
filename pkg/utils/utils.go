// Copyright (c) 2024 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package utils

func Ternary[T any](condition bool, first, second T) T {
	if condition {
		return first
	}

	return second
}
