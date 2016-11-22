/**
 * tmpllib are a collection of functions useful to add to the default Go template/text and template/html template definitions
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
package tmplfuncs

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
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
	}
)

// normalizeDate takes a expands years to four digits, month and day to two digits
// E.g. 4/3/2016 becomes 04/03/2016
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
