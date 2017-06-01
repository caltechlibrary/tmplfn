
# Template Basics

_tmplfn_ wraps the Go template libraries and adds additional functions.
Go's template packages are is similar to [Mustache]() and [Handlebars]()
template systems. Those elements delimited by an opening `{{` and closing
`}}` will be evalauted by the template engine.  Typically you put a
variable name there where a period represents the root object holding
any variable content. This is easier to see with an example.

If our data is 

```json
    { "Name": "Robert" }
```

And our template is

```
    Hello {{ .Name }}
```

Assembling and executing our template would result in

```
    Hello Robert
```

## Go template additional compared to handlebars and mustache

The basic Go template system provides additional functionality.
This in includes the ability to use simple logic (e.g. _if_, _else_
and _with_) as well as to iterate over array or objects (e.g. 
_range_). The basic template system also allows you to define 
and use sub-templates and blocks (e.g. _define_, _template_, _block_).
This makes it easier to share templated element.



