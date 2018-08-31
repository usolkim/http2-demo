package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("../wwwroot")))

	http.HandleFunc("/server_push.html", func(w http.ResponseWriter, r *http.Request) {
		if pusher, ok := w.(http.Pusher); ok {
			// Push is supported.

			push(pusher, "/js/vue.js")
			push(pusher, "/js/element-ui.js")
			push(pusher, "/css/element-ui.css")
			push(pusher, "/js/jquery.js")
			push(pusher, "/images/2MB.jpg")
		}

		content, err := ioutil.ReadFile("../wwwroot/server_push.html")
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(http.StatusText(404)))
			return
		}
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Etag", strconv.Itoa(rand.Int()))
		w.Header().Add("Content-Type", "text/html")
		w.Write(content)
	})

	go func() {
		log.Fatal(http.ListenAndServe(":80", logRequest(http.DefaultServeMux)))
	}()

	log.Fatal(http.ListenAndServeTLS(":443", "../cert/localhost.cert", "../cert/localhost.key", logRequest(http.DefaultServeMux)))
}

func push(pusher http.Pusher, resource string) {
	if err := pusher.Push(resource, nil); err != nil {
		log.Printf("Failed to push: %s (%v)", resource, err)
	} else {
		log.Printf("Success to push: %s\n", resource)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.Header.Get("User-Agent"))
		c, err := r.Cookie("latency")
		if err == nil {
			t := c.Value
			if t != "0ms" {
				v, err := time.ParseDuration(t)
				if err == nil {
					time.Sleep(v)
				}
			}
		}
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Etag", strconv.Itoa(rand.Int()))
		handler.ServeHTTP(w, r)
	})
}
