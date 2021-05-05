# go-http-bootstrap

`go-http-bootstrap` is an HTTP middleware package for including Bootstrap (v5.0.0) assets in web applications.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/go-http-bootstrap.svg)](https://pkg.go.dev/github.com/aaronland/go-http-bootstrap)

`go-http-bootstrap` is an HTTP middleware package for including Bootstrap.js assets in web applications. It exports two principal methods:

* `bootstrap.AppendAssetHandlers(*http.ServeMux)` which is used to append HTTP handlers to a `http.ServeMux` instance for serving Bootstrap CSS and JavaScript files, and related assets.
* `bootstrap.AppendResourcesHandler(http.Handler, *BootstrapOptions)` which is used to rewrite any HTML produced by previous handler to include the necessary markup to load Bootstrap

This package doesn't specify any code or methods for how Bootstrap.js is used. It only provides method for making Bootstraps available to existing applications.

## Example

See [cmd/example](cmd/example/main.go) for a working example.

## See also

* https://getbootstrap.com/
