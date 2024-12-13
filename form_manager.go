package dhtml

import (
	"log"
)

var FormManager formManagerT

func init() {
	FormManager = formManagerT{
		list: make(map[string]*FormHandler, 0),
	}
}

type formManagerT struct {
	//form handlers list
	list map[string]*FormHandler
}

func (m *formManagerT) Register(form *FormHandler) {
	id := SafeId(form.Id)

	if m.IsRegistered(id) {
		log.Panicf("Form id '%s' already registered", id)
	}

	form.Id = id
	m.list[id] = form
}

func (m *formManagerT) IsRegistered(id string) bool {
	id = SafeId(id)

	_, ok := m.list[id]
	return ok
}

func (m *formManagerT) GetHandler(id string) *FormHandler {
	id = SafeId(id)

	if form_handler, ok := m.list[id]; ok {
		return form_handler
	} else {
		log.Panicf("Form id '%s' not registered", id)
		return nil
	}
}

func (m *formManagerT) RenderForm(id string, fc *FormContext) *HtmlPiece {
	id = SafeId(id)

	//log.Printf("DBG formManagerT.RenderForm fc: %+v\n", fc)

	return renderForm(m.GetHandler(id), fc)
}
