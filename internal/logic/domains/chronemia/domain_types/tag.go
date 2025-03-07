package domaintypes

type tag struct{}

type timeMap struct {
	timeMapId
}
type timeMapId struct {
	value int
}
type name struct {
	value string
}
type createdAt struct {
	value Time.time
}
type plan struct{}
type updatedAt struct {
	value Time.time
}
type user struct {
	userId
}
type userId struct {
	value int
}
type color struct{}
type duration struct {
	value Time.duration
}
type budgetZone struct {
	budgetZoneId
}
type budgetZoneId struct {
	value int
}
