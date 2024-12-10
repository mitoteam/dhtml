package dhtml

import (
	"fmt"
	"log"
)

type RenderTemplateFunc func(t *HtmlTemplate) HtmlPiece

type HtmlTemplate struct {
	regions NamedHtmlPieces
	renderF RenderTemplateFunc
}

// force interfaces implementation
var _ fmt.Stringer = (*HtmlTemplate)(nil)

func NewHtmlTemplate(renderF RenderTemplateFunc) *HtmlTemplate {
	return &HtmlTemplate{
		regions: NewNamedHtmlPieces(),
		renderF: renderF,
	}
}

// Adds some content to named region
func (t *HtmlTemplate) ClearRegions() *HtmlTemplate {
	t.regions = NewNamedHtmlPieces()
	return t
}

// Adds some content to named region
func (t *HtmlTemplate) Add(region_name string, v any) *HtmlTemplate {
	t.regions.Add(region_name, v)
	return t
}

// Returns region's contents
func (t *HtmlTemplate) Region(region_name string) *HtmlPiece {
	return t.regions.Get(region_name)
}

func (t *HtmlTemplate) String() string {
	if t.renderF == nil {
		log.Panicf("renderF function is nil")
	}

	p := t.renderF(t)

	return p.String()
}
