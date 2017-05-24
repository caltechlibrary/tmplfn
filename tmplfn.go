/**
 * tmplfn are a collection of functions useful to add to the default Go template/text and template/html template definitions
 *
 * @author R. S. Doiel
 *
 * Copyright (c) 2016, Caltech
 * All rights not granted herein are expressly reserved by Caltech.
 *
 * Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 *
 * 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */
package tmplfn

import (
	"encoding/json"
	"fmt"
	"go/doc"
	"log"
	"net/url"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
	// Version of tmplfn package
	Version = `v0.0.5`

	// TimeMap provides a common set of time/date related functions for use in text/template or html/template
	TimeMap = template.FuncMap{
		"year": func(s string) string {
			var (
				dt  time.Time
				err error
			)
			if s == "now" {
				dt = time.Now()
			} else {
				dt, err = time.Parse("2006-01-02", normalizeDate(s))
				if err != nil {
					return ""
				}
			}
			return dt.Format("2006")
		},
		"rfc3339": func(s string) string {
			var (
				dt  time.Time
				err error
			)
			if s == "now" {
				dt = time.Now()
			} else {
				dt, err = time.Parse("2006-01-02", normalizeDate(s))
				if err != nil {
					return ""
				}
			}
			return dt.Format(time.RFC3339)
		},
		"rfc1123": func(s string) string {
			var (
				dt  time.Time
				err error
			)
			if s == "now" {
				dt = time.Now()
			} else {
				dt, err = time.Parse("2006-01-02", normalizeDate(s))
				if err != nil {
					return ""
				}
			}
			return dt.Format(time.RFC1123)
		},
		"rfc1123z": func(s string) string {
			var (
				dt  time.Time
				err error
			)
			if s == "now" {
				dt = time.Now()
			} else {
				dt, err = time.Parse("2006-01-02", normalizeDate(s))
				if err != nil {
					return ""
				}
			}
			return dt.Format(time.RFC1123Z)
		},
		"rfc822z": func(s string) string {
			var (
				dt  time.Time
				err error
			)
			if s == "now" {
				dt = time.Now()
			} else {
				dt, err = time.Parse("2006-01-02", normalizeDate(s))
				if err != nil {
					return ""
				}
			}
			return dt.Format(time.RFC822Z)
		},
		"rfc822": func(s string) string {
			var (
				dt  time.Time
				err error
			)
			if s == "now" {
				dt = time.Now()
			} else {
				dt, err = time.Parse("2006-01-02", normalizeDate(s))
				if err != nil {
					return ""
				}
			}
			return dt.Format(time.RFC822)
		},
		"datefmt": func(dt, outputFmtYMD, outputFmtYM, outputFmtY string) string {
			var (
				inputFmt  string
				outputFmt string
			)
			// NOTE: Date input formats can be YYYY, YYYY-MM and YYYY,
			// we need to define output formats for each
			switch {
			case len(dt) == 4:
				inputFmt = "2006"
				outputFmt = outputFmtY
			case len(dt) > 4 && len(dt) <= 7:
				inputFmt = "2006-01"
				outputFmt = outputFmtYM
			default:
				inputFmt = "2006-01-02"
				outputFmt = outputFmtYMD
			}
			//intputFmt: 2006-01-02
			//outputfmt: Jan _2, 2006
			d, err := time.Parse(inputFmt, dt)
			if err != nil {
				return fmt.Sprintf("%s, %s", dt, err.Error())
			}
			return d.Format(outputFmt)
		},
	}

	PageMap = template.FuncMap{
		"nl2p": func(s string) string {
			return strings.Replace(strings.Replace(s, "\n\n", "<p>", -1), "\n", "<br />", -1)
		},
		"contains": func(s, substring string) bool {
			return strings.Contains(s, substring)
		},
		"title": func(s string) string {
			return strings.Title(s)
		},
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"arraylength": func(a []string) int {
			return len(a)
		},
		"mapsize": func(m map[string]string) int {
			return len(m)
		},
		"prevPage": func(from, size, max int) int {
			next := from - size
			if next < 0 {
				return 0
			}
			return next
		},
		"nextPage": func(from, size, max int) int {
			next := from + size
			if next > max {
				return from
			}
			return next
		},
		"getType": func(t interface{}) string {
			switch tp := t.(type) {
			default:
				return fmt.Sprintf("%T", tp)
			}
		},
		"asList": func(li []interface{}, sep string) string {
			var l []string
			for _, item := range li {
				l = append(l, fmt.Sprintf("%s", item))
			}
			return strings.Join(l, sep)
		},
		"synopsis": func(s string) string {
			return doc.Synopsis(s)
		},
		"encodeURIComponent": func(s string) string {
			u, err := url.Parse(s)
			if err != nil {
				log.Printf("Bad encoding request: %s, %s\n", s, err)
				return ""
			}
			return strings.Replace(u.String(), "&", "%26", -1)
		},
		"stringify": func(data interface{}, prettyPrint bool) string {
			if prettyPrint == true {
				if buf, err := json.MarshalIndent(data, "", "\t"); err == nil {
					return string(buf)
				}
			} else if buf, err := json.Marshal(data); err == nil {
				return string(buf)
			}
			return ""
		},
		"codeblock": func(src string, start int, end int, hint string) string {
			result := []string{}
			lines := strings.Split(src, "\n")
			if start < 1 {
				start = 0
			}
			if end < 1 {
				end = len(lines)
			}
			if (end - start) > 0 {
				result = append(result, fmt.Sprintf("```%s", hint))
			}
			for i, line := range lines[start:end] {
				if len(line) > 0 {
					result = append(result, fmt.Sprintf("    %s", line))
				} else if i > 0 && i < (end-1) {
					result = append(result, "")
				}
			}
			if len(result) > 0 {
				result = append(result, "```")
			}
			return strings.Join(result, "\n")
		},
	}
)

// normalizeDate takes a expands years to four digits, month and day to two digits
// E.g. 2016-4-3 becomes 2016-04-03
func normalizeDate(in string) string {
	parts := strings.Split(in, "-")
	if len(parts) == 1 {
		parts = append(parts, "01")
		parts = append(parts, "01")
	}
	if len(parts) == 2 {
		parts = append(parts, "01")
	}
	for i := 0; i < len(parts); i++ {
		x, err := strconv.Atoi(parts[i])
		if err != nil {
			x = 1
		}
		if i == 0 {
			parts[i] = fmt.Sprintf("%0.4d", x)
		} else {
			parts[i] = fmt.Sprintf("%0.2d", x)
		}
	}
	return strings.Join(parts, "-")
}

// Join take one or more func maps and returns an aggregate one.
func Join(maps ...template.FuncMap) template.FuncMap {
	result := make(template.FuncMap)
	for _, m := range maps {
		for key, fn := range m {
			result[key] = fn
		}
	}
	return result
}

// Assemble support a very simple template setup of an outer HTML file with a content include
// along with common template functions.
func Assemble(tmplFuncs template.FuncMap, templateFilenames ...string) (*template.Template, error) {
	if len(templateFilenames) > 0 {
		return template.New(path.Base(templateFilenames[0])).Funcs(tmplFuncs).ParseFiles(templateFilenames...)
	}
	return nil, fmt.Errorf("No template names specified")
}

// AssembleString like Assemble but using a string as a source for the template
func AssembleString(tmplFuncs template.FuncMap, src string) (*template.Template, error) {
	return template.New("master").Funcs(tmplFuncs).Parse(src)
}

// AssembleTemplateMap takes a map of template names and source and returns a map of template names and templates
func AssembleTemplateMap(tmplFuncs template.FuncMap, templateSrcMap map[string]string) (map[string]*template.Template, error) {
	tmpls := map[string]*template.Template{}
	for tName, tSrc := range templateSrcMap {
		if tmpl, err := template.New(tName).Funcs(tmplFuncs).Parse(src); err == nil {
			tmpls[tName] = tmpl
		} else {
			return nil, fmt.Errorf("%s parse error, %s", tName, err)
		}
	}
	return tmpls, nil
}
