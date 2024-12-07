package dhtml

import (
	"log"
)

// HTML forms

type FormHandler struct {
	id string

	RenderF   func() *FormElement
	ValidateF func()
	SubmitF   func()
}

var FormManager formManagerT

type formManagerT struct {
	list map[string]*FormHandler
}

func (m *formManagerT) Register(id string, form *FormHandler) {
	id = SafeId(id)

	if m.list == nil {
		m.list = make(map[string]*FormHandler, 0)
	} else {
		if _, ok := m.list[id]; ok {
			log.Panicf("Form id '%s' already registered", id)
		}
	}

	form.id = id
	m.list[id] = form
}

func (m *formManagerT) IsRegistered(id string) bool {
	_, ok := m.list[id]
	return ok
}

func (m *formManagerT) GetHandler(id string) *FormHandler {
	if form_handler, ok := m.list[id]; ok {
		return form_handler
	} else {
		log.Panicf("Form id '%s' not registered", id)
		return nil
	}
}

func (m *formManagerT) Render(id string) *FormElement {
	form_handler := m.GetHandler(id)

	return form_handler.RenderF()
}
