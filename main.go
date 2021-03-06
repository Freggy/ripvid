package main

import (
	"flag"
	"fmt"
	"github.com/freggy/ripvid/youtube"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ip := flag.String("address", "localhost", "server address")
	port := flag.Int("port", 1337, "server port")
	router := mux.NewRouter()

	router.HandleFunc("/{id}/video", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, youtube.DownloadVideo)
	})

	router.HandleFunc("/{id}/audio", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r, youtube.DownloadAudio)
	})

	flag.Parse()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *ip, *port), router))
}

func handle(
	w http.ResponseWriter,
	r *http.Request,
	retrieve func(string, string, string) (string, error),
) {
	format := r.FormValue("format")
	filename := r.FormValue("filename")
	id := mux.Vars(r)["id"]

	if filename == "" {
		filename = id
	}

	if format == "" {
		e(w, http.StatusBadRequest, "format is missing")
		return
	}

	path, err := retrieve(filename, format, id)
	if err != nil {
		e(w, http.StatusInternalServerError, fmt.Sprintf("retrieve: %v", err))
		return
	}

	f, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		e(w, http.StatusInternalServerError, fmt.Sprintf("open file: %v", err))
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.%s", filename, format))
	http.ServeContent(w, r, filename, time.Now(), f)

	if err := os.Remove(path); err != nil {
		log.Printf("remove: %v\n", err)
	}
}

func e(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"error":{"msg":"%s"}}`, msg)
}
