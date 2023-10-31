package html

//tmpl:bind page.tmpl.html
type Page struct {
	Title string
	// TitlePrefix is prepended to the title if set
	TitlePrefix string
	// TitleSuffix is appended to the title if set
	TitleSuffix string

	Links   []Link   `tmpl:"link"`
	Scripts []Script `tmpl:"script"`
}
