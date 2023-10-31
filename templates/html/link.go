package html

type RelType string

const (
	RelTypeAlternate   RelType = "alternate"
	RelTypeAuthor      RelType = "author"
	RelTypeDnsPrefetch RelType = "dns-prefetch"
	RelTypeHelp        RelType = "help"
	RelTypeIcon        RelType = "icon"
	RelTypeLicense     RelType = "license"
	RelTypeNext        RelType = "next"
	RelTypePingback    RelType = "pingback"
	RelTypePreconnect  RelType = "preconnect"
	RelTypePrefetch    RelType = "prefetch"
	RelTypePreload     RelType = "preload"
	RelTypePrerender   RelType = "prerender"
	RelTypePrev        RelType = "prev"
	RelTypeSearch      RelType = "search"
	RelTypeStylesheet  RelType = "stylesheet"
)

// Link represents an HTML linked document
// https://www.w3schools.com/tags/tag_link.asp
//
//tmpl:bind link.tmpl.html
type Link struct {
	// Href specifies the location of the linked document
	Href string
	// Media specifies on what device the linked document will be displayed
	Media *string
	// Rel specifies the relationship between the current document and the linked document
	Rel RelType
	// Type specifies the media type of the linked document
	Type string
	// As specifies the type of content being loaded
	As string
	// CrossOrigin specifies how the element handles crossorigin requests
	CrossOrigin string
}
