package swagger

import (
	"bytes"
	"html/template"
	"net/http"
)

type SwaggerOpts struct {
	SwaggerSpec []byte
	DocsURL     string
	SpecURL     string
	RedocURL    string
	Title       string
}

func ensureDefaults(opts *SwaggerOpts) {
	if opts.RedocURL == "" {
		opts.RedocURL = redocLatest
	}
	if opts.SpecURL == "" {
		opts.SpecURL = "/swagger.json"
	}
	if opts.DocsURL == "" {
		opts.DocsURL = "/docs"
	}
}

func WithSwagger(swagger SwaggerOpts, handler http.Handler) http.Handler {

	ensureDefaults(&swagger)
	tmpl := template.Must(template.New("redoc").Parse(redocTemplate))

	buf := bytes.NewBuffer(nil)
	_ = tmpl.Execute(buf, swagger)
	b := buf.Bytes()

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path == swagger.SpecURL {
			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			rw.WriteHeader(http.StatusOK)

			_, _ = rw.Write(swagger.SwaggerSpec)
			return
		}

		if r.URL.Path == swagger.DocsURL {
			rw.Header().Set("Content-Type", "text/html; charset=utf-8")

			rw.WriteHeader(http.StatusOK)

			_, _ = rw.Write(b)
			return
		}

		handler.ServeHTTP(rw, r)
	})
}

const (
	redocLatest   = "https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"
	redocTemplate = `<!DOCTYPE html>
<html>
  <head>
    <title>{{ .Title }}</title>
    <!-- needed for adaptive design -->
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!--
    ReDoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='{{ .SpecURL }}'></redoc>
    <script src="{{ .RedocURL }}"> </script>
  </body>
</html>
`
)
