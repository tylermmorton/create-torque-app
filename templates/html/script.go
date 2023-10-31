package html

import "html/template"

//tmpl:bind script.tmpl.html
type Script struct {
	Src       string
	Type      string
	Content   template.JS
	Integrity string
	Defer     bool
	Async     bool
}

// IsValid returns true if the script should be rendered.
func (s Script) IsValid() bool {
	return s.Src != "" || s.Content != ""
}

func (Script) Attr(name, value string) template.HTMLAttr {
	if value == "" {
		return "\b"
	}
	return template.HTMLAttr(name + "=\"" + value + "\"")
}
