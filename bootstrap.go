package bootstrap

import (
	"github.com/aaronland/go-http-bootstrap/resources"
	_ "log"
	"net/http"
	"path/filepath"
	"strings"
)

type BootstrapOptions struct {
	JS  []string
	CSS []string
}

func DefaultBootstrapOptions() *BootstrapOptions {

	opts := &BootstrapOptions{
		CSS: []string{"/css/bootstrap.min.css"},
		JS:  []string{"/javascript/bootstrap.min.js"},
	}

	return opts
}

func AppendResourcesHandler(next http.Handler, opts *BootstrapOptions) http.Handler {
	return AppendResourcesHandlerWithPrefix(next, opts, "")
}

func AppendResourcesHandlerWithPrefix(next http.Handler, opts *BootstrapOptions, prefix string) http.Handler {

	js := opts.JS
	css := opts.CSS

	if prefix != "" {

		for i, path := range js {
			js[i] = appendPrefix(prefix, path)
		}

		for i, path := range css {
			css[i] = appendPrefix(prefix, path)
		}
	}

	ext_opts := &resources.AppendResourcesOptions{
		JS:  js,
		CSS: css,
	}

	return resources.AppendResourcesHandler(next, ext_opts)
}

func AssetsHandler() (http.Handler, error) {

	fs := assetFS()
	return http.FileServer(fs), nil
}

func AppendAssetHandlers(mux *http.ServeMux) error {
	return AppendAssetHandlersWithPrefix(mux, "")
}

func AppendAssetHandlersWithPrefix(mux *http.ServeMux, prefix string) error {

	asset_handler, err := AssetsHandler()

	if err != nil {
		return nil
	}

	for _, path := range AssetNames() {

		path := strings.Replace(path, "static", "", 1)

		if prefix != "" {
			path = appendPrefix(prefix, path)
		}

		mux.Handle(path, asset_handler)
	}

	return nil
}

func appendPrefix(prefix string, path string) string {

	prefix = strings.TrimRight(prefix, "/")

	if prefix != "" {
		path = strings.TrimLeft(path, "/")
		path = filepath.Join(prefix, path)
	}

	return path
}
