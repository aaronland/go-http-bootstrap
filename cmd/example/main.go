package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"log"
	"net/http"
)

func Handler() http.Handler {

	index := `
<!doctype html>
<html lang="en-us">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <title>Bootstrap</title>
  </head>

  <body>
   <div class="card">
   	<h1 class="card-header">Card header</h1>
	<div class="card-body">Card body</div>
	<div class="card-footer">Card footer</div>
   </div>
  </body>
</html>`

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		rsp.Write([]byte(index))
	}

	return http.HandlerFunc(fn)
}

func main() {

	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")

	flag.Parse()

	mux := http.NewServeMux()

	idx_handler := Handler()

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()
	idx_handler = bootstrap.AppendResourcesHandler(idx_handler, bootstrap_opts)

	mux.Handle("/", idx_handler)

	err := bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatal(err)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening for requests on %s\n", endpoint)

	err = http.ListenAndServe(endpoint, mux)

	if err != nil {
		log.Fatal(err)
	}
}
