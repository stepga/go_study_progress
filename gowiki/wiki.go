package main

import (
	"fmt"
	"log"
	"os"
)

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
	p1 := &Page{Title: "First Test Page", Body: []byte("This is a sample Page.")}
	err := p1.save()
	if err != nil {
		log.Printf("[ERROR] Saving Page '%s' failed: %v\n", p1.filename(), err)
	}
	p2, err := loadPage("First Test Page")
	if err != nil {
		log.Printf("[ERROR] Loading Page '%s' failed: %v\n", p2.filename(), err)
	} else {
		fmt.Printf("%s\n", p2.Body)
	}
}
