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
var formDataStore map[string]*FormData

func init() {
	formDataStore = make(map[string]*FormData)
}

type FormData struct {
	build_id string
	args     url.Values
	values   url.Values

	errorList map[string][]HtmlPiece //map of error lists by form item name
}

func NewFormData() *FormData {
	return &FormData{
		build_id:  "fd_" + mttools.RandomString(64),
		args:      make(url.Values),
		values:    make(url.Values),
		errorList: make(map[string][]HtmlPiece, 0),
	}
}

func (fd *FormData) GetValue(name string) any {
	if v, ok := fd.values[name]; ok {
		if len(v) == 1 { //url.Values values are always arrays
			return v[0]
		} else {
			return v
		}
	} else {
		return nil
	}
}

func (fd *FormData) SetItemError(form_item_name string, v any) {
	if _, ok := fd.errorList[form_item_name]; !ok {
		fd.errorList[form_item_name] = make([]HtmlPiece, 0, 1)
	}

	fd.errorList[form_item_name] = append(fd.errorList[form_item_name], *Piece(v))
}

func (fd *FormData) SetError(name string, v any) {
	//empty item name = common error
	fd.SetItemError("", v)
}

func (fd *FormData) ClearErrors() {
	fd.errorList = make(map[string][]HtmlPiece, 0)
}

func (fd *FormData) RenderErrors() (out HtmlPiece) {
	for name, itemErrors := range fd.errorList {
		for _, itemError := range itemErrors {
			errorOut := Div().Class("item-error")

			if name != "" {
				errorOut.Textf("%s: ", name)
			}

			errorOut.Append(itemError)

			out.Append(errorOut)
		}
	}

	return out
}

// ========= Form processing and rendering ===========
// Process form data, build it and render
func RenderForm(fh *FormHandler, request *http.Request) *HtmlPiece {
	var out HtmlPiece
	var form_data *FormData
	var is_rebuild bool

	// check if it is being re-build (from POST request)
	if build_id := request.PostFormValue("form_build_id"); build_id != "" {
		if form_data, is_rebuild = formDataStore[build_id]; is_rebuild {
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
		form_data = NewFormData()
	}

	//out.Append(Dbg("%s: %s", request.Method, form_data.build_id))
	//out.Append(Dbg("RAW post data: %v", request.PostForm))
	//out.Append(Dbg("FormData.values: %v", form_data.values))

	if is_rebuild {
		fh.ValidateF(form_data)

		if len(form_data.errorList) > 0 {
			out.Append(Div().Class("errors").Append(form_data.RenderErrors()))
		}
	}

	form := NewForm()
	form.formData = form_data

	fh.RenderF(form)

	// form.Append(build_id_hidden.renderF().String()) //DBG
	form.Append(NewFormHidden("form_build_id", form_data.build_id))

	formDataStore[form_data.build_id] = form_data

	out.Append(form)

	//wrap it all into container <div>
	div := Div().Class("dhtml-form").
		Append(out)

	return Piece(div)
}
