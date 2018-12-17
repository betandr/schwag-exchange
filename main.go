package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		index(w)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
	return
}

func index(out io.Writer) {
	templ := `<!DOCTYPE html>
        <html>
        <head>
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <style>
        body, html {
          height: 100%;
          margin: 0;
        }

        .bg {
          background-image: url("https://storage.googleapis.com/images.schwag.exchange/coming_soon.jpg");
          height: 100%;
          background-position: center;
          background-repeat: no-repeat;
          background-size: cover;
        }
        </style>
        <title>Swag Exchange</title>
        </head>
        <body>

        <div class="bg"></div>

        <!-- Follow the progress of Schwag Exchange at https://github.com/betandr/schwag_exchange -->

        </body>
        </html>`

	var issueList = template.Must(template.New("swag").Parse(templ))

	if err := issueList.Execute(out, nil); err != nil {
		log.Fatal(err)
	}
}
