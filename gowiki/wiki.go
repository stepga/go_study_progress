package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
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

// TODO: redirect to index with embedded error message
// instead of showing a 404 error page
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	pathValidator := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	matches := pathValidator.FindStringSubmatch(r.URL.Path)
	if matches == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return matches[2], nil
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
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	var port_str = fmt.Sprintf(":%d", port)
	log.Printf("[INFO] listening on port: %s\n", port_str)
	log.Fatalf("[ERROR] ListenAndServe: %v\n", http.ListenAndServe(port_str, nil))
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
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
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
