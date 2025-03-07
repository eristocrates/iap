package domaintypes

type task struct {
	name
	taskId
}
type taskId struct {
	value int
}
type order struct{}
type minSessionToKeep struct{}
type parent struct {
	name
}
type parentName struct {
	value string
}
type ancestors struct{}
type shortNote struct{}
type name struct {
	value string
}
