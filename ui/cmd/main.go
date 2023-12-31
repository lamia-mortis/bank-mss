package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"ui/resources"
)

func main() {
	host, found := os.LookupEnv("UI_HOST")

	if !found {
		host = "127.0.0.1"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:8080", host),
		Handler: router(),
	}

	srv.ListenAndServe()
}

func router() http.Handler {
	mux := http.NewServeMux()

	// index
	mux.HandleFunc("/", indexHandler)

	// static files
	staticFS, _ := fs.Sub(resources.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", httpFS)

	// some api
	mux.HandleFunc("/api/v1/greeting", greetingAPI)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}

	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := resources.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}

	rawFile, _ := resources.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}

func greetingAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, there!"))
}
