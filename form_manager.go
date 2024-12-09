package dhtml

import (
	"log"
	"net/http"
	"net/url"

	"github.com/mitoteam/mttools"
)

// HTML forms

type FormHandler struct {
	id string

	RenderF   func(form *FormElement)
	ValidateF func(fd *FormData)
	SubmitF   func(fd *FormData)
}

var FormManager formManagerT

type formManagerT struct {
	//form handlers list
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

// ========= FormData =========
type FormData struct {
	build_id string
	args     url.Values
	values   url.Values
}

var formDataStore map[string]*FormData

func init() {
	formDataStore = make(map[string]*FormData)
}

// ========= Form processing and rendering ===========
// Process form data, build it and render
func RenderForm(fh *FormHandler, request *http.Request) (out HtmlPiece) {
	var form_data *FormData

	if build_id := request.PostFormValue("form_build_id"); build_id != "" {
		if fd, ok := formDataStore[build_id]; ok {
			form_data = fd

			//hydrate form_data.values from POST data
			for name := range form_data.values {
				if postValue, ok := request.PostForm[name]; ok {
					form_data.values[name] = postValue
				}
			}
		}
	}

	// new form being built
	if form_data == nil {
		form_data = &FormData{
			build_id: "fd_" + mttools.RandomString(64),
			args:     make(url.Values),
			values:   make(url.Values),
		}
	}

	out.Append(Div().Textf("DBG %s: %s", request.Method, form_data.build_id))
	out.Append(Div().Textf("DBG post data: %v", request.PostForm))

	form := NewForm()
	form.formData = form_data

	fh.RenderF(form)

	// form.Append(build_id_hidden.renderF().String()) //DBG
	form.Append(NewFormHidden("form_build_id", form_data.build_id))

	formDataStore[form_data.build_id] = form_data

	out.Append(form)
	return out
}
