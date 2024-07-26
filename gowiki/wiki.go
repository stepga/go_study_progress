package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html", "index.html", "index_header.html"))
var port = 8080

type Page struct {
	Title string
	Body  []byte
	// XXX: ^ `io` libraries expect rather `[]byte` types instead of `string`
}

func (p *Page) filename() string {
	return filename(p.Title)
}

func filename(title string) string {
	return fmt.Sprintf("%s.txt", title)
}

func (p *Page) save() error {
	return os.WriteFile(p.filename(), p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	body, err := os.ReadFile(filename(title))
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	var port_str = fmt.Sprintf(":%d", port)
	log.Printf("[INFO] listening on port: %s\n", port_str)
	log.Fatalf("[ERROR] ListenAndServe: %v\n", http.ListenAndServe(port_str, nil))
}

func printHeader(w http.ResponseWriter, h http.Header) {
	keys := make([]string, 0)
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Fprintln(w, "Header:")
	for _, k := range keys {
		fmt.Fprintf(w, " %s: %v\n", k, h[k])
	}
}

func listFilesWithSuffix(suffix string) []string {
	dir_entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatalf("[ERROR] ReadDir(\"./\"): %v\n", err)
	}

	ret := make([]string, 0)
	for _, dir_entry := range dir_entries {
		if strings.HasSuffix(dir_entry.Name(), suffix) {
			fname := dir_entry.Name()
			fname = fname[:len(fname)-len(suffix)]
			ret = append(ret, fname)
		}
	}
	sort.Strings(ret)
	return ret
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	txt_file_names := listFilesWithSuffix(".txt")
	renderTemplate(w, "index_header", nil)
	for _, fname := range txt_file_names {
		renderTemplate(w, "index", &Page{Title: fname})
	}
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	request_path := r.URL.Path[1:]
	if len(request_path) > 0 {
		fmt.Fprintln(w, "Requested URL Path:")
		fmt.Fprintf(w, " '%s'\n", request_path)
	}
	printHeader(w, r.Header)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	if len(title) == 0 {
		http.Error(w, "A Page's minimum title length is 1", http.StatusBadRequest)
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	if len(title) == 0 {
		http.Error(w, "A Page's minimum title length is 1", http.StatusBadRequest)
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	// TODO: instead of displaying plain error messages:
	//       - redirect to /index with an embedded error message
	//       - rework `renderTemplate`
	if len(title) == 0 {
		http.Error(w, "A Page's minimum title length is 1", http.StatusBadRequest)
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
