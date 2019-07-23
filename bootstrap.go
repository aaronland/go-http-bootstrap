package bootstrap

import (
	"github.com/aaronland/go-http-bootstrap/resources"
	_ "log"
	"net/http"
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

	ext_opts := &resources.AppendResourcesOptions{
		JS:  opts.JS,
		CSS: opts.CSS,
	}

	return resources.AppendResourcesHandler(next, ext_opts)
}

func AssetsHandler() (http.Handler, error) {

	fs := assetFS()
	return http.FileServer(fs), nil
}

func AppendAssetHandlers(mux *http.ServeMux) error {

	asset_handler, err := AssetsHandler()

	if err != nil {
		return nil
	}

	for _, path := range AssetNames() {
		path := strings.Replace(path, "static", "", 1)
		mux.Handle(path, asset_handler)
	}

	return nil
}
