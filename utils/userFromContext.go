package utils

import (
	"context"

	"github.com/google/uuid"
)

func UserIdFromContext(ctx context.Context) uuid.UUID {
	return uuid.New()
}
