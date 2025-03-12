package taskman

import (
	taskvals "iap/internal/core/components/modauxdo/domains/taskman/entities/task/values"
	sharevals "iap/internal/shared/values"
)

type task struct {
	id   sharevals.Id
	name taskvals.Name
}
