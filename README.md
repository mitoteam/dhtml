# dhtml - Go html renderer

[![Reference](https://pkg.go.dev/badge/github.com/mitoteam/dhtml.svg)](https://pkg.go.dev/github.com/mitoteam/dhtml)
![GitHub code size](https://img.shields.io/github/languages/code-size/mitoteam/dhtml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mitoteam/dhtml)](https://goreportcard.com/report/github.com/mitoteam/dhtml)
![GitHub](https://img.shields.io/github/license/mitoteam/dhtml)


[![GitHub Version](https://img.shields.io/github/v/release/mitoteam/dhtml?logo=github)](https://github.com/mitoteam/dhtml)
[![GitHub Release](https://img.shields.io/github/release-date/mitoteam/dhtml)](https://github.com/mitoteam/dhtml/releases)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/mitoteam/dhtml)](https://github.com/mitoteam/dhtml/commits)

Go html renderer inspired by [Drupal](https://www.drupal.org)'s renderable arrays concept. Allows to built some elements tree in Go and then render it as html.

Ready to try beta version.

Code Example:
```go
//Build elements tree, <html> as root element
html := dhtml.Tag("html").
  Comment("some <html> escaped comment").
  //appending children elements
  Append(
    dhtml.Tag("head").
      Append(
        //element with attributes
        dhtml.Tag("link").Attribute("href", "/assets/vendor/bootstrap.min.css").Attribute("rel", "stylesheet"),
      ),
  ).
  //body element
  Append(dhtml.Tag("body").
    Append(
      // dhtml.Div() is a shorthand for dhtml.Tag("div")
      dhtml.Div().Id("test").
        Attribute("data-mt-test", "some attribute").
        //classes deduplication
        Class("border").Class("m-3").Class("p-3").
        Content("some <escaped> text"),
    ),
    Append(
      dhtml.Div()
        //multiple classes
        Classes([]string{"border", "border-danger", "border-5"}).
        Content("another text in red rectangle"),
    ),
  )

//return it as HTML string
return html.String()
```

And yes, one more time: **Thank you, Drupal!**
