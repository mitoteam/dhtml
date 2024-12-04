package dhtml

// Element is something that can be turned in list of html tags.
// Very simple elements are: tags itself, html comments or just plain text content.
// It could be much more complex things like Bootstrap's card for example.
// Whole HTML document is element as well (see dhtml.Document helper).
type ElementI interface {
	GetTags() TagsList
}

// helper types
type TagsList []*Tag
