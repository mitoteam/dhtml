# dhtml - Go html renderer

Go html renderer inspired by [Drupal](https://www.drupal.org/)'s renderable arrays concept. Allows to built some elements tree in Go and then render it as html.

NOT READY FOR PRODUCTION YET!

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
      // dhtml.Div() is an alias for dhtml.Tag("div")
      dhtml.Div().Id("test").
        Attribute("data-mt-test", "some attribute").
        Class("border").Class("m-3").Class("p-3"). //css classes handling
        Content("some <escaped> text"),
    ),
  )

//return it as HTML string
return html.Render()
```

And yes, one more time: **Thank you, Drupal!**
