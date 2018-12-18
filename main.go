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
	    <style>
		a {
		  color: white;
		}
	    .centered {
	      position: fixed;
	      top: 50%;
	      left: 50%;
	      margin-top: -150px;
	      margin-left: -111px;
	    }
		.footer {
		  position: fixed;
		  font-family: Futura,Trebuchet MS,Arial,sans-serif;
		  left: 0;
		  bottom: 0;
		  width: 100%;
		  background-color: #3367d6;
		  color: white;
		  text-align: center;
		}
	    </style>
		<link rel="shortcut icon" href="https://storage.googleapis.com/images.schwag.exchange/favicon.ico" type="image/x-icon">
		<link rel="icon" href="https://storage.googleapis.com/images.schwag.exchange/favicon.ico" type="image/x-icon">
	    <!-- Global site tag (gtag.js) - Google Analytics -->
	    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-131145973-1"></script>
	    <script>
	        window.dataLayer = window.dataLayer || [];
	        function gtag(){dataLayer.push(arguments);}
	        gtag('js', new Date());
	        gtag('config', 'UA-131145973-1');
	    </script>
	    <title>Swag Exchange</title>
	    </head>
	    <body>
	    <!--  -->
	    <!-- Follow the progress at  -->

	    <img class="centered" src="https://storage.googleapis.com/images.schwag.exchange/schwaggy.png" alt="Schwaggy the Hamster" />

		<div class="footer">
		  <p>"<a href="https://github.com/betandr/schwag_exchange">Schwag Exchange</a>"
		  is a reference Go/Kubernetes/Istio application by <a href="https://bet.andr.io/about/">Beth Anderson</a>.</p>
		</div>
	    </body>
	    </html>`

	var issueList = template.Must(template.New("swag").Parse(templ))

	if err := issueList.Execute(out, nil); err != nil {
		log.Fatal(err)
	}
}
