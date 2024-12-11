package dhtml

import (
	"net/http"

	"github.com/mitoteam/mttools"
)

type FormHandler struct {
	Id string

	RenderF   func(form *FormElement, fd *FormData)
	ValidateF func(fd *FormData)
	SubmitF   func(fd *FormData)
}

// ========= FormContext =========

type FormContext struct {
	params      mttools.Values // copied to form data each time form is rendered (even if it is being rebuild)
	args        mttools.Values // copied to form data on first build only and stored between builds
	w           http.ResponseWriter
	r           *http.Request
	redirectUrl string // issue an redirect to this URL (FormData redirectUrl has priority)
}

func NewFormContext(w http.ResponseWriter, r *http.Request) *FormContext {
	fc := &FormContext{
		w: w, r: r,
		params: mttools.NewValues(),
		args:   mttools.NewValues(),
	}

	return fc
}

func (fc *FormContext) SetParam(key string, v any) *FormContext {
	fc.params.Set(key, v)
	return fc
}

func (fc *FormContext) GetParam(key string) any {
	return fc.params.Get(key)
}

func (fc *FormContext) SetArg(key string, v any) *FormContext {
	fc.args.Set(key, v)
	return fc
}

func (fc *FormContext) GetArg(key string) any {
	return fc.args.Get(key)
}

func (fc *FormContext) SetRedirect(url string) *FormContext {
	fc.redirectUrl = url
	return fc
}

// ========= FormData =========
var formDataStore map[string]*FormData

func init() {
	formDataStore = make(map[string]*FormData)
}

type FormErrorsT map[string][]HtmlPiece

type FormData struct {
	build_id string
	args     mttools.Values
	params   mttools.Values
	values   mttools.Values
	labels   NamedHtmlPieces

	errorList   FormErrorsT //map of error lists by form item name
	rebuild     bool        // rebuild form with same data again
	redirectUrl string      // issue an redirect to this URL
}

func NewFormData() *FormData {
	return &FormData{
		build_id:  "fd_" + mttools.RandomString(64),
		args:      mttools.NewValues(),
		params:    mttools.NewValues(),
		values:    mttools.NewValues(),
		errorList: make(FormErrorsT, 0),
		labels:    NewNamedHtmlPieces(),
	}
}

func (fd *FormData) GetArg(name string) any {
	return fd.args.Get(name)
}

func (fd *FormData) GetAllArgs() mttools.Values {
	return fd.args
}

func (fd *FormData) GetParam(name string) any {
	return fd.params.Get(name)
}

func (fd *FormData) GetValue(name string) any {
	return fd.values.Get(name)
}

func (fd *FormData) SetValue(name string, v any) *FormData {
	fd.values.Set(name, v)
	return fd
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

func (fd *FormData) GetLabel(form_item_name string) *HtmlPiece {
	if piece, ok := fd.labels.GetOk(form_item_name); ok {
		return piece
	}

	return Piece(form_item_name) //not found, so return element's name as label
}

func (fd *FormData) SetItemError(form_item_name string, v any) {
	if _, ok := fd.errorList[form_item_name]; !ok {
		fd.errorList[form_item_name] = make([]HtmlPiece, 0, 1)
	}

	fd.errorList[form_item_name] = append(fd.errorList[form_item_name], *Piece(v))
}

func (fd *FormData) SetError(v any) {
	//empty item name = common error
	fd.SetItemError("", v)
}

func (fd *FormData) HasError() bool {
	return len(fd.errorList) > 0
}

func (fd *FormData) GetErrors() FormErrorsT {
	return fd.errorList
}

func (fd *FormData) ClearErrors() {
	fd.errorList = make(map[string][]HtmlPiece, 0)
}

// Process form data, build it and render
func renderForm(fh *FormHandler, fc *FormContext) *HtmlPiece {
	var formOut HtmlPiece
	var fd *FormData
	var isPost bool //false - form is being built for the first time, true - processing form submit by POST request

	// check if it is being re-build (from POST request)
	if build_id := fc.r.PostFormValue("form_build_id"); build_id != "" {
		if fd, isPost = formDataStore[build_id]; isPost {
			//refresh params from context
			fd.params.CopyFrom(fc.params)

			//re-hydrate form_data.values from POST data
			for name := range fd.values.GetNamesIterator() {
				if postValue, ok := fc.r.PostForm[name]; ok {
					if len(postValue) == 1 {
						fd.values.Set(name, postValue[0])
					} else {
						fd.values.Set(name, postValue)
					}
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

		if fd.HasError() {
			formOut.Append(FormManager.renderErrorsF(fd))
			fd.SetRebuild(true) //and display form again
		} else {
			//there were no errors
			fh.SubmitF(fd)
		}

		if !fd.rebuild {
			delete(formDataStore, fd.build_id)

			//check redirect (first from FormData, then from FormContext)
			var redirectUrl = fd.redirectUrl

			if redirectUrl == "" {
				redirectUrl = fc.redirectUrl
			}

			if redirectUrl != "" {
				http.Redirect(fc.w, fc.r, redirectUrl, http.StatusSeeOther)
				//fc.w.
				return NewHtmlPiece() //empty html
			}

			//we are not rebuilding, not redirecting = new form should be built from scratch
			fd = nil
		}
	}

	if fd == nil {
		fd = NewFormData()
		fd.args.CopyFrom(fc.args)
		fd.params.CopyFrom(fc.params)
	}

	form := NewForm().Append(NewFormHidden("form_build_id", fd.build_id))
	form.formData = fd

	fh.RenderF(form, fd)

	formDataStore[fd.build_id] = fd

	formOut.Append(form)

	//wrap it all into container <div>
	div := Div().Class("dhtml-form").Attribute("data-form-id", fh.Id).
		Append(formOut)

	return Piece(div)
}
