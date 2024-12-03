package dhtml

type (
	ElementI interface {
		GetTags() TagsList
	}

	// helper type
	ElementsList []ElementI
	TagsList     []*Tag
)
