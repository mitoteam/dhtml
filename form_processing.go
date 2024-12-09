package dhtml

import (
	"net/http"
	"net/url"

	"github.com/mitoteam/mttools"
)

type FormHandler struct {
	Id string

	RenderF   func(form *FormElement, fd *FormData)
	ValidateF func(fd *FormData)
	SubmitF   func(fd *FormData)
}

// ========= FormData =========
var formDataStore map[string]*FormData

func init() {
	formDataStore = make(map[string]*FormData)
}

type FormErrors map[string][]HtmlPiece

type FormData struct {
	build_id string
	args     url.Values
	values   url.Values

	errorList   FormErrors //map of error lists by form item name
	rebuild     bool       // rebuild form with same data again
	redirectUrl string     // issue an redirect to this URL
}

func NewFormData() *FormData {
	return &FormData{
		build_id:  "fd_" + mttools.RandomString(64),
		args:      make(url.Values),
		values:    make(url.Values),
		errorList: make(FormErrors, 0),
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

func (fd *FormData) IsRebuild() bool {
	return fd.rebuild
}

func (fd *FormData) SetRebuild(v bool) {
	fd.rebuild = v
}

func (fd *FormData) GetRedirect() string {
	return fd.redirectUrl
}

func (fd *FormData) SetRedirect(url string) {
	fd.redirectUrl = url
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

func (fd *FormData) GetErrors() FormErrors {
	return fd.errorList
}

func (fd *FormData) ClearErrors() {
	fd.errorList = make(map[string][]HtmlPiece, 0)
}

// Process form data, build it and render
func renderForm(fh *FormHandler, w http.ResponseWriter, r *http.Request) *HtmlPiece {
	var formOut HtmlPiece
	var fd *FormData
	var isPost bool //false - form is being built for the first time, true - processing form submit by POST request

	// check if it is being re-build (from POST request)
	if build_id := r.PostFormValue("form_build_id"); build_id != "" {
		if fd, isPost = formDataStore[build_id]; isPost {
			//hydrate form_data.values from POST data
			for name := range fd.values {
				if postValue, ok := r.PostForm[name]; ok {
					fd.values[name] = postValue
				}
			}
		}
	}

	//out.Append(Dbg("%s: %s", request.Method, form_data.build_id))
	//out.Append(Dbg("RAW post data: %v", request.PostForm))
	//out.Append(Dbg("FormData.values: %v", form_data.values))

	if isPost {
		fd.SetRedirect("")
		fd.SetRebuild(false)
		fd.ClearErrors()

		fh.ValidateF(fd)

		if len(fd.errorList) > 0 {
			formOut.Append(FormManager.renderErrorsF(fd))
			fd.SetRebuild(true) //and display form again
		} else {
			//there were no errors
			fh.SubmitF(fd)
		}

		if !fd.rebuild {
			delete(formDataStore, fd.build_id)

			if fd.redirectUrl != "" {
				http.Redirect(w, r, fd.redirectUrl, http.StatusSeeOther)
				return NewHtmlPiece() //empty html
			}

			//we are not rebuilding, should be completely new FormData
			fd = nil
		}
	}

	if fd == nil {
		fd = NewFormData()
	}

	form := NewForm()
	form.formData = fd

	fh.RenderF(form, fd)

	// form.Append(build_id_hidden.renderF().String()) //DBG
	form.Append(NewFormHidden("form_build_id", fd.build_id))

	formDataStore[fd.build_id] = fd

	formOut.Append(form)

	//wrap it all into container <div>
	div := Div().Class("dhtml-form").Attribute("data-form-id", fh.Id).
		Append(formOut)

	return Piece(div)
}
