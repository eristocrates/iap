package thoughtman

import (
	values "iap/internal/components/domains/thoughtman/entities/thought/values"
	shared "iap/internal/shared/values"
)

type thought struct {
	shared.Id
	values.Title
	shared.CreatedAt
	// TODO decide on how to have edit rollback on each thought.
}
