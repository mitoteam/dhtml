package dhtml

// simple <table> element
type TableElement struct {
	tag    *Tag // <table>
	header *TableRowElement
	tbody  *Tag

	emptyLabel string
}

// force interfaces implementation
var _ ElementI = (*TableElement)(nil)

func NewTable() *TableElement {
	return &TableElement{
		tag:    NewTag("table"),
		header: NewTableRow(),
		tbody:  NewTag("tbody"),
	}
}

func (e *TableElement) Class(v ...any) *TableElement {
	e.tag.Class(v...)
	return e
}

func (e *TableElement) EmptyLabel(label string) *TableElement {
	e.emptyLabel = label
	return e
}

// Adds header value (<tr><th>v</th></tr>) to the table
func (e *TableElement) Header(v any) *TableElement {
	th := NewTag("th").Append(v)
	e.header.tag.Append(th)
	return e
}

// Adds <TR> to the table data
func (e *TableElement) AppendRow(row *TableRowElement) *TableElement {
	e.tbody.Append(row)
	return e
}

// Adds class(es) to <tbody> tag
func (e *TableElement) BodyClass(v ...any) *TableElement {
	e.tbody.Class(v...)
	return e
}

// Creates new row, appends it to table and returns back
func (e *TableElement) NewRow() (row *TableRowElement) {
	row = NewTableRow()
	e.AppendRow(row)

	return row
}

// Returns added rows count.
func (e *TableElement) RowCount() int {
	return e.tbody.ChildrenCount()
}

func (e *TableElement) GetTags() TagList {
	//empty label set and no rows added - just show label
	if e.emptyLabel != "" && !e.tbody.HasChildren() {
		return EmptyLabel(e.emptyLabel).GetTags()
	}

	if e.header.tag.HasChildren() {
		e.tag.Append(NewTag("thead").Append(e.header.tag))
	}

	if e.tbody.HasChildren() {
		e.tag.Append(e.tbody)
	}

	return e.tag.GetTags()
}

//========== <TR> ===============

type TableRowElement struct {
	tag *Tag // <tr>
}

var _ ElementI = (*TableRowElement)(nil)

func NewTableRow() *TableRowElement {
	return &TableRowElement{tag: NewTag("tr")}
}

// Add <TD> to the row
func (e *TableRowElement) AppendCell(cell *TableCellElement) *TableRowElement {
	e.tag.Append(cell)
	return e
}

// Add new <TD> with given content to the row
func (e *TableRowElement) Cell(v any) *TableCellElement {
	cell := NewTableCell().Append(v)
	e.AppendCell(cell)
	return cell
}

func (e *TableRowElement) Class(v ...any) *TableRowElement {
	e.tag.Class(v...)
	return e
}

func (e *TableRowElement) GetTags() TagList {
	return e.tag.GetTags()
}

//========== <TD> ===============

type TableCellElement struct {
	tag *Tag // <td>
}

var _ ElementI = (*TableCellElement)(nil)

func NewTableCell() *TableCellElement {
	return &TableCellElement{tag: NewTag("td")}
}

func (e *TableCellElement) Class(v ...any) *TableCellElement {
	e.tag.Class(v...)
	return e
}

func (e *TableCellElement) Append(v any) *TableCellElement {
	e.tag.Append(v)
	return e
}

func (e *TableCellElement) GetTags() TagList {
	return e.tag.GetTags()
}
