package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//go:embed web/index.html
var webFS embed.FS

var indexTmpl = template.Must(template.ParseFS(webFS, "web/index.html"))

// maxCount caps a single generation request so an accidental huge value
// (or an abusive request) can't exhaust memory or time out the response.
const maxCount = 500000

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if err := indexTmpl.Execute(w, nil); err != nil {
		log.Println("template error:", err)
	}
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form", http.StatusBadRequest)
		return
	}

	start, err := strconv.ParseInt(strings.TrimSpace(r.FormValue("start")), 10, 64)
	if err != nil || start < 1 {
		http.Error(w, "некорректный начальный номер", http.StatusBadRequest)
		return
	}
	count, err := strconv.ParseInt(strings.TrimSpace(r.FormValue("count")), 10, 64)
	if err != nil || count < 1 {
		http.Error(w, "некорректное количество номеров", http.StatusBadRequest)
		return
	}
	if count > maxCount {
		http.Error(w, fmt.Sprintf("слишком много номеров (максимум %d)", maxCount), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="NEW-Pool.csv"`)
	if err := WritePool(w, start, count); err != nil {
		log.Println("generate error:", err)
	}
}

func main() {
	loadDotEnv(".env")
	port := getPort()

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/generate", generateHandler)

	addr := ":" + port
	log.Printf("gen-phones listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
