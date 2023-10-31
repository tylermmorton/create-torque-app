package landing

import (
	"github.com/tylermmorton/create-torque-app/templates/html"
	"github.com/tylermmorton/tmpl"
	"github.com/tylermmorton/torque"
	"net/http"
)

// Template can be used to render the landing page.
var Template = tmpl.MustCompile(&LandingPage{})

//tmpl:bind landing.tmpl.html
type LandingPage struct {
	// Page is a template for a base html page.
	// It exposes the `body` template slot.
	html.Page `tmpl:"page"` // <- name the template, so it can be used as a target
}

// RouteModule is the torque route module for the landing page.
type RouteModule struct{}

var _ interface {
	torque.Loader
	torque.Renderer
} = &RouteModule{}

func (rm *RouteModule) Load(req *http.Request) (any, error) {
	return nil, nil
}

func (rm *RouteModule) Render(wr http.ResponseWriter, req *http.Request, loaderData any) error {
	return Template.Render(wr,
		&LandingPage{
			Page: html.Page{
				TitlePrefix: "Welcome",
				Title:       "create-torque-app",
				Links: []html.Link{
					{Rel: "stylesheet", Href: "/s/app.css"},
				},
				Scripts: []html.Script{
					{Src: "https://unpkg.com/htmx.org@1.9.6"},
				},
			},
		},
		tmpl.WithName("body"),   // <- assign the landing page template to the `body` slot
		tmpl.WithTarget("page"), // <- render the `page` template, which contains the `body`
	)
}
