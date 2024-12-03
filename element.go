package dhtml

type (
	ElementI interface {
		GetTags() []*Tag
	}

	ElementsList []ElementI
)
