package architecture

type Concept struct {
	Mutable   bool
	Influence Approach
}

type Entity struct {
	Concept
}

type Value struct {
	Concept
}

type Port struct{}
type Adapter struct{}
type Direction int
