package cmd

import (
	"context"
)

func Execute(ctx context.Context) int {
	if err := runScreenshotCopy(ctx).Execute(); err != nil {
		return 1
	}
	return 0
}
