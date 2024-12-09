package dhtml

import (
	"log"
	"net/http"
)

var FormManager formManagerT

func init() {
	FormManager = formManagerT{
		list: make(map[string]*FormHandler, 0),

		//simple default errors renderer
		renderErrorsF: func(fd *FormData) (out HtmlPiece) {
			container := Div().Class("form-errors")

			for name, itemErrors := range fd.errorList {
				for _, itemError := range itemErrors {
					errorOut := Div().Class("item-error")

					if name != "" {
						errorOut.Attribute("data-form-item-name", name).Textf("%s: ", name)
					}

					errorOut.Append(itemError)

					container.Append(errorOut)
				}
			}

			out.Append(container)
			return out
		},
	}
}

type formManagerT struct {
	//form handlers list
	list map[string]*FormHandler

	// function to render validation errors
	renderErrorsF func(fd *FormData) HtmlPiece
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

func (m *formManagerT) SetRenderErrorsF(f func(fd *FormData) HtmlPiece) *formManagerT {
	m.renderErrorsF = f
	return m
}

func (m *formManagerT) RenderForm(id string, w http.ResponseWriter, r *http.Request) *HtmlPiece {
	id = SafeId(id)

	return renderForm(m.GetHandler(id), w, r)
}
