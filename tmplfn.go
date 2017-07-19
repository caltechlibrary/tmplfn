/**
 * tmplfn are a collection of functions useful to add to the default Go template/text and template/html template definitions
 *
 * @author R. S. Doiel
 *
 * Copyright (c) 2017, Caltech
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
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"

	// Caltech Library Packages
	"github.com/caltechlibrary/dotpath"
	"github.com/caltechlibrary/tmplfn/numbers"
)

var (
	// Version of tmplfn package
	Version = `v0.0.15`

	// Time provides a common set of time/date related functions for use in text/template or html/template
	Time = template.FuncMap{
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

	Math = template.FuncMap{
		"int":      numbers.Int,
		"int64":    numbers.Int64,
		"float32":  numbers.Float32,
		"float64":  numbers.Float64,
		"add":      numbers.Add,
		"sub":      numbers.Subtract,
		"multiply": numbers.Multiply,
		"divide":   numbers.Divide,
		"modulo":   numbers.Modulo,
		"addi":     numbers.Addi,
		"subi":     numbers.Subtract,
		"typeof": func(t interface{}) string {
			if t == nil {
				return "<nil>"
			}
			return fmt.Sprintf("%T", t)
		},
	}

	Strings = template.FuncMap{
		// concat concatenates strings together
		"concat": func(strs ...string) string {
			return strings.Join(strs, "")
		},
		// has_prefix returns true if prefix matches, false otherwise
		"has_prefix":  strings.HasPrefix,
		"has_suffix":  strings.HasSuffix,
		"contains":    strings.Contains,
		"trim_prefix": strings.TrimPrefix,
		"trim_suffix": strings.TrimSuffix,
		"lowercase":   strings.ToLower,
		"uppercase":   strings.ToUpper,
		"title":       strings.Title,
		"replace":     strings.Replace,
		// join joins an array of strings with separator
		"join": func(li []interface{}, sep string) string {
			var l []string
			for _, item := range li {
				if item == nil {
					l = append(l, "")
				} else {
					l = append(l, fmt.Sprintf("%s", item))
				}
			}
			return strings.Join(l, sep)
		},
	}

	Page = template.FuncMap{
		"nl2p": func(s string) string {
			return strings.Replace(strings.Replace(s, "\n\n", "<p>", -1), "\n", "<br />", -1)
		},
		"previ": func(pos, move_size, min_pos, max_pos int, wrap bool) int {
			prev := pos - move_size
			if prev < min_pos {
				if wrap == false {
					return min_pos
				}
				return max_pos
			}
			return prev
		},
		"nexti": func(pos, move_size, min_pos, max_pos int, wrap bool) int {
			next := pos + move_size
			if next > max_pos {
				if wrap == false {
					return pos
				}
				return min_pos
			}
			return next
		},
		"synopsis": func(s string) string {
			return doc.Synopsis(s)
		},
		"urldecode": func(s string) string {
			sDecoded, err := url.QueryUnescape(s)
			if err != nil {
				log.Printf("Bad encoding request: %q, %s\n", s, err)
				return ""
			}
			return sDecoded
		},
		"urlencode": func(s string) string {
			return url.QueryEscape(s)
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

	// Iterables produces lists that then can supply the template range function with values
	Iterables = template.FuncMap{
		// ints returns an array of int. Both start and end are inclusive. If start <= end the ascending by inc else descending by inc
		"ints": func(start, end, inc int) []int {
			var result []int
			if start == end {
				return []int{start}
			} else if start < end {
				for i := start; i <= end; i = i + inc {
					result = append(result, i)
				}
			} else {
				for i := end; i >= start; i = i - inc {
					result = append(result, i)
				}
			}
			return result
		},
		// int64s returns an array of int64. Both start and end are inclusive. If start <= end the ascending by inc else descending by inc
		"int64s": func(start, end, inc int64) []int64 {
			var result []int64
			if start == end {
				return []int64{start}
			} else if start < end {
				for i := start; i <= end; i = i + inc {
					result = append(result, i)
				}
			} else {
				for i := end; i >= start; i = i - inc {
					result = append(result, i)
				}
			}
			return result
		},
		// float32s returns an array of float32. Both start and end are inclusive. If start <= end the ascending by inc else descending by inc
		"float32s": func(start, end, inc float32) []float32 {
			var result []float32
			if start == end {
				return []float32{start}
			} else if start < end {
				for i := start; i <= end; i = i + inc {
					result = append(result, i)
				}
			} else {
				for i := end; i >= start; i = i - inc {
					result = append(result, i)
				}
			}
			return result
		},
		// float64s returns an array of float64. Both start and end are inclusive. If start <= end the ascending by inc else descending by inc
		"float64s": func(start, end, inc float64) []float64 {
			var result []float64
			if start == end {
				return []float64{start}
			} else if start < end {
				for i := start; i <= end; i = i + inc {
					result = append(result, i)
				}
			} else {
				for i := end; i >= start; i = i - inc {
					result = append(result, i)
				}
			}
			return result
		},
		// cols2rows takes a list of columns and returns a 2d array of rows and columns
		// number of rows will match the largest number of cells in the columns included, empty/missing cells will
		// be added using an empty string
		"cols2rows": func(cols ...[]interface{}) [][]interface{} {
			var (
				row     []interface{}
				rows    [][]interface{}
				maxRows int
			)
			// Find the max rows
			for _, col := range cols {
				if len(col) >= maxRows {
					maxRows = len(col)
				}
			}
			// From zero to maxRows assemble a row and add to rows
			for i := 0; i < maxRows; i++ {
				// reset row
				row = []interface{}{}
				// build row
				for _, col := range cols {
					if i < len(col) {
						row = append(row, col[i])
					} else {
						row = append(row, "")
					}
				}
				// add row to rows
				rows = append(rows, row)
			}
			// For each column add a cell to the row
			return rows
		},
	}

	//Booleans provides a set of functions working with Boolean data
	Booleans = template.FuncMap{
		// CountTrue takes one or more booling variables and counts the ones which are true.
		// returns an integer
		"count_true": func(booleans ...bool) int {
			cnt := 0
			for _, val := range booleans {
				if val == true {
					cnt++
				}
			}
			return cnt
		},
	}

	//Path methods for working with paths (E.g. path.Base(), path.Ext() and path.Dir() in Go path package)
	Path = template.FuncMap{
		// basename works similar to Unix command and will trim any extensions provided in additional to the path
		"basename": func(args ...string) string {
			p := path.Base(args[0])
			if len(args) > 1 {
				for _, ext := range args[1:] {
					if strings.HasSuffix(p, ext) {
						p = strings.TrimSuffix(p, ext)
					}
				}
			}
			return p
		},
		"base": path.Base,
		"ext":  path.Ext,
		"dir":  path.Dir,
	}

	//Url methods are for working with URLs and extracting useful parts
	Url = template.FuncMap{
		// return the scheme (e.g. https, http) of URL
		"url_scheme": func(args ...string) string {
			if u, err := url.Parse(args[0]); err == nil {
				return u.Scheme
			}
			return ""
		},
		// return the host (e.g. example.org ) of URL
		"url_host": func(args ...string) string {
			if u, err := url.Parse(args[0]); err == nil {
				return u.Host
			}
			return ""
		},
		// return the path (e.g. /about/index.html ) of URL
		"url_path": func(args ...string) string {
			if u, err := url.Parse(args[0]); err == nil {
				return u.Path
			}
			return ""
		},
	}

	//Dotpath methods from datatools/dotpath in templates
	Dotpath = template.FuncMap{
		//dotpath takes a dot path (string), the data to operate on (e.g. map[string]interface{}) and default data to return on fail,
		// dotpath returns the dot path result on success or the default fail value if not.
		"dotpath": func(p string, data interface{}, defaultVal interface{}) interface{} {
			if val, err := dotpath.Eval(p, data); err == nil {
				return val
			}
			return defaultVal
		},
		// has_dotpath takes a dot path (string), the data to check, the value to return if DOES exists and the value to return if NOT exists)
		// the return value will be either the exists value or does not exist value
		"has_dotpath": func(p string, data interface{}, existsVal interface{}, notExistsVal interface{}) interface{} {
			if _, err := dotpath.Eval(p, data); err == nil {
				return existsVal
			}
			return notExistsVal
		},
	}

	// Console holds functions that interact with the console where
	// the template processing is happening (e.g. for a web service
	// writing to the console log of the web server).
	Console = template.FuncMap{
		// writelog writes something to the log using Log.Println()
		// it returns an empty string because template Funcs need to
		// return something.
		"writelog": func(v ...interface{}) string {
			log.Println(v...)
			return ""
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
	result := template.FuncMap{}
	for _, m := range maps {
		for key, fn := range m {
			result[key] = fn
		}
	}
	return result
}

// AllFuncs() returns a Join of func maps available in tmplfn
func AllFuncs() template.FuncMap {
	return Join(Booleans, Console, Dotpath, Iterables, Math, Page, Markdown, Path, Strings, Time, Url)
}

// Src is a mapping of template source to names and byte arrays.
// It is useful to create a series of defaults templates that can be
// overwritten by user supplied template files.
type Tmpl struct {
	// Holds the function map for templates
	FuncMap template.FuncMap

	// Code holds a map of names to byte arrays, the byte arrays hold the template source code
	// the names can be either filename or other names defined by the implementor
	Code map[string][]byte
}

// New creates a pointer to a template.Template and  empty map of names to byte arrays
// pointing at an empty byte array
func New(fm template.FuncMap) *Tmpl {
	return &Tmpl{
		FuncMap: fm,
		Code:    map[string][]byte{},
	}
}

// ReadFiles takes the given file, filenames or directory name(s) and reads the byte array(s)
// into the Code map.  If a filename is a directory then the directory is scanned
// for files ending in ".tmpl" and those are loaded into the Code map. It does
// NOT parse/assemble templates. The basename in the path is used as the name
// of the template (e.g. templates/page.tmpl would be stored as page.tmpl.
func (t *Tmpl) ReadFiles(fNames ...string) error {
	for _, fname := range fNames {
		if info, err := os.Stat(fname); err != nil {
			return err
		} else if info.IsDir() == true {
			if files, err := ioutil.ReadDir(fname); err == nil {
				for _, file := range files {
					tname := path.Base(file.Name())
					pname := path.Join(fname, file.Name())
					ext := path.Ext(pname)
					if file.IsDir() != true && ext == ".tmpl" {
						if src, err := ioutil.ReadFile(pname); err != nil {
							return err
						} else {
							t.Code[tname] = src
						}
					}
				}
			} else {
				return err
			}
		} else if src, err := ioutil.ReadFile(fname); err == nil {
			tname := path.Base(fname)
			t.Code[tname] = src
		} else {
			return err
		}
	}
	return nil
}

// Add takes a name and source (byte array) and updates t.Code with it.
// It is like Merge but for a single file. The name provided in Add is
// used as the key to the template source code map.
func (t Tmpl) Add(name string, src []byte) error {
	t.Code[name] = src
	if _, ok := t.Code[name]; ok != true {
		return fmt.Errorf("failed to add %s", name)
	}
	return nil
}

// ReadMap works like ReadFiles but takes the name/source pairs from a map rather
// than the file system. It expected template names to end in ".tmpl" like ReadFiles()
// Note the basename of the key provided in the sourceMap is used as the key
// in the Code source code map (e.g. /templates/page.tmpl is stored as page.tmpl)
func (t Tmpl) ReadMap(sourceMap map[string][]byte) error {
	for fname, src := range sourceMap {
		tname := path.Base(fname)
		ext := path.Ext(tname)
		if ext == ".tmpl" {
			t.Code[tname] = src
		}
	}
	if len(t.Code) == 0 {
		return fmt.Errorf("No templates found")
	}
	return nil
}

// Assemble mimics template.ParseFiles() but works with the properties of
// a Tmpl struct.
func (t Tmpl) Assemble() (*template.Template, error) {
	if len(t.Code) == 0 {
		// Mimmic template.ParseFiles() error
		return nil, fmt.Errorf("tmplfn.Assemble(): no template sources to parse")
	}
	var tpl *template.Template
	// Scan the individual templates and parse errors.
	for tName, tSrc := range t.Code {
		s := string(tSrc)
		name := path.Base(tName)
		// This is patterned after template.ParseFiles() internal calls to parseFiles()
		// First template becomes return value if not already defined,
		// use subsequent New calls to associate all the templates together.
		// Otherwise we create a new template associated with t.Template
		var tmpl *template.Template
		if tpl == nil {
			tpl = template.New(name).Funcs(t.FuncMap)
		}
		if name == tpl.Name() {
			tmpl = tpl
		} else {
			tmpl = tpl.New(name).Funcs(t.FuncMap)
		}
		if _, err := tmpl.Parse(s); err != nil {
			return nil, err
		}
	}
	return tpl, nil
}
