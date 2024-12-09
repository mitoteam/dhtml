package dhtml

import "log"

type FormItemI interface {
	ElementI
	GetName() string
	GetId() string

	GetValue() any
	SetValue(v any)
}

type FormItemBase struct {
	name         string
	defaultValue any
	value        any
	renderF      RenderFunc
}

func (fi *FormItemBase) GetName() string {
	return fi.name
}

func (fi *FormItemBase) GetId() string {
	return "id_" + fi.name
}

func (fi *FormItemBase) GetValue() any {
	if fi.value != nil {
		return fi.value
	}

	return fi.defaultValue
}

func (fi *FormItemBase) SetValue(v any) {
	fi.value = v
}

func (fi *FormItemBase) GetTags() TagsList {
	if fi.renderF == nil {
		log.Panic("Form item render function not set")
		return nil
	} else {
		return Div().Class("form-item").
			Append(fi.renderF()).
			GetTags()
	}
}

// ========= FormItemExtI ================

type FormItemExtI interface {
	FormItemI
	GetLabel() HtmlPiece
	GetNote() HtmlPiece
}

type FormItemExtBase struct {
	FormItemBase
	label HtmlPiece
	note  HtmlPiece
}

func (fi *FormItemExtBase) GetLabel() HtmlPiece {
	return fi.label
}

func (fi *FormItemExtBase) GetNote() HtmlPiece {
	return fi.note
}