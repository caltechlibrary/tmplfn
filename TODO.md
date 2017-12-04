
# Action Items

## Next

+ [ ] unslug, slug
+ [ ] english_title
+ [ ] Documentation, examples, tutorials of using tmplfn in Go as well as the functions in text/templates
    + [ ] Basic text/templates built-in functionality
    + [ ] Funtionality added by tmplfn

## Someday, Maybe

+ [ ] Re-organize code so tmplfn.go holds the text template mappings but function collections are their own packages
+ [ ] Review other template function systems, align with their names and parameters where it makes sense
+ [ ] Implement a simpler dotpath function than how Go template's index function works
+ [ ] Add a codeblock function that will read in a file (with optional line range) and render a code example like tripple back tick does in Markdown

## Ideas

look at go-chart and see if it would be useful for generating embedded SVG output in Go templates.

the markdown fns should support news paper like columns from a single markdown doc. this could split on a specified html element like div or class the output with inline CSS for reflowing content. need to research current CSS3 support for the CSS solution.
