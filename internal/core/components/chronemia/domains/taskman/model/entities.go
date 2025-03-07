package taskman

type task struct {
	name
	taskId
}

type parent struct {
	name
}

type ancestors struct{}
type shortNote struct{}
