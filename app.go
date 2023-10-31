package main

import (
	"embed"
	"github.com/tylermmorton/create-torque-app/routes/landing"
	"github.com/tylermmorton/torque"
	"io/fs"
	"log"
	"net/http"
)

//go install github.com/tylermmorton/tmpl/cmd/tmpl@latest
//go:generate tmpl bind ./...

//go:embed .build/assets
var embeddedAssets embed.FS

func main() {
	staticAssets, err := fs.Sub(embeddedAssets, ".build/assets")
	if err != nil {
		log.Fatalf("failed to create static assets filesystem: %+v", err)
	}

	r := torque.NewRouter(
		torque.WithRouteModule("/", &landing.RouteModule{}),
		torque.WithFileSystemServer("/s", staticAssets),
	)

	log.Printf("[torque] Listening on http://localhost:8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("failed to start server: %+v", err)
	}
}
